package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"my-auction-market-api/internal/contracts/my_auction"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 从环境变量获取 RPC URL
	rpcURL := os.Getenv("ETH_RPC_URL")
	if rpcURL == "" {
		log.Fatal("Please set ETH_RPC_URL environment variable")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	auctionContractAddress := common.HexToAddress("0x0603f34e8857e813FFC84768F3227F05462AC353")
	auctionContract, err := my_auction.NewMyXAuctionV2(auctionContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	usdValue, err := auctionContract.ConvertToUSDValue(&bind.CallOpts{},
		common.HexToAddress("0x0000000000000000000000000000000000000000"), big.NewInt(int64(0.1*math.Pow10(18))), //1 ETH
		// common.HexToAddress("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238"), big.NewInt(int64(1*math.Pow10(6))), //2 USDC
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("usdValue:", usdValue.Uint64())
	fmt.Println("usdValue:", big.NewFloat(0).Quo(big.NewFloat(float64(usdValue.Int64())), big.NewFloat(float64(math.Pow10(8)))))

	price, err := auctionContract.GetTokenPrice(&bind.CallOpts{}, common.HexToAddress("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("price:", price.Price.Uint64())
	fmt.Println("price:", price.Decimals)

	//function createAuction(address _nftAddress, uint256 _tokenId, uint256 _startPrice, uint256 _startTime, uint256 _endTime) returns(uint256)
	// startTime := time.Now().Unix() + 1000000000
	// tx, err := auctionContract.CreateAuction(&bind.TransactOpts{
	// 	From: common.HexToAddress("0x7cf2fb7ab63b7b064f3c9ea52d1432fa75ec67c4"),
	// },
	// 	common.HexToAddress("0xc61adf5878cb7f414d5fc63aba32835f3da97c87"), //nftAddress
	// 	big.NewInt(2), //nftId
	// 	big.NewInt(int64(298.228*math.Pow10(8))), //startPrice in USD
	// 	big.NewInt(startTime),                    //startTime
	// 	big.NewInt(startTime+1000000000),         //endTime
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(tx.Hash())
	fmt.Println(common.BigToAddress(big.NewInt(0)) == common.HexToAddress("0x0000000000000000000000000000000000000000"))
	// totalAuctionsCreated, totalBidsPlaced, platformFee, totalValueLocked
	totalAuctionsCreated, totalBidsPlaced, platformFee, totalValueLocked, err := auctionContract.GetAuctionSimpleStats(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("totalAuctionsCreated:", totalAuctionsCreated.Uint64())
	fmt.Println("totalBidsPlaced:", totalBidsPlaced.Uint64())
	fmt.Println("platformFee:", platformFee.Uint64())
	fmt.Println("totalValueLocked:", totalValueLocked.Uint64())
}
