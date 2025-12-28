package utils

import (
	"fmt"
	"math/big"
	"my-auction-market-api/internal/contracts/erc20_metadata"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

const (
	// WeiPerETH 1 ETH = 10^18 wei
	WeiPerETH = 1000000000000000000

	// ETHAddress ETH 代币地址（零地址表示原生 ETH）
	ETHAddress = "0x0000000000000000000000000000000000000000"

	// ChainIDMainnet 以太坊主网 ChainID
	ChainIDMainnet int64 = 1
	// ChainIDSepolia Sepolia 测试网 ChainID
	ChainIDSepolia int64 = 11155111
)

var (
	// USDCAddressesByChainID 不同网络的 USDC 地址映射
	// 注意：不同网络 USDC 地址不同，需要根据实际网络配置
	USDCAddressesByChainID = map[int64]string{
		ChainIDMainnet: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", // 主网 USDC
		ChainIDSepolia: "0x1c7d4b196cb0c7b01d743fbc6116a902379c7238", // Sepolia 测试网 USDC
	}
)

// GetUSDCAddress 根据 ChainID 获取对应网络的 USDC 地址
func GetUSDCAddress(chainID int64) (string, error) {
	address, exists := USDCAddressesByChainID[chainID]
	if !exists {
		return "", fmt.Errorf("USDC address not configured for chain ID %d", chainID)
	}
	return address, nil
}

// WeiToETH 将 wei 转换为 ETH (decimal.Decimal)
func WeiToETH(wei *big.Int) decimal.Decimal {
	if wei == nil {
		return decimal.Zero
	}
	return decimal.NewFromBigInt(wei, -18)
}

// ETHToWei 将 ETH (decimal.Decimal) 转换为 wei
func ETHToWei(eth decimal.Decimal) *big.Int {
	// 乘以 10^18 并转换为 big.Int
	weiDecimal := eth.Mul(decimal.NewFromInt(WeiPerETH))
	return weiDecimal.BigInt()
}

// WeiToETHString 将 wei (string) 转换为 ETH (decimal.Decimal)
func WeiToETHString(weiStr string) (decimal.Decimal, error) {
	wei, ok := new(big.Int).SetString(weiStr, 10)
	if !ok {
		return decimal.Zero, nil
	}
	return WeiToETH(wei), nil
}

// ETHToWeiString 将 ETH (decimal.Decimal) 转换为 wei (string)
func ETHToWeiString(eth decimal.Decimal) string {
	wei := ETHToWei(eth)
	return wei.String()
}

// IsSupportedToken 检查代币是否是平台支持的支付代币
// chainID: 当前网络的 ChainID，用于确定支持的 USDC 地址
func IsSupportedToken(tokenAddress string, chainID int64) bool {
	// 规范化地址（转为小写）进行比较
	normalizedAddr := strings.ToLower(tokenAddress)

	// ETH 在所有网络都支持
	if normalizedAddr == strings.ToLower(ETHAddress) {
		return true
	}

	// 根据 ChainID 获取对应的 USDC 地址
	usdcAddress, err := GetUSDCAddress(chainID)
	if err != nil {
		return false
	}

	return normalizedAddr == strings.ToLower(usdcAddress)
}

// ERC20Token 获取 ERC20 代币信息
// 注意：此函数会检查代币是否是平台支持的支付代币
// chainID: 当前网络的 ChainID，用于确定支持的 USDC 地址
func ERC20Token(client *ethclient.Client, tokenAddress string, chainID int64) (name string, symbol string, decimals uint8, totalSupply *big.Int, err error) {
	// 规范化地址进行比较
	normalizedAddr := strings.ToLower(tokenAddress)

	// 检查是否是平台支持的代币
	if !IsSupportedToken(tokenAddress, chainID) {
		usdcAddress, _ := GetUSDCAddress(chainID)
		return "", "", 0, nil, fmt.Errorf("token %s is not a supported payment token. Supported tokens: ETH (%s), USDC (%s)",
			tokenAddress, ETHAddress, usdcAddress)
	}

	if normalizedAddr == strings.ToLower(ETHAddress) {
		return "Ethereum", "ETH", 18, nil, nil
	}

	// USDC 或其他支持的 ERC20 代币
	erc20Contract, err := erc20_metadata.NewIERC20MetadataSolIERC20Metadata(
		common.HexToAddress(tokenAddress),
		bind.ContractBackend(client))
	if err != nil {
		return "", "", 0, nil, fmt.Errorf("failed to create ERC20 contract: %w", err)
	}
	name, err = erc20Contract.Name(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, nil, fmt.Errorf("failed to get name: %w", err)
	}
	symbol, err = erc20Contract.Symbol(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, nil, fmt.Errorf("failed to get symbol: %w", err)
	}
	decimals, err = erc20Contract.Decimals(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, nil, fmt.Errorf("failed to get decimals: %w", err)
	}
	totalSupply, err = erc20Contract.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, nil, fmt.Errorf("failed to get total supply: %w", err)
	}
	return
}
