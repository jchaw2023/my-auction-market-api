package ethereum

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"my-auction-market-api/internal/config"
	"my-auction-market-api/internal/logger"
)

// EtherscanClient Etherscan API客户端
type EtherscanClient struct {
	apiKey  string
	chainID int64
	baseURL string
	client  *http.Client
}

// NewEtherscanClient 创建Etherscan客户端
func NewEtherscanClient(cfg config.EtherscanConfig) *EtherscanClient {
	baseURL := "https://api.etherscan.io/v2/api"
	return &EtherscanClient{
		apiKey:  cfg.APIKey,
		chainID: cfg.ChainID,
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// NFTTransaction Etherscan返回的NFT交易记录
type NFTTransaction struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	TokenID           string `json:"tokenID"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	MethodID          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
	Confirmations     string `json:"confirmations"`
}

// EtherscanResponse Etherscan API响应
type EtherscanResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Result  []NFTTransaction `json:"result"`
}

// GetNFTTransactions 获取地址的NFT交易记录
func (e *EtherscanClient) GetNFTTransactions(address string, startBlock uint64, page int) ([]NFTTransaction, error) {
	// 构建请求URL
	/**
	测试版只能由这个api获取NFT交易记录
	https://api.etherscan.io/v2/api?apikey=YourApiKeyToken&chainid=1&address=0x0603f34e8857e813FFC84768F3227F05462AC353&module=account&action=tokennfttx&startblock=0&sort=desc&page=1&offset=100
	*/
	params := url.Values{}
	params.Set("apikey", e.apiKey)
	params.Set("chainid", fmt.Sprintf("%d", e.chainID))
	params.Set("address", strings.ToLower(address))
	params.Set("module", "account")
	params.Set("action", "tokennfttx")
	// params.Set("startblock", "0")
	params.Set("sort", "desc")
	params.Set("page", fmt.Sprintf("%d", page))
	params.Set("offset", "100")
	if startBlock > 0 {
		params.Set("startblock", fmt.Sprintf("%d", startBlock))
	}
	reqURL := fmt.Sprintf("%s?%s", e.baseURL, params.Encode())

	logger.Info("fetching NFT transactions from Etherscan: %s", address)

	// 发送HTTP请求
	resp, err := e.client.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("failed to request Etherscan API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Etherscan API returned status %d: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var etherscanResp EtherscanResponse
	if err := json.NewDecoder(resp.Body).Decode(&etherscanResp); err != nil {
		return nil, fmt.Errorf("failed to decode Etherscan response: %w", err)
	}

	// 检查API响应状态
	if etherscanResp.Status != "1" {
		if etherscanResp.Message == "No transactions found" {
			return []NFTTransaction{}, nil
		}
		return nil, fmt.Errorf("Etherscan API error: %s", etherscanResp.Message)
	}

	// 统一规范化所有地址（在数据解析时处理）
	for i := range etherscanResp.Result {
		etherscanResp.Result[i].From = strings.ToLower(etherscanResp.Result[i].From)
		etherscanResp.Result[i].To = strings.ToLower(etherscanResp.Result[i].To)
		etherscanResp.Result[i].ContractAddress = strings.ToLower(etherscanResp.Result[i].ContractAddress)
	}

	return etherscanResp.Result, nil
}
