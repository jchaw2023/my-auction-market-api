package handlers

import (
	"github.com/gin-gonic/gin"

	"my-auction-market-api/internal/config"
	"my-auction-market-api/internal/response"
)

// EthereumConfigResponse 以太坊配置响应
type EthereumConfigResponse struct {
	RPCURL                 string `json:"rpcUrl"`
	AuctionContractAddress string `json:"auctionContractAddress"`
	ChainID                int64  `json:"chainId"`
}

// GetEthereumConfig godoc
// @Summary      Get Ethereum configuration
// @Description  Get Ethereum network configuration including RPC URL, contract address, and chain ID
// @Tags         config
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response{data=EthereumConfigResponse}
// @Failure      500  {object}  response.Response
// @Router       /config/ethereum [get]
func GetEthereumConfig(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		response.Success(c, EthereumConfigResponse{
			RPCURL:                 cfg.Ethereum.RPCURL,
			AuctionContractAddress: cfg.Ethereum.AuctionContractAddress,
			ChainID:                cfg.Ethereum.ChainID,
		})
	}
}
