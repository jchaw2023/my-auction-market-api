package services

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"my-auction-market-api/internal/config"
	erc721_nft "my-auction-market-api/internal/contracts/erc721_nft"
	"my-auction-market-api/internal/database"
	"my-auction-market-api/internal/ethereum"
	"my-auction-market-api/internal/logger"
	"my-auction-market-api/internal/models"
	"my-auction-market-api/internal/page"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

// NFTMetadata NFT元数据结构（符合ERC-721标准）
type NFTMetadata struct {
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	Image           string                 `json:"image"`
	ExternalURL     string                 `json:"external_url,omitempty"`
	AnimationURL    string                 `json:"animation_url,omitempty"`
	Attributes      []Attribute            `json:"attributes,omitempty"`
	Properties      map[string]interface{} `json:"properties,omitempty"`
	BackgroundColor string                 `json:"background_color,omitempty"`
}

// Attribute NFT属性
type Attribute struct {
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType string      `json:"display_type,omitempty"`
}

type NFTService struct {
	etherscanClient *ethereum.EtherscanClient
	ethClient       *ethereum.Client
	config          config.EthereumConfig
	httpClient      *resty.Client
}

// OnNFTApproved 处理NFT授权事件，更新授权状态
// ownerAddress: NFT拥有者的钱包地址
// nftContractAddressStr: NFT合约地址
// tokenId: Token ID
// 返回值: nftId, error
func (s *NFTService) OnNFTApproved(ownerAddress string, nftContractAddressStr string, tokenId uint64) (string, error) {
	// 1. 根据 ownerAddress 获取 user_id
	var user models.User
	normalizedOwnerAddr := strings.ToLower(ownerAddress)
	if err := database.DB.Where("wallet_address = ?", normalizedOwnerAddr).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("User not found for wallet address: %s", ownerAddress)
			return "", fmt.Errorf("user not found for wallet address: %w", err)
		}
		logger.Error("Failed to query user by wallet address %s: %s", ownerAddress, err.Error())
		return "", fmt.Errorf("failed to query user by wallet address %s: %w", ownerAddress, err)
	}

	// 2. 根据 nftContractAddressStr 和 tokenId 生成 nftId
	normalizedContractAddr := strings.ToLower(nftContractAddressStr)
	nftID := GenerateNFTID(normalizedContractAddr, tokenId)

	// 3. 更新数据库表 nft_ownerships 的 approved 为 true
	updates := map[string]interface{}{
		"approved":   true,
		"updated_at": time.Now(),
	}

	result := database.DB.Model(&models.NFTOwnership{}).
		Where("nft_id = ? AND user_id = ?", nftID, user.ID).
		Updates(updates)

	if result.Error != nil {
		logger.Error("Failed to update NFT ownership approved status: NFTID=%s, UserID=%d, Error=%s",
			nftID, user.ID, result.Error.Error())
		return nftID, fmt.Errorf("failed to update NFT ownership approved status: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		logger.Warn("No NFT ownership record found to update: NFTID=%s, UserID=%d",
			nftID, user.ID)
		return nftID, fmt.Errorf("no NFT ownership record found to update: %w", result.Error)
	}

	logger.Info("Updated NFT approval status: NFTID=%s, UserID=%d, Contract=%s, TokenID=%d",
		nftID, user.ID, normalizedContractAddr, tokenId)
	return nftID, nil
}

func NewNFTService(ethCfg config.EthereumConfig, etherscanCfg config.EtherscanConfig) (*NFTService, error) {
	// 初始化以太坊客户端
	ethClient, err := ethereum.NewClient(ethCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Ethereum client: %w", err)
	}

	// 初始化 HTTP 客户端（使用 resty）
	httpClient := resty.New().
		SetTimeout(10*time.Second).
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", "NFT-Market-API/1.0")

	return &NFTService{
		etherscanClient: ethereum.NewEtherscanClient(etherscanCfg),
		ethClient:       ethClient,
		config:          ethCfg,
		httpClient:      httpClient,
	}, nil
}

// GetEthClient 获取以太坊客户端（用于链上查询）
func (s *NFTService) GetEthClient() *ethereum.Client {
	return s.ethClient
}

// Close 关闭以太坊客户端连接
func (s *NFTService) Close() {
	if s.ethClient != nil {
		s.ethClient.Close()
	}
}

// GetOwnedNFTs 获取用户拥有的NFT（从链上查询）
// TODO: 实现业务逻辑
func (s *NFTService) GetOwnedNFTs(userID uint64, contractAddress string, query page.PageQuery) ([]models.NFT, int64, error) {
	// TODO: 实现业务逻辑
	// 1. 从JWT或数据库获取用户钱包地址
	// 2. 调用以太坊客户端查询NFT
	// 3. 返回NFT列表
	return nil, 0, nil
}

// SyncNFTs 同步用户NFT（使用Etherscan API）
func (s *NFTService) SyncNFTs(userID uint64) (*models.NFTSyncResult, error) {
	// 1. 从数据库获取用户钱包地址
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// 规范化用户钱包地址
	walletAddress := strings.ToLower(user.WalletAddress)

	// 查询最后一次同步的 NFTOwnership 记录，判断是否之前同步过
	// 如果存在记录，使用其 blockNumber 作为起始区块；如果不存在，说明是首次同步，从区块 0 开始
	var lastSyncedOwnership models.NFTOwnership
	var startBlockNumber uint64 = 0

	if err := database.DB.Model(&models.NFTOwnership{}).
		Where("user_id = ?", userID).
		Order("block_number DESC").
		Preload("NFT").
		First(&lastSyncedOwnership).Error; err != nil {
		// 如果没有找到记录，说明是首次同步
		if err == gorm.ErrRecordNotFound {
			logger.Info("User %d has no previous sync record, starting first sync from block 0", userID)
		} else {
			// 其他数据库错误
			logger.Error("failed to get last synced NFTOwnership for user %d: %s", userID, err.Error())
			logger.Info("Continuing with first sync from block 0")
			return nil, fmt.Errorf("failed to get last synced NFTOwnership for user %d: %w", userID, err)
		}
		// startBlockNumber 保持为 0，表示从最开始同步
	} else {
		// 找到了上次同步记录，使用其 blockNumber 作为起始区块
		startBlockNumber = lastSyncedOwnership.BlockNumber
		logger.Info("User %d has previous sync record: NFTID=%s, Timestamp=%d, BlockNumber=%d, starting incremental sync from block %d",
			userID, lastSyncedOwnership.NFTID, lastSyncedOwnership.Timestamp, lastSyncedOwnership.BlockNumber, startBlockNumber)
	}

	// 从链上同步NFT数据（根据是否有上次同步记录决定起始区块号）
	var newNFTDatas *[]ChainNFTData
	syncNFTsFromChainResponse, err := s.SyncNFTsFromChain(walletAddress, startBlockNumber, 1)
	if err != nil {
		logger.Error("failed to sync NFTs from chain: %s", err.Error())
		return nil, fmt.Errorf("failed to sync NFTs from chain!")
	}
	if syncNFTsFromChainResponse.TotalFound > 0 {
		newNFTDatas = &syncNFTsFromChainResponse.NFTs
	}

	// 初始化同步结果
	result := &models.NFTSyncResult{
		TotalFound:  0,
		TotalSynced: 0,
		TotalFailed: 0,
	}

	if newNFTDatas != nil {
		result.TotalFound = len(*newNFTDatas)
		for _, nftdata := range *newNFTDatas {
			//将nftdata原数据保存到数据库:如果存在则更新 如果metadata description image nft_name任意一项不存在则更新记录
			if err := s.saveOrUpdateNFTFromChainData(nftdata, userID); err != nil {
				logger.Error("failed to save/update NFT from chain data: %s, error: %s", nftdata.NFTID, err.Error())
				result.TotalFailed++
				// 继续处理下一个NFT，不中断流程
			} else { //开始处理用户关系 就是nft_ownerships里面的数据
				if err := s.saveOrUpdateNFTOwnership(nftdata, userID); err != nil {
					logger.Error("failed to save/update NFT ownership: %s, error: %s", nftdata.NFTID, err.Error())
					result.TotalFailed++
					// 继续处理下一个NFT，不中断流程
				} else {
					result.TotalSynced++
				}
			}
		}
	}

	return result, nil
}

// GenerateNFTID 生成NFT的唯一ID（基于合约地址和TokenID）
// 注意：contractAddress 应该在调用前已规范化
func GenerateNFTID(contractAddress string, tokenID uint64) string {
	// 地址应在数据解析时已规范化，此处直接使用
	// 生成唯一标识：contractAddress:tokenID
	key := fmt.Sprintf("%s:%d", strings.ToLower(contractAddress), tokenID)
	// 使用MD5哈希生成唯一ID（32个字符）
	hash := md5.Sum([]byte(key))
	// 返回hex编码的32个字符（完整MD5哈希值）
	return hex.EncodeToString(hash[:])
}

// ChainNFTData 链上NFT数据（不包含数据库操作）
type ChainNFTData struct {
	NFTID           string `json:"nftId"`           // NFT唯一标识
	ContractAddress string `json:"contractAddress"` // 合约地址
	TokenID         uint64 `json:"tokenId"`         // Token ID
	IsOwned         bool   `json:"isOwned"`         // 是否拥有
	Timestamp       uint64 `json:"timestamp"`       // 交易时间戳
	BlockNumber     uint64 `json:"blockNumber"`     // 区块号
	TransactionHash string `json:"transactionHash"` // 交易哈希
	From            string `json:"from"`            // 发送方地址
	To              string `json:"to"`              // 接收方地址
}

// SyncNFTsFromChainResponse 从链上同步NFT的响应
type SyncNFTsFromChainResponse struct {
	NFTs       []ChainNFTData `json:"nfts"`       // NFT列表
	TotalFound int            `json:"totalFound"` // 找到的NFT数量
	HasMore    bool           `json:"hasMore"`    // 是否还有更多数据
	NextPage   int            `json:"nextPage"`   // 下一页页码
	NextBlock  uint64         `json:"nextBlock"`  // 下一个起始区块号
}

// GetSyncStatus 获取NFT同步状态
// TODO: 实现业务逻辑
func (s *NFTService) GetSyncStatus(userID uint64) (*models.NFTSyncStatus, error) {
	// TODO: 实现业务逻辑
	// 1. 查询用户的上次同步信息
	// 2. 查询数据库中NFT总数
	// 3. 检查是否有正在进行的同步任务
	return nil, nil
}

// GetByID 根据ID获取NFT
// TODO: 实现业务逻辑
func (s *NFTService) GetByID(id uint64) (*models.NFT, error) {
	// TODO: 实现业务逻辑
	// 1. 从数据库查询NFT
	// 2. 返回NFT信息
	return nil, nil
}

// VerifyOwnership 验证NFT所有权
// TODO: 实现业务逻辑
func (s *NFTService) VerifyOwnership(userID uint64, payload models.NFTOwnershipVerifyPayload) (bool, error) {
	// TODO: 实现业务逻辑
	// 1. 从数据库获取用户钱包地址
	// 2. 调用以太坊客户端验证所有权
	// 3. 返回验证结果
	return false, nil
}

// GetMyNFTsList 获取用户的NFT列表（返回 nft_ownerships 数据，关联 NFT 元数据）
// 从 nft_ownerships 表查询，通过 Preload 关联查询对应的 NFT 元数据
// statusFilter: 状态筛选参数，可选值：holding, selling, sold, transfered, all（默认 all）
func (s *NFTService) GetMyNFTsList(userID uint64, statusFilter string) ([]models.NFTOwnership, error) {
	var ownerships []models.NFTOwnership

	// 构建查询
	query := database.DB.Model(&models.NFTOwnership{}).
		Where("user_id = ?", userID)

	// 如果指定了状态筛选且不是 "all"，则添加状态过滤条件
	if statusFilter != "" && statusFilter != "all" {
		query = query.Where("status = ?", statusFilter)
	}

	// 从 nft_ownerships 表查询用户的NFT关系记录
	// 通过 Preload("NFT") 关联查询对应的 NFT 元数据
	if err := query.
		Preload("NFT").
		Order("timestamp DESC").
		Find(&ownerships).Error; err != nil {
		return nil, fmt.Errorf("failed to query NFT ownerships: %w", err)
	}

	return ownerships, nil
}

// GetMyNFTOwnershipByNFTID 根据 nftId 和 userId 获取单个 NFT 关系数据
func (s *NFTService) GetMyNFTOwnershipByNFTID(userID uint64, nftID string) (*models.NFTOwnership, error) {
	var ownership models.NFTOwnership

	// 从 nft_ownerships 表查询指定 NFT 的关系记录
	// 通过 Preload("NFT") 关联查询对应的 NFT 元数据
	if err := database.DB.Model(&models.NFTOwnership{}).
		Where("user_id = ? AND nft_id = ?", userID, nftID).
		Preload("NFT").
		First(&ownership).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("NFT ownership not found: nftId=%s, userId=%d", nftID, userID)
		}
		return nil, fmt.Errorf("failed to query NFT ownership: %w", err)
	}

	return &ownership, nil
}

// SyncNFTsFromChain 按照BlockNumber和分页参数同步钱包所有NFT的链上最新数据
// 该方法只获取链上数据，不处理数据库记录
// 通过递增pageNum循环获取所有数据，直到返回的transactions为空
func (s *NFTService) SyncNFTsFromChain(walletAddress string, startBlock uint64, startPage int) (*SyncNFTsFromChainResponse, error) {
	// 规范化钱包地址

	// 使用map来跟踪每个NFT的最终状态（以最新交易为准）
	// Etherscan返回的交易按时间降序排列（最新的在前面），所以只需要处理每个NFT的第一条交易
	nftMap := make(map[string]*ChainNFTData)

	// 循环获取所有页面的数据
	pageNum := startPage
	maxBlockNumber := uint64(0)

	for {
		// 调用Etherscan API获取NFT交易记录
		transactions, err := s.etherscanClient.GetNFTTransactions(walletAddress, startBlock, pageNum)
		if err != nil {
			logger.Error("failed to get NFT transactions from Etherscan (page %d): %s", pageNum, err.Error())
			return nil, fmt.Errorf("failed to get NFT transactions from Etherscan (page %d)!", pageNum)
		}

		// 如果返回的交易为空，说明已经获取完所有数据
		if len(transactions) == 0 {
			break
		}

		// 处理交易记录
		for _, tx := range transactions {
			if tx.TokenSymbol == "" {
				continue
			}

			// 解析TokenID
			tokenID, err := strconv.ParseUint(tx.TokenID, 10, 64)
			if err != nil {
				logger.Error("failed to parse token ID: %s", tx.TokenID)
				continue
			}

			// 解析时间戳
			timestamp, err := strconv.ParseUint(tx.TimeStamp, 10, 64)
			if err != nil {
				logger.Error("failed to parse timestamp: %s", tx.TimeStamp)
				continue
			}

			// 解析区块号
			blockNumber, err := strconv.ParseUint(tx.BlockNumber, 10, 64)
			if err != nil {
				logger.Error("failed to parse block number: %s", tx.BlockNumber)
				continue
			}

			// 记录最大区块号
			if blockNumber > maxBlockNumber {
				maxBlockNumber = blockNumber
			}

			// 生成NFT的唯一ID（用于去重）
			key := GenerateNFTID(tx.ContractAddress, tokenID)

			// 如果NFT已经在map中，说明已经处理过更新的交易（因为按降序排列），跳过
			if _, exists := nftMap[key]; exists {
				continue
			}

			// 这是该NFT的第一条交易（最新的），根据交易类型确定最终状态
			isOwned := false
			if tx.To == walletAddress {
				// 转入：用户是接收方，说明当前拥有
				isOwned = true
			} else if tx.From == walletAddress {
				// 转出：用户是发送方，说明当前不拥有
				isOwned = false
			}

			// 记录该NFT的链上数据
			nftMap[key] = &ChainNFTData{
				NFTID:           key,
				ContractAddress: tx.ContractAddress, // 已规范化
				TokenID:         tokenID,
				IsOwned:         isOwned,
				Timestamp:       timestamp,
				BlockNumber:     blockNumber,
				TransactionHash: tx.Hash,
				From:            tx.From, // 已规范化
				To:              tx.To,   // 已规范化
			}
		}

		// 递增页码，继续获取下一页
		pageNum++
	}

	// 转换为切片
	nfts := make([]ChainNFTData, 0, len(nftMap))
	for _, nft := range nftMap {
		nfts = append(nfts, *nft)
	}

	// 计算下一个起始区块号（使用最大区块号）
	nextBlock := maxBlockNumber
	if nextBlock == 0 {
		nextBlock = startBlock
	}

	return &SyncNFTsFromChainResponse{
		NFTs:       nfts,
		TotalFound: len(nfts),
		HasMore:    false, // 已经获取完所有数据
		NextPage:   0,     // 没有更多数据
		NextBlock:  nextBlock,
	}, nil
}

// saveOrUpdateNFTFromChainData 保存或更新NFT数据（从链上数据）
// 如果NFT不存在则写入，如果存在且metadata、description、image、nft_name任意一项不存在则更新记录
func (s *NFTService) saveOrUpdateNFTFromChainData(chainData ChainNFTData, userID uint64) error {
	// 生成NFT唯一ID
	nftID := chainData.NFTID

	// 先检查NFT是否已存在
	var existingNFT models.NFT
	result := database.DB.Where("nft_id = ?", nftID).First(&existingNFT)

	// 获取NFT合约信息（用于获取TokenURI和Metadata）
	contractAddr := common.HexToAddress(chainData.ContractAddress)
	mynft, err := erc721_nft.NewMyNFT(contractAddr, s.ethClient.GetClient())
	if err != nil {
		return fmt.Errorf("failed to create MyNFT contract for %s: %w", chainData.ContractAddress, err)
	}

	tokenIDBigInt := big.NewInt(int64(chainData.TokenID))
	opts := &bind.CallOpts{Context: context.Background()}

	// 获取合约基本信息
	mynftName, _ := mynft.Name(opts)
	mynftSymbol, _ := mynft.Symbol(opts)
	mynftTokenURI, _ := mynft.TokenURI(opts, tokenIDBigInt)
	mynftOwner, err := mynft.OwnerOf(opts, tokenIDBigInt)
	if err != nil {
		return fmt.Errorf("failed to get owner for NFT %s:%d: %w", chainData.ContractAddress, chainData.TokenID, err)
	}

	normalizedOwnerAddr := strings.ToLower(mynftOwner.String())

	if result.Error != nil {
		// NFT不存在，创建新记录
		var metadata *NFTMetadata
		var metadataJSONStr string

		// 获取Metadata
		if mynftTokenURI != "" {
			metadataJSON, err := s.fetchMetadataFromURI(mynftTokenURI)
			if err != nil {
				logger.Error("failed to fetch metadata from URI %s: %s", mynftTokenURI, err.Error())
			} else {
				var parsedMetadata NFTMetadata
				if err := json.Unmarshal(metadataJSON, &parsedMetadata); err != nil {
					logger.Error("failed to unmarshal metadata JSON: %s", err.Error())
				} else {
					metadata = &parsedMetadata
					metadataJSONStr = string(metadataJSON)
				}
			}
		}

		if metadata == nil {
			metadata = &NFTMetadata{}
		}

		// 创建新记录
		nftRecord := models.NFT{
			NFTID:           nftID,
			UserID:          userID,
			ContractAddress: chainData.ContractAddress,
			TokenID:         chainData.TokenID,
			TokenURI:        mynftTokenURI,
			ContractName:    mynftName,
			ContractSymbol:  mynftSymbol,
			NftOwnerAddress: normalizedOwnerAddr,
			NftName:         metadata.Name,
			Image:           metadata.Image,
			Description:     metadata.Description,
			Metadata:        metadataJSONStr,
			LastSyncedAt:    time.Now(),
		}

		if err := database.DB.Create(&nftRecord).Error; err != nil {
			return fmt.Errorf("failed to create NFT: %w", err)
		}
		logger.Info("Created new NFT record from chain data: %s (contract: %s, token: %d)", nftID, chainData.ContractAddress, chainData.TokenID)
	} else {
		// NFT存在，检查是否需要更新
		needUpdate := false
		updates := map[string]interface{}{}

		// 检查metadata、description、image、nft_name是否为空
		if existingNFT.Metadata == "" || existingNFT.Description == "" || existingNFT.Image == "" || existingNFT.NftName == "" {
			needUpdate = true

			// 获取Metadata
			var metadata *NFTMetadata
			var metadataJSONStr string

			if mynftTokenURI != "" {
				metadataJSON, err := s.fetchMetadataFromURI(mynftTokenURI)
				if err != nil {
					logger.Error("failed to fetch metadata from URI %s: %s", mynftTokenURI, err.Error())
				} else {
					var parsedMetadata NFTMetadata
					if err := json.Unmarshal(metadataJSON, &parsedMetadata); err != nil {
						logger.Error("failed to unmarshal metadata JSON: %s", err.Error())
					} else {
						metadata = &parsedMetadata
						metadataJSONStr = string(metadataJSON)
					}
				}
			}

			if metadata != nil {
				// 只更新为空的字段
				if existingNFT.Metadata == "" && metadataJSONStr != "" {
					updates["metadata"] = metadataJSONStr
				}
				if existingNFT.Description == "" && metadata.Description != "" {
					updates["description"] = metadata.Description
				}
				if existingNFT.Image == "" && metadata.Image != "" {
					updates["image"] = metadata.Image
				}
				if existingNFT.NftName == "" && metadata.Name != "" {
					updates["nft_name"] = metadata.Name
				}
			}

			// 更新其他可能缺失的字段
			if existingNFT.TokenURI == "" && mynftTokenURI != "" {
				updates["token_uri"] = mynftTokenURI
			}
			if existingNFT.ContractName == "" && mynftName != "" {
				updates["contract_name"] = mynftName
			}
			if existingNFT.ContractSymbol == "" && mynftSymbol != "" {
				updates["contract_symbol"] = mynftSymbol
			}
		}

		// 始终更新nft_owner_address和last_synced_at（无论是否需要更新其他字段）
		updates["nft_owner_address"] = normalizedOwnerAddr
		updates["last_synced_at"] = time.Now()

		// 如果有需要更新的字段，执行更新
		if needUpdate && len(updates) > 2 { // 除了nft_owner_address和last_synced_at还有其他字段需要更新
			if err := database.DB.Model(&existingNFT).
				Where("nft_id = ?", nftID).
				Updates(updates).Error; err != nil {
				return fmt.Errorf("failed to update NFT: %w", err)
			}
			logger.Info("Updated NFT record from chain data: %s (contract: %s, token: %d)", nftID, chainData.ContractAddress, chainData.TokenID)
		} else {
			// 即使不需要更新其他字段，也要更新nft_owner_address和last_synced_at
			if err := database.DB.Model(&existingNFT).
				Where("nft_id = ?", nftID).
				Updates(updates).Error; err != nil {
				return fmt.Errorf("failed to update NFT ownership: %w", err)
			}
			logger.Debug("Updated NFT owner address: %s (contract: %s, token: %d, owner: %s)",
				nftID, chainData.ContractAddress, chainData.TokenID, normalizedOwnerAddr)
		}
	}

	return nil
}

// saveOrUpdateNFTOwnership 保存或更新NFT用户关系（nft_ownerships表）
// nft_id 与 user_id 是联合唯一索引
func (s *NFTService) saveOrUpdateNFTOwnership(chainData ChainNFTData, userID uint64) error {
	// owner_address 表示在当前用户操作下NFT的去向（交易中的To地址），不一定是最终的拥有者
	// 如果用户是发送方，owner_address是接收方地址；如果用户是接收方，owner_address是用户自己的地址
	normalizedOwnerAddr := strings.ToLower(chainData.To)

	// 获取平台合约地址（用于判断是否转移到平台）
	platformContractAddr := strings.ToLower(s.config.AuctionContractAddress)
	// 判断status
	status := models.NFTOwnershipStatusHolding // 默认状态
	if !chainData.IsOwned {
		// 用户不拥有该NFT
		if normalizedOwnerAddr == platformContractAddr {
			// 接收方是平台合约，不修改status（保持原状态）
			status = "" // 空字符串表示不更新
		} else {
			// 接收方不是平台合约，设置为transfered
			status = models.NFTOwnershipStatusTransfered
		}
	}

	// 检查是否已存在该关系记录
	var existingOwnership models.NFTOwnership
	result := database.DB.Where("nft_id = ? AND user_id = ?", chainData.NFTID, userID).First(&existingOwnership)

	now := time.Now()

	if result.Error != nil {
		// 记录不存在，创建新记录
		// 如果status为空（不拥有且转移到平台），使用transfered作为初始状态
		if status == "" {
			status = models.NFTOwnershipStatusTransfered
		}

		ownership := models.NFTOwnership{
			NFTID:        chainData.NFTID,
			UserID:       int64(userID),
			OwnerAddress: normalizedOwnerAddr,
			Status:       status,
			Approved:     false, // 默认未授权
			Timestamp:    int64(chainData.Timestamp),
			BlockNumber:  chainData.BlockNumber,
			LastSyncedAt: &now,
			CreatedAt:    &now,
			UpdatedAt:    &now,
		}

		if err := database.DB.Create(&ownership).Error; err != nil {
			return fmt.Errorf("failed to create NFT ownership: %w", err)
		}
		logger.Info("Created NFT ownership: NFTID=%s, UserID=%d, Status=%s, Owner=%s",
			chainData.NFTID, userID, status, normalizedOwnerAddr)
	} else {
		// 记录已存在，更新记录
		updates := map[string]interface{}{
			"owner_address":  normalizedOwnerAddr,
			"timestamp":      int64(chainData.Timestamp),
			"block_number":   chainData.BlockNumber,
			"last_synced_at": &now,
			"updated_at":     &now,
		}

		// 只有当status不为空时才更新status
		if status != "" {
			updates["status"] = status
		}

		if err := database.DB.Model(&existingOwnership).
			Where("nft_id = ? AND user_id = ?", chainData.NFTID, userID).
			Updates(updates).Error; err != nil {
			return fmt.Errorf("failed to update NFT ownership: %w", err)
		}

		statusLog := status
		if status == "" {
			statusLog = existingOwnership.Status + "(unchanged)"
		}
		logger.Info("Updated NFT ownership: NFTID=%s, UserID=%d, Status=%s, Owner=%s",
			chainData.NFTID, userID, statusLog, normalizedOwnerAddr)
	}

	return nil
}

// fetchMetadataFromURI 从 tokenURI 获取 metadata JSON 数据（使用 resty）
func (s *NFTService) fetchMetadataFromURI(tokenURI string) ([]byte, error) {
	// 处理 IPFS 协议
	if strings.HasPrefix(tokenURI, "ipfs://") {
		// 将 ipfs:// 转换为 HTTP URL
		ipfsHash := strings.TrimPrefix(tokenURI, "ipfs://")
		tokenURI = fmt.Sprintf("https://ipfs.io/ipfs/%s", ipfsHash)
	} else if strings.HasPrefix(tokenURI, "ipfs/") {
		// 处理 ipfs/ 格式
		ipfsHash := strings.TrimPrefix(tokenURI, "ipfs/")
		tokenURI = fmt.Sprintf("https://ipfs.io/ipfs/%s", ipfsHash)
	}

	// 使用 resty 发送 GET 请求
	resp, err := s.httpClient.R().SetContext(context.Background()).
		Get(tokenURI)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch metadata: %w", err)
	}

	// 检查状态码
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return resp.Body(), nil
}
