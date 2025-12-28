package ethereum

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"my-auction-market-api/internal/config"
	"my-auction-market-api/internal/logger"
)

type Client struct {
	client     *ethclient.Client
	config     config.EthereumConfig
	nftABI     *abi.ABI
	auctionABI *abi.ABI
}

func NewClient(cfg config.EthereumConfig) (*Client, error) {
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}

	logger.Info("connected to Ethereum node: %s", cfg.RPCURL)

	return &Client{
		client: client,
		config: cfg,
	}, nil
}

func NewClientWithWSS(cfg config.EthereumConfig) (*Client, error) {
	// 优先使用 WSS URL，如果没有配置则使用 RPC URL（但订阅功能可能不可用）
	url := cfg.WssURL
	if url == "" {
		url = cfg.RPCURL
		logger.Warn("WSS URL not configured, using RPC URL instead. WebSocket subscriptions may not work.")
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node at %s: %w", url, err)
	}

	logger.Info("connected to Ethereum node (WebSocket): %s", url)

	return &Client{
		client: client,
		config: cfg,
	}, nil
}
func (c *Client) GetClient() *ethclient.Client {
	return c.client
}

func (c *Client) GetChainID() *big.Int {
	return big.NewInt(c.config.ChainID)
}

func (c *Client) GetAuctionContractAddress() common.Address {
	return common.HexToAddress(c.config.AuctionContractAddress)
}

func (c *Client) GetBlockNumber(ctx context.Context) (uint64, error) {
	header, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

func (c *Client) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	balance, err := c.client.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *Client) GetTransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return c.client.TransactionReceipt(ctx, txHash)
}

func (c *Client) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return c.client.FilterLogs(ctx, query)
}

func (c *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return c.client.CallContract(ctx, msg, blockNumber)
}

func (c *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return c.client.PendingNonceAt(ctx, account)
}

func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.client.SuggestGasPrice(ctx)
}

func (c *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return c.client.EstimateGas(ctx, msg)
}

func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.client.SendTransaction(ctx, tx)
}

func (c *Client) GetTransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	return c.client.TransactionByHash(ctx, txHash)
}

func (c *Client) Close() {
	if c.client != nil {
		c.client.Close()
	}
}

// returns a transaction auth for contract interactions
func (c *Client) GetAuth(ctx context.Context, privateKey string) (*bind.TransactOpts, error) {
	// This is a placeholder - in production, you should use a secure key management system
	// For now, this function signature is provided for future implementation
	if privateKey == "" {
		return nil, fmt.Errorf("private key is empty")
	}
	cr, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to convert private key to ECDSA: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(cr, c.GetChainID())
	if err != nil {
		return nil, fmt.Errorf("failed to create keyed transactor: %w", err)
	}
	return auth, nil
}
