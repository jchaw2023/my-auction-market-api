// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package my_auction

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MyXAuctionV2Auction is an auto generated low-level Go binding around an user-defined struct.
type MyXAuctionV2Auction struct {
	NftAddress      common.Address
	TokenId         *big.Int
	Seller          common.Address
	HighestBidder   common.Address
	HighestBidToken common.Address
	HighestBid      *big.Int
	HighestBidValue *big.Int
	StartPrice      *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	Ended           bool
	Cancelled       bool
	BidCount        *big.Int
}

// MyXAuctionV2FeeTier is an auto generated low-level Go binding around an user-defined struct.
type MyXAuctionV2FeeTier struct {
	Threshold *big.Int
	FeeRate   *big.Int
}

// MyXAuctionV2MetaData contains all meta data concerning the MyXAuctionV2 contract.
var MyXAuctionV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cancelledBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmountValue\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalBid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBidValue\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"endedBy\",\"type\":\"address\"}],\"name\":\"AuctionForceEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBidValue\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBidValue\",\"type\":\"uint256\"}],\"name\":\"BidValueTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"DynamicFeeEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tierIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"name\":\"FeeTierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTApprovalCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"PlatformFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTVL\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBidsPlaced\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"change\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isIncrease\",\"type\":\"bool\"}],\"name\":\"TotalValueLockedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approveNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctionBidCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBidValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"bidCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_usdValue\",\"type\":\"uint256\"}],\"name\":\"calculateDynamicFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelNFTApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"cancelUserAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertToUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"endAuctionAndClaimNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeTiers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"forceEndAuctionAndClaimNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllFeeTiers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"internalType\":\"structMyXAuctionV2.FeeTier[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBidValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"bidCount\",\"type\":\"uint256\"}],\"internalType\":\"structMyXAuctionV2.Auction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuctionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuctionSimpleStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuctionStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAuctions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBids\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentPlatformFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"activeAuctions\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getAuctionsBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBidValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"bidCount\",\"type\":\"uint256\"}],\"internalType\":\"structMyXAuctionV2.Auction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeTierCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalValueLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserCreatedAuctionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserCreatedAuctions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBidValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"bidCount\",\"type\":\"uint256\"}],\"internalType\":\"structMyXAuctionV2.Auction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"priceFeeds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setDynamicFeeEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_feeRates\",\"type\":\"uint256[]\"}],\"name\":\"setFeeTiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setPlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_priceFeeds\",\"type\":\"address[]\"}],\"name\":\"setPriceFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAuctionsCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBidsPlaced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValueLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDynamicFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userCreatedAuctionIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080806040523460aa575f51602061381e5f395f51905f525460ff8160401c16609b576002600160401b03196001600160401b038216016049575b60405161376f90816100af8239f35b6001600160401b0319166001600160401b039081175f51602061381e5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80603a565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c9081630a48a7fe146127e75750806312e8e2c31461274c57806318aa08f31461272f5780631a8d52a1146126f75780631d57042c14612332578063230ed44a146122f557806326232a2e146122d85780632b48d9a3146121e75780632ecdee5b14612128578063387296d3146120495780633f4ba83a14611fc05780633fe3b2ce14611d6457806343638e8614611d3c578063450ed73514611cb057806354fd4d5014611c54578063571a26a014611b765780635c975abb14611b535780635cd8a76b14611a315780636da7c22c14611a13578063715018a6146119aa578063742f0a901461122e57806376e11286146111c357806378bd7935146111885780637948c325146110bb5780637a222a4a14610fcf5780638129fc1c14610e345780638456cb5914610da257806385f63a1914610cf55780638da5cb5b14610cc057806392aeb53a14610b985780639600dd8314610265578063961c9ae41461073757806396b5a755146104295780639cde2871146103d25780639dcb511a14610392578063a1127c561461036f578063b26025aa14610247578063b6304c5614610348578063ba8b196e1461032a578063c44e66401461030c578063cb600160146102e2578063cf2d2178146102c4578063d02641a01461028f578063dcd4325d14610265578063ec18154e146102475763f2fde38b14610218575f80fd5b346102445760203660031901126102445761024161023461281b565b61023c61316c565b6130fb565b80f35b80fd5b50346102445780600319360112610244576020600754604051908152f35b50346102445760403660031901126102445761027f61281b565b5061024160ff6003541615612b39565b503461024457602036600319011261024457604060ff6102b56102b061281b565b612fcd565b83519182529091166020820152f35b50346102445780600319360112610244576020600a54604051908152f35b50346102445760203660031901126102445760406020916004358152600683522054604051908152f35b50346102445780600319360112610244576020600154604051908152f35b50346102445780600319360112610244576020600554604051908152f35b5034610244576020366003190112610244576020610367600435612f3f565b604051908152f35b5034610244578060031936011261024457602060ff600954166040519015158152f35b5034610244576020366003190112610244576020906001600160a01b036103b761281b565b168152808252604060018060a01b0391205416604051908152f35b5034610244576040366003190112610244576103ec61281b565b6001600160a01b03168152600b602052604081208054602435929083101561024457602061041a8484612a97565b90549060031b1c604051908152f35b50346102445760203660031901126102445760043561044661316c565b61045560ff6003541615612b39565b61045d6135bc565b61046a6001548210612b7a565b61047381612a0d565b5090600a82019081549160ff83161580610728575b61049190612bbf565b600284019260018060a01b0384541633148015610708575b156106b35761ffff19166101011790558254600684018054909386926001600160a01b03169180610670575b505460018601546001600160a01b0390911690823b15610661576040516323b872dd60e01b81523060048201526001600160a01b0392909216602483015260448201529082908290606490829084905af180156106655761064c575b50506003830180546001600160a01b0316806105b3575b5060018060a01b0390541691600560018060a01b03600486015416940154905490604051938452602084015260408301527f10ec610e9b6f8628e9a57f51bfda4adcb5db4c761791045d0f85214b681c7c4860603393a460015f5160206136fa5f395f51905f525580f35b60048501546001600160a01b0316806106385750508480808060018060a01b038554166005890154905af16105e6612eb4565b50156105f3575b5f610548565b60405162461bcd60e51b815260206004820152601760248201527f4661696c656420746f20726566756e64206269646465720000000000000000006044820152606490fd5b610647916005870154916135f4565b6105ed565b816106569161286a565b61066157835f610531565b8380fd5b6040513d84823e3d90fd5b608061068c5f5160206136da5f395f51905f5292600754612c1a565b8060075560055488549060405192835260208301526040820152856060820152a15f6104d5565b60405162461bcd60e51b815260206004820152602760248201527f4f6e6c792073656c6c6572206f72206f776e65722063616e2063616e63656c2060448201526630bab1ba34b7b760c91b6064820152608490fd5b505f5160206136ba5f395f51905f52546001600160a01b031633146104a9565b50600883901c60ff1615610488565b50346102445760a03660031901126102445761075161281b565b60243590604435906064356084359161076f60ff6003541615612b39565b82821015610b48578315610af8576040516331a9108f60e11b8152600481018690526001600160a01b03919091169390602081602481885afa908115610aed578791610aa7575b50336001600160a01b0390911603610a6b578591843b15610a5c576040516323b872dd60e01b8152336004820152306024820152604481018790528381606481838a5af1908115610a60578491610a47575b5050604051936108178561284d565b85855260208501918783526040860193338552606087018681526080880187815260a089019088825260c08a019289845260e08b019485526101008b019586526101208b019687526101408b01988a8a526101608c01988b8a526101808d019b8c52600154600160401b811015610a31578060016108989201600155612a0d565b9d909d610a1b57518d546001600160a01b03199081166001600160a01b03928316178f55915160018f810191909155925160028f0180548416918316919091179055925160038e0180548316918516919091179055925160048d018054909416921691909117909155905160058a015590516006890155905160078801559051600880880191909155915160098701559251600a86018054935161ffff1990941660ff921515929092169190911792151590911b61ff00169190911790559051600b9290920191909155545f19810192908311610a075761097a600454612c27565b600455338452600b60205260408420805494600160401b8610156109f35750846109ab916001602097018155612a97565b81549060031b9085821b915f19901b1916179055604051908152827f8777bed2a899ba8843de663a1f6ed1c48d071cc8bde08e2488b59c00c1993f76853393a4604051908152f35b634e487b7160e01b81526041600452602490fd5b634e487b7160e01b84526011600452602484fd5b5050634e487b7160e01b8f5260048f905260248ffd5b5050634e487b7160e01b8f52604160045260248ffd5b81610a519161286a565b610a5c57825f610808565b8280fd5b6040513d86823e3d90fd5b60405162461bcd60e51b8152602060048201526014602482015273165bdd481b5d5cdd081bdddb881d1a194813919560621b6044820152606490fd5b90506020813d602011610ae5575b81610ac26020938361286a565b81010312610ae157516001600160a01b0381168103610ae1575f6107b6565b8680fd5b3d9150610ab5565b6040513d89823e3d90fd5b60405162461bcd60e51b815260206004820152602260248201527f5374617274207072696365206d7573742062652067726561746572207468616e604482015261020360f41b6064820152608490fd5b60405162461bcd60e51b815260206004820152602260248201527f53746172742074696d65206d757374206265206265666f726520656e642074696044820152616d6560f01b6064820152608490fd5b50346102445760403660031901126102445760043567ffffffffffffffff8111610cbc57610bca903690600401612a29565b60243567ffffffffffffffff8111610a5c57610bea903690600401612a29565b90610bf361316c565b8051825103610c7e57825b8151811015610c7a57600190610c286001600160a01b03610c1f8387612ad9565b51161515612ef3565b818060a01b03610c388286612ad9565b5116828060a01b03610c4a8386612ad9565b51168652856020526040862090838060a01b03166bffffffffffffffffffffffff60a01b82541617905501610bfe565b8380f35b60405162461bcd60e51b8152602060048201526016602482015275082e4e4c2f2e640d8cadccee8d040dad2e6dac2e8c6d60531b6044820152606490fd5b5080fd5b50346102445780600319360112610244575f5160206136ba5f395f51905f52546040516001600160a01b039091168152602090f35b503461024457602036600319011261024457600435610d1261316c565b610d2160ff6003541615612b39565b610d2e6001548210612b7a565b610d6a610d3a82612a0d565b50600a81016001815460ff81161580610d93575b610d5790612bbf565b42600985015560ff19161790558261319f565b33907f77e5df52a34e08f2c8bf4ea9bea6de1ef93675f92c17e6a246847ab529d385ac8380a380f35b50600881901c60ff1615610d4e565b5034610244578060031936011261024457610dbb61316c565b60035460ff8116610dfe5760019060ff1916176003557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b60405162461bcd60e51b815260206004820152600e60248201526d105b1c9958591e481c185d5cd95960921b6044820152606490fd5b50346102445780600319360112610244575f51602061371a5f395f51905f525467ffffffffffffffff60ff8260401c1615911680159081610fc7575b6001149081610fbd575b159081610fb4575b50610fa55780600167ffffffffffffffff195f51602061371a5f395f51905f525416175f51602061371a5f395f51905f5255610f75575b610ec1613636565b610ec9613636565b610ed2336130fb565b610eda613636565b610ee2613636565b60015f5160206136fa5f395f51905f52558160025560ff196003541660035581600455816005558160075560ff196009541660095581600a55610f225780f35b60ff60401b195f51602061371a5f395f51905f5254165f51602061371a5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b600160401b60ff60401b195f51602061371a5f395f51905f525416175f51602061371a5f395f51905f5255610eb9565b63f92ee8a960e01b8252600482fd5b9050155f610e82565b303b159150610e7a565b829150610e70565b50346102445760403660031901126102445760043560015480821015611076578161100f9161100060243583612aac565b9080821161106e575b50612c1a565b9061101982612d66565b925b828110611034576040518061103086826129c9565b0390f35b8061105261104c61104760019486612aac565b612a0d565b50612db5565b61105c8287612ad9565b526110678186612ad9565b500161101b565b90505f611009565b60405162461bcd60e51b815260206004820152601960248201527f537461727420696e646578206f7574206f6620626f756e6473000000000000006044820152606490fd5b50346102445760203660031901126102445760043580158015809203610a5c5781906110e561316c565b9061117d575b1561112d5760207f5f0dd08ee0bed9364680dd75e895a44ad73286aa34423ee7531fe7841f4e76d19160ff196009541660ff821617600955604051908152a180f35b60405162461bcd60e51b815260206004820152602260248201527f4d7573742073657420666565207469657273206265666f726520656e61626c696044820152616e6760f01b6064820152608490fd5b5060085415156110eb565b5034610244576020366003190112610244576111a2612d07565b506101a06111b461104c600435612a0d565b6111c16040518092612931565bf35b5034610244576040366003190112610244576111dd61281b565b6024356001600160a01b0381169190829003610a5c576111fb61316c565b611206821515612ef3565b6001600160a01b0316825260208290526040822080546001600160a01b031916909117905580f35b506060366003190112610244576024356044356001600160a01b0381166004358183036119a65761126460ff6003541615612b39565b61126c6135bc565b8315611957578115908161190c575b6112886001548210612b7a565b61129181612a0d565b5093600a85015460ff81161590816118fd575b50156118b8576008850154421015806118aa575b1561185b57838752602087905260408720546112e99187916112e4906001600160a01b03161515612e5e565b612c4e565b9460068501805480155f1461185557506007860154935b838952600660205261131560408a2054612c27565b94848a5260066020528560408b205585600b890155611335600554612c27565b60055588106117865715611591575b6003860180546004880180546005909901805485546001600160a01b03198086163317909655948b168a17909255859055928990556001600160a01b039081169791928a9290911688611499575b5050505f5160206136da5f395f51905f5296608096959493926114307f8c875a887f1bc54cb0fe2951468ca736412e5429164f585257ed15372e1d6b2e93600754988115155f14611487576113f2856113ed84600754612c1a565b612aac565b6007555b60405194859433994291879260a094919796959260c08501988552602085015260408401526060830152600180841b031660808201520152565b0390a46007549080821190828215611479579061144c91612c1a565b600554604051938452602084015260408301526060820152a160015f5160206136fa5f395f51905f525580f35b61148291612c1a565b61144c565b611491858b612aac565b6007556113f6565b93979695949380611548575081808092895af16114b4612eb4565b5015611504575f5160206136da5f395f51905f52966080966114307f8c875a887f1bc54cb0fe2951468ca736412e5429164f585257ed15372e1d6b2e935b93508a92995081949596979850611392565b606460405162461bcd60e51b815260206004820152602060248201527f4661696c656420746f20726566756e642070726576696f7573206269646465726044820152fd5b5f5160206136da5f395f51905f529992507f8c875a887f1bc54cb0fe2951468ca736412e5429164f585257ed15372e1d6b2e9361158c60809a938a611430946135f4565b6114f2565b3461174157604051636eb1769f60e11b8152336004820152306024820152602081604481895afa80156116bc5783918a9161170c575b50106116c7576040516370a0823160e01b8152336004820152602081602481895afa80156116bc5783918a91611683575b501061163e576040516323b872dd60e01b6020820152336024820152306044820152606480820184905281526116399061163360848261286a565b86613661565b611344565b60405162461bcd60e51b815260206004820152601a60248201527f496e73756666696369656e7420746f6b656e2062616c616e63650000000000006044820152606490fd5b9150506020813d6020116116b4575b8161169f6020938361286a565b810103126116b0578290515f6115f8565b5f80fd5b3d9150611692565b6040513d8b823e3d90fd5b60405162461bcd60e51b815260206004820152601c60248201527f496e73756666696369656e7420746f6b656e20616c6c6f77616e6365000000006044820152606490fd5b9150506020813d602011611739575b816117286020938361286a565b810103126116b0578290515f6115c7565b3d915061171b565b60405162461bcd60e51b815260206004820152601d60248201527f43616e6e6f742073656e642045544820666f72204552433230206269640000006044820152606490fd5b50600386015490546040805193845260208401959095524294830194909452606082018790526001600160a01b0316608082015260a081019290925233917fe814151408fcc11f8069f77262418f2fdd23e91da33bd56d4c27c2d959521f599060c090a460405162461bcd60e51b815260206004820152603960248201527f4269642076616c7565206d7573742062652067726561746572207468616e207460448201527f6865206d696e696d756d2072657175697265642076616c7565000000000000006064820152608490fd5b93611300565b60405162461bcd60e51b815260206004820152602160248201527f41756374696f6e20686173206e6f742073746172746564206f7220656e6465646044820152600160fd1b6064820152608490fd5b5060098501544211156112b8565b60405162461bcd60e51b815260206004820152601e60248201527f41756374696f6e2068617320656e646564206f722063616e63656c6c656400006044820152606490fd5b60ff915060081c16155f6112a4565b34851461127b57606460405162461bcd60e51b815260206004820152602060248201527f45544820616d6f756e74206d757374206d617463682062696420616d6f756e746044820152fd5b60405162461bcd60e51b815260206004820152602160248201527f42696420616d6f756e74206d7573742062652067726561746572207468616e206044820152600360fc1b6064820152608490fd5b8480fd5b50346102445780600319360112610244576119c361316c565b5f5160206136ba5f395f51905f5280546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b50346102445780600319360112610244576020600854604051908152f35b50346102445780600319360112610244575f51602061371a5f395f51905f525460ff8160401c16908115611b3d575b50611b2e575f51602061371a5f395f51905f52805468ffffffffffffffffff1916680100000000000000021790555f5160206136ba5f395f51905f5254611aba906001600160a01b0316611ab2613636565b61023c613636565b611ac2613636565b611aca613636565b60015f5160206136fa5f395f51905f525560ff60401b195f51602061371a5f395f51905f5254165f51602061371a5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160028152a180f35b63f92ee8a960e01b8152600490fd5b6002915067ffffffffffffffff1610155f611a60565b5034610244578060031936011261024457602060ff600354166040519015158152f35b50346102445760203660031901126102445760043590600154821015610244576101a0611ba283612a0d565b50805460018201546002830154600384015460048501546005860154600687015460078801546008808a015460098b0154600a8c0154600b909c0154604080516001600160a01b039d8e168152602081019c909c52998c16998b0199909952968a1660608a015294909816608088015260a087019290925260c086015260e085015261010084015261012083015260ff84811615156101408401529390921c9092161515610160830152610180820152f35b503461024457806003193601126102445760408051611c73828261286a565b6005815260208101640322e302e360dc1b81528251938492602084525180928160208601528585015e828201840152601f01601f19168101030190f35b5034610244576020366003190112610244576001600160a01b03611cd261281b565b168152600b60205260408120805490611cea82612d66565b925b828110611d01576040518061103086826129c9565b80611d2061104c611d1460019486612a97565b90549060031b1c612a0d565b611d2a8287612ad9565b52611d358186612ad9565b5001611cec565b5034610244576040366003190112610244576020610367611d5b61281b565b60243590612c4e565b50346116b05760203660031901126116b057600435611d8860ff6003541615612b39565b611d906135bc565b611d9d6001548210612b7a565b611da681612a0d565b5060028101546001600160a01b031633819003611f7b57600a8201805460ff81161580611f6c575b611dd790612bbf565b60038401546001600160a01b0316611f275760048401546001600160a01b0316611ed4576101019061ffff19161790556001808060a01b03835416920154823b156116b0576040516323b872dd60e01b81523060048201526001600160a01b039290921660248301526044820152905f908290606490829084905af18015611ec957611eb4575b50604051908282528260208301528260408301527f10ec610e9b6f8628e9a57f51bfda4adcb5db4c761791045d0f85214b681c7c4860603393a460015f5160206136fa5f395f51905f5255602060405160018152f35b611ec19192505f9061286a565b5f905f611e5e565b6040513d5f823e3d90fd5b60405162461bcd60e51b815260206004820152602560248201527f41756374696f6e2068617320746f6b656e20626964732c2063616e6e6f742063604482015264185b98d95b60da1b6064820152608490fd5b60405162461bcd60e51b815260206004820152601f60248201527f41756374696f6e2068617320626964732c2063616e6e6f742063616e63656c006044820152606490fd5b50600881901c60ff1615611dce565b60405162461bcd60e51b815260206004820152601e60248201527f4f6e6c792073656c6c65722063616e2063616e63656c2061756374696f6e00006044820152606490fd5b346116b0575f3660031901126116b057611fd861316c565b60035460ff8116156120175760ff19166003557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b60405162461bcd60e51b815260206004820152600a602482015269139bdd081c185d5cd95960b21b6044820152606490fd5b346116b0575f3660031901126116b0576008546120658161288c565b90612073604051928361286a565b80825260208201908160085f527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee35f915b8383106120fb578486604051918291602083019060208452518091526040830191905f5b8181106120d6575050500390f35b82518051855260209081015181860152869550604090940193909201916001016120c8565b6002602060019260405161210e81612831565b8554815284860154838201528152019201920191906120a4565b346116b0575f3660031901126116b0576004546005549060025460ff600354165f915f600154905b81811061217a57505060a09460405194855260208501526040840152151560608301526080820152f35b60ff600a61218783612a0d565b5001541615806121d0575b806121b9575b6121a5575b600101612150565b936121b1600191612c27565b94905061219d565b5060096121c582612a0d565b500154421115612198565b5060086121dc82612a0d565b500154421015612192565b346116b05760203660031901126116b05760043561220361316c565b61221260ff6003541615612b39565b61221f6001548210612b7a565b61222881612a0d565b50600a81019182549260ff841615806122c9575b61224590612bbf565b6009830154603b1981019081116122b557421061227057600161226e9460ff191617905561319f565b005b60405162461bcd60e51b815260206004820152601960248201527f41756374696f6e20686173206e6f7420656e64656420796574000000000000006044820152606490fd5b634e487b7160e01b5f52601160045260245ffd5b50600884901c60ff161561223c565b346116b0575f3660031901126116b0576020600254604051908152f35b346116b05760203660031901126116b0576004356008548110156116b05761231e604091612901565b506001815491015482519182526020820152f35b346116b05760403660031901126116b05760043567ffffffffffffffff81116116b0576123639036906004016128a4565b60243567ffffffffffffffff81116116b0576123839036906004016128a4565b61238b61316c565b8051825190600182018092116122b5570361268c5781511561264757600a825111612602576008545f60085580612587575b506123d56103e86123cd83612acc565b511115612aed565b6123de81612acc565b51600a555f90815b835183101561254d576123f98385612ad9565b5111156124fa57600182018083116122b55761241b6103e86123cd8385612ad9565b6124258385612ad9565b516124308284612ad9565b516040519161243e83612831565b825260208201908152600854600160401b8110156124e6578060016124669201600855612901565b9390936124d357859360016040937f050378e8543487fde7e7bd39a32579544cecfec6b7bd3f81082141eb9c255c4895518355519101556124b26124aa8589612ad9565b519186612ad9565b5182519182526020820152a260016124ca8385612ad9565b519201916123e6565b634e487b7160e01b5f525f60045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b60405162461bcd60e51b815260206004820152602560248201527f5468726573686f6c6473206d75737420626520696e20617363656e64696e672060448201526437b93232b960d91b6064820152608490fd5b600160ff1960095416176009557f5f0dd08ee0bed9364680dd75e895a44ad73286aa34423ee7531fe7841f4e76d1602060405160018152a1005b600181901b906001600160ff1b038116036122b55760085f525f5b8181106125af57506123bd565b805f6002927ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee301555f7ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee4820155016125a2565b60405162461bcd60e51b815260206004820152601760248201527f546f6f206d616e7920746965727320286d6178203130290000000000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601a60248201527f4174206c65617374206f6e6520746965722072657175697265640000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152603a60248201527f466565207261746573206172726179206d7573742068617665206f6e65206d6f60448201527f726520656c656d656e74207468616e207468726573686f6c64730000000000006064820152608490fd5b346116b05760203660031901126116b0576001600160a01b0361271861281b565b165f52600b602052602060405f2054604051908152f35b346116b0575f3660031901126116b0576020600454604051908152f35b346116b05760203660031901126116b05760043561276861316c565b6103e881116127aa5760407fd347e206f25a89b917fc9482f1a2d294d749baa4dc9bde7fb495ee11fe49164391600254908060025582519182526020820152a1005b60405162461bcd60e51b81526020600482015260156024820152744665652063616e6e6f74206578636565642031302560581b6044820152606490fd5b346116b0575f3660031901126116b05760809060045460055460025490600754928452602084015260408301526060820152f35b600435906001600160a01b03821682036116b057565b6040810190811067ffffffffffffffff8211176124e657604052565b6101a0810190811067ffffffffffffffff8211176124e657604052565b90601f8019910116810190811067ffffffffffffffff8211176124e657604052565b67ffffffffffffffff81116124e65760051b60200190565b9080601f830112156116b05781356128bb8161288c565b926128c9604051948561286a565b81845260208085019260051b8201019283116116b057602001905b8282106128f15750505090565b81358152602091820191016128e4565b60085481101561291d5760085f5260205f209060011b01905f90565b634e487b7160e01b5f52603260045260245ffd5b80516001600160a01b039081168352602080830151908401526040808301518216908401526060808301518216908401526080808301519091169083015260a0808201519083015260c0808201519083015260e0808201519083015261010080820151908301526101208082015190830152610140808201511515908301526101608082015115159083015261018090810151910152565b60206040818301928281528451809452019201905f5b8181106129ec5750505090565b90919260206101a082612a026001948851612931565b0194019291016129df565b60015481101561291d5760015f52600c60205f20910201905f90565b9080601f830112156116b057813590612a418261288c565b92612a4f604051948561286a565b82845260208085019360051b8201019182116116b057602001915b818310612a775750505090565b82356001600160a01b03811681036116b057815260209283019201612a6a565b805482101561291d575f5260205f2001905f90565b919082018092116122b557565b818102929181159184041417156122b557565b80511561291d5760200190565b805182101561291d5760209160051b010190565b15612af457565b60405162461bcd60e51b815260206004820152601a60248201527f46656520726174652063616e6e6f7420657863656564203130250000000000006044820152606490fd5b15612b4057565b60405162461bcd60e51b815260206004820152601260248201527110dbdb9d1c9858dd081a5cc81c185d5cd95960721b6044820152606490fd5b15612b8157565b60405162461bcd60e51b8152602060048201526016602482015275105d58dd1a5bdb88191bd95cc81b9bdd08195e1a5cdd60521b6044820152606490fd5b15612bc657565b60405162461bcd60e51b815260206004820152602660248201527f41756374696f6e2068617320616c726561647920656e646564206f722063616e60448201526518d95b1b195960d21b6064820152608490fd5b919082039182116122b557565b5f1981146122b55760010190565b908160209103126116b0575160ff811681036116b05790565b612c7660ff91612c5d81612fcd565b50906001600160a01b031680612ca25750601293612ab9565b9116604d81116122b557600a0a908115612c8e570490565b634e487b7160e01b5f52601260045260245ffd5b60206004916040519283809263313ce56760e01b82525afa5f9181612cd6575b50612cd05750601293612ab9565b93612ab9565b612cf991925060203d602011612d00575b612cf1818361286a565b810190612c35565b905f612cc2565b503d612ce7565b60405190612d148261284d565b5f610180838281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e0820152826101008201528261012082015282610140820152826101608201520152565b90612d708261288c565b612d7d604051918261286a565b8281528092612d8e601f199161288c565b01905f5b828110612d9e57505050565b602090612da9612d07565b82828501015201612d92565b90604051612dc28161284d565b82546001600160a01b039081168252600184015460208301526002840154811660408301526003840154811660608301526004840154166080820152600583015460a0820152600683015460c0820152600783015460e08201526008808401546101008301526009840154610120830152600a84015460ff8082161515610140850152911c161515610160820152600b90920154610180830152565b15612e6557565b60405162461bcd60e51b815260206004820152602160248201527f50726963652066656564206e6f742073657420666f72207468697320746f6b656044820152603760f91b6064820152608490fd5b3d15612eee573d9067ffffffffffffffff82116124e65760405191612ee3601f8201601f19166020018461286a565b82523d5f602084013e565b606090565b15612efa57565b60405162461bcd60e51b815260206004820152601a60248201527f496e76616c6964207072696365206665656420616464726573730000000000006044820152606490fd5b9060ff60095416158015612fac575b612fa557600854805b612f635750600a549150565b5f1981018181116122b557612f7781612901565b5054841015612f90575080156122b5575f190180612f57565b6001929350612f9f9150612901565b50015490565b6002549150565b5060085415612f4e565b519069ffffffffffffffffffff821682036116b057565b6001600160a01b039081165f9081526020819052604090205416612ff2811515612e5e565b604051633fabe5a360e21b81529060a082600481845afa918215611ec9575f926130ac575b505f8213156130675760206004916040519283809263313ce56760e01b82525afa908115611ec9575f9161304a57509091565b613063915060203d602011612d0057612cf1818361286a565b9091565b60405162461bcd60e51b815260206004820152601960248201527f496e76616c69642070726963652066726f6d206f7261636c65000000000000006044820152606490fd5b90915060a0813d60a0116130f3575b816130c860a0938361286a565b810103126116b0576130d981612fb6565b506130eb608060208301519201612fb6565b50905f613017565b3d91506130bb565b6001600160a01b03168015613159575f5160206136ba5f395f51905f5280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b631e4fbdf760e01b5f525f60045260245ffd5b5f5160206136ba5f395f51905f52546001600160a01b0316330361318c57565b63118cdaa760e01b5f523360045260245ffd5b81546006830180545f94926001600160a01b03169080613579575b506003830180549091906001600160a01b0316156134b5578154600185015491906001600160a01b0316813b156134b1576040516323b872dd60e01b81523060048201526001600160a01b03919091166024820152604481019290925286908290606490829084905af180156134a657908691613491575b5090600584019586549260ff6009541680613486575b1561347e576132578554612f3f565b8061345d575b506004860180549094906001600160a01b0316806133e45750600287015483918291829182916001600160a01b03165af1613296612eb4565b50156133945780613313575b5050907f4805050c098aacbeda24d1234b31e570f6a2236ed167fd9516d8e87e926883999360a093925b600180861b039054169654926007600180871b0360028401541693600180881b039054169154920154926040519485526020850152604084015260608301526080820152a3565b5f5160206136ba5f395f51905f525492949392829182918291906001600160a01b03165af1613340612eb4565b501561334f5790915f806132a2565b60405162461bcd60e51b815260206004820152601f60248201527f4661696c656420746f207472616e736665722066656520746f206f776e6572006044820152606490fd5b60405162461bcd60e51b815260206004820152602260248201527f4661696c656420746f207472616e736665722066756e647320746f2073656c6c60448201526132b960f11b6064820152608490fd5b60028801547f4805050c098aacbeda24d1234b31e570f6a2236ed167fd9516d8e87e926883999860a0989796909550939192613429916001600160a01b0316846135f4565b80613436575b50506132cc565b61345691600180891b035f5160206136ba5f395f51905f525416906135f4565b5f8061342f565b613477915061346f6127109186612ab9565b048094612c1a565b925f61325d565b600254613257565b506008541515613248565b8161349b9161286a565b6119a657845f613232565b6040513d88823e3d90fd5b8780fd5b91505060028201916001808060a01b03845416910154823b156116b0576040516323b872dd60e01b81523060048201526001600160a01b039290921660248301526044820152905f908290606490829084905af18015611ec957613563575b5060a07f4805050c098aacbeda24d1234b31e570f6a2236ed167fd9516d8e87e9268839991600180831b03905416604051908582526020820152846040820152846060820152846080820152a3565b6135709193505f9061286a565b5f9160a0613514565b60806135955f5160206136da5f395f51905f5292600754612c1a565b80600755600554855490604051928352602083015260408201525f6060820152a15f6131ba565b60025f5160206136fa5f395f51905f5254146135e55760025f5160206136fa5f395f51905f5255565b633ee5aeb560e01b5f5260045ffd5b60405163a9059cbb60e01b60208201526001600160a01b0390921660248301526044808301939093529181526136349161362f60648361286a565b613661565b565b60ff5f51602061371a5f395f51905f525460401c161561365257565b631afcd79f60e31b5f5260045ffd5b905f602091828151910182855af115611ec9575f513d6136b057506001600160a01b0381163b155b6136905750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b6001141561368956fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300fe9b183fcbd8416e8d2ddcc6180ff8e7a33c363b7191deccac34f694c4faab299b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220cc3f6a0472fde69dba7add4502072058a85338a25ac8ae85a55ceacd5ca2815564736f6c63430008200033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// MyXAuctionV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MyXAuctionV2MetaData.ABI instead.
var MyXAuctionV2ABI = MyXAuctionV2MetaData.ABI

// MyXAuctionV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyXAuctionV2MetaData.Bin instead.
var MyXAuctionV2Bin = MyXAuctionV2MetaData.Bin

// DeployMyXAuctionV2 deploys a new Ethereum contract, binding an instance of MyXAuctionV2 to it.
func DeployMyXAuctionV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyXAuctionV2, error) {
	parsed, err := MyXAuctionV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyXAuctionV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyXAuctionV2{MyXAuctionV2Caller: MyXAuctionV2Caller{contract: contract}, MyXAuctionV2Transactor: MyXAuctionV2Transactor{contract: contract}, MyXAuctionV2Filterer: MyXAuctionV2Filterer{contract: contract}}, nil
}

// MyXAuctionV2 is an auto generated Go binding around an Ethereum contract.
type MyXAuctionV2 struct {
	MyXAuctionV2Caller     // Read-only binding to the contract
	MyXAuctionV2Transactor // Write-only binding to the contract
	MyXAuctionV2Filterer   // Log filterer for contract events
}

// MyXAuctionV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MyXAuctionV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyXAuctionV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MyXAuctionV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyXAuctionV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyXAuctionV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyXAuctionV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyXAuctionV2Session struct {
	Contract     *MyXAuctionV2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyXAuctionV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyXAuctionV2CallerSession struct {
	Contract *MyXAuctionV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MyXAuctionV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyXAuctionV2TransactorSession struct {
	Contract     *MyXAuctionV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MyXAuctionV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MyXAuctionV2Raw struct {
	Contract *MyXAuctionV2 // Generic contract binding to access the raw methods on
}

// MyXAuctionV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyXAuctionV2CallerRaw struct {
	Contract *MyXAuctionV2Caller // Generic read-only contract binding to access the raw methods on
}

// MyXAuctionV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyXAuctionV2TransactorRaw struct {
	Contract *MyXAuctionV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMyXAuctionV2 creates a new instance of MyXAuctionV2, bound to a specific deployed contract.
func NewMyXAuctionV2(address common.Address, backend bind.ContractBackend) (*MyXAuctionV2, error) {
	contract, err := bindMyXAuctionV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2{MyXAuctionV2Caller: MyXAuctionV2Caller{contract: contract}, MyXAuctionV2Transactor: MyXAuctionV2Transactor{contract: contract}, MyXAuctionV2Filterer: MyXAuctionV2Filterer{contract: contract}}, nil
}

// NewMyXAuctionV2Caller creates a new read-only instance of MyXAuctionV2, bound to a specific deployed contract.
func NewMyXAuctionV2Caller(address common.Address, caller bind.ContractCaller) (*MyXAuctionV2Caller, error) {
	contract, err := bindMyXAuctionV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2Caller{contract: contract}, nil
}

// NewMyXAuctionV2Transactor creates a new write-only instance of MyXAuctionV2, bound to a specific deployed contract.
func NewMyXAuctionV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MyXAuctionV2Transactor, error) {
	contract, err := bindMyXAuctionV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2Transactor{contract: contract}, nil
}

// NewMyXAuctionV2Filterer creates a new log filterer instance of MyXAuctionV2, bound to a specific deployed contract.
func NewMyXAuctionV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MyXAuctionV2Filterer, error) {
	contract, err := bindMyXAuctionV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2Filterer{contract: contract}, nil
}

// bindMyXAuctionV2 binds a generic wrapper to an already deployed contract.
func bindMyXAuctionV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyXAuctionV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyXAuctionV2 *MyXAuctionV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyXAuctionV2.Contract.MyXAuctionV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyXAuctionV2 *MyXAuctionV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.MyXAuctionV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyXAuctionV2 *MyXAuctionV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.MyXAuctionV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyXAuctionV2 *MyXAuctionV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyXAuctionV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyXAuctionV2 *MyXAuctionV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyXAuctionV2 *MyXAuctionV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.contract.Transact(opts, method, params...)
}

// AuctionBidCount is a free data retrieval call binding the contract method 0xcb600160.
//
// Solidity: function auctionBidCount(uint256 ) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) AuctionBidCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "auctionBidCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionBidCount is a free data retrieval call binding the contract method 0xcb600160.
//
// Solidity: function auctionBidCount(uint256 ) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) AuctionBidCount(arg0 *big.Int) (*big.Int, error) {
	return _MyXAuctionV2.Contract.AuctionBidCount(&_MyXAuctionV2.CallOpts, arg0)
}

// AuctionBidCount is a free data retrieval call binding the contract method 0xcb600160.
//
// Solidity: function auctionBidCount(uint256 ) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) AuctionBidCount(arg0 *big.Int) (*big.Int, error) {
	return _MyXAuctionV2.Contract.AuctionBidCount(&_MyXAuctionV2.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address nftAddress, uint256 tokenId, address seller, address highestBidder, address highestBidToken, uint256 highestBid, uint256 highestBidValue, uint256 startPrice, uint256 startTime, uint256 endTime, bool ended, bool cancelled, uint256 bidCount)
func (_MyXAuctionV2 *MyXAuctionV2Caller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NftAddress      common.Address
	TokenId         *big.Int
	Seller          common.Address
	HighestBidder   common.Address
	HighestBidToken common.Address
	HighestBid      *big.Int
	HighestBidValue *big.Int
	StartPrice      *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	Ended           bool
	Cancelled       bool
	BidCount        *big.Int
}, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		NftAddress      common.Address
		TokenId         *big.Int
		Seller          common.Address
		HighestBidder   common.Address
		HighestBidToken common.Address
		HighestBid      *big.Int
		HighestBidValue *big.Int
		StartPrice      *big.Int
		StartTime       *big.Int
		EndTime         *big.Int
		Ended           bool
		Cancelled       bool
		BidCount        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NftAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.HighestBidder = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.HighestBidToken = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.HighestBid = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.HighestBidValue = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.StartPrice = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Ended = *abi.ConvertType(out[10], new(bool)).(*bool)
	outstruct.Cancelled = *abi.ConvertType(out[11], new(bool)).(*bool)
	outstruct.BidCount = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address nftAddress, uint256 tokenId, address seller, address highestBidder, address highestBidToken, uint256 highestBid, uint256 highestBidValue, uint256 startPrice, uint256 startTime, uint256 endTime, bool ended, bool cancelled, uint256 bidCount)
func (_MyXAuctionV2 *MyXAuctionV2Session) Auctions(arg0 *big.Int) (struct {
	NftAddress      common.Address
	TokenId         *big.Int
	Seller          common.Address
	HighestBidder   common.Address
	HighestBidToken common.Address
	HighestBid      *big.Int
	HighestBidValue *big.Int
	StartPrice      *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	Ended           bool
	Cancelled       bool
	BidCount        *big.Int
}, error) {
	return _MyXAuctionV2.Contract.Auctions(&_MyXAuctionV2.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address nftAddress, uint256 tokenId, address seller, address highestBidder, address highestBidToken, uint256 highestBid, uint256 highestBidValue, uint256 startPrice, uint256 startTime, uint256 endTime, bool ended, bool cancelled, uint256 bidCount)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) Auctions(arg0 *big.Int) (struct {
	NftAddress      common.Address
	TokenId         *big.Int
	Seller          common.Address
	HighestBidder   common.Address
	HighestBidToken common.Address
	HighestBid      *big.Int
	HighestBidValue *big.Int
	StartPrice      *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	Ended           bool
	Cancelled       bool
	BidCount        *big.Int
}, error) {
	return _MyXAuctionV2.Contract.Auctions(&_MyXAuctionV2.CallOpts, arg0)
}

// BaseFeeRate is a free data retrieval call binding the contract method 0xcf2d2178.
//
// Solidity: function baseFeeRate() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) BaseFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "baseFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseFeeRate is a free data retrieval call binding the contract method 0xcf2d2178.
//
// Solidity: function baseFeeRate() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) BaseFeeRate() (*big.Int, error) {
	return _MyXAuctionV2.Contract.BaseFeeRate(&_MyXAuctionV2.CallOpts)
}

// BaseFeeRate is a free data retrieval call binding the contract method 0xcf2d2178.
//
// Solidity: function baseFeeRate() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) BaseFeeRate() (*big.Int, error) {
	return _MyXAuctionV2.Contract.BaseFeeRate(&_MyXAuctionV2.CallOpts)
}

// CalculateDynamicFeeRate is a free data retrieval call binding the contract method 0xb6304c56.
//
// Solidity: function calculateDynamicFeeRate(uint256 _usdValue) view returns(uint256 feeRate)
func (_MyXAuctionV2 *MyXAuctionV2Caller) CalculateDynamicFeeRate(opts *bind.CallOpts, _usdValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "calculateDynamicFeeRate", _usdValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateDynamicFeeRate is a free data retrieval call binding the contract method 0xb6304c56.
//
// Solidity: function calculateDynamicFeeRate(uint256 _usdValue) view returns(uint256 feeRate)
func (_MyXAuctionV2 *MyXAuctionV2Session) CalculateDynamicFeeRate(_usdValue *big.Int) (*big.Int, error) {
	return _MyXAuctionV2.Contract.CalculateDynamicFeeRate(&_MyXAuctionV2.CallOpts, _usdValue)
}

// CalculateDynamicFeeRate is a free data retrieval call binding the contract method 0xb6304c56.
//
// Solidity: function calculateDynamicFeeRate(uint256 _usdValue) view returns(uint256 feeRate)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) CalculateDynamicFeeRate(_usdValue *big.Int) (*big.Int, error) {
	return _MyXAuctionV2.Contract.CalculateDynamicFeeRate(&_MyXAuctionV2.CallOpts, _usdValue)
}

// ConvertToUSDValue is a free data retrieval call binding the contract method 0x43638e86.
//
// Solidity: function convertToUSDValue(address _token, uint256 _amount) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) ConvertToUSDValue(opts *bind.CallOpts, _token common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "convertToUSDValue", _token, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToUSDValue is a free data retrieval call binding the contract method 0x43638e86.
//
// Solidity: function convertToUSDValue(address _token, uint256 _amount) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) ConvertToUSDValue(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _MyXAuctionV2.Contract.ConvertToUSDValue(&_MyXAuctionV2.CallOpts, _token, _amount)
}

// ConvertToUSDValue is a free data retrieval call binding the contract method 0x43638e86.
//
// Solidity: function convertToUSDValue(address _token, uint256 _amount) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) ConvertToUSDValue(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _MyXAuctionV2.Contract.ConvertToUSDValue(&_MyXAuctionV2.CallOpts, _token, _amount)
}

// FeeTiers is a free data retrieval call binding the contract method 0x230ed44a.
//
// Solidity: function feeTiers(uint256 ) view returns(uint256 threshold, uint256 feeRate)
func (_MyXAuctionV2 *MyXAuctionV2Caller) FeeTiers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Threshold *big.Int
	FeeRate   *big.Int
}, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "feeTiers", arg0)

	outstruct := new(struct {
		Threshold *big.Int
		FeeRate   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Threshold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeeRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FeeTiers is a free data retrieval call binding the contract method 0x230ed44a.
//
// Solidity: function feeTiers(uint256 ) view returns(uint256 threshold, uint256 feeRate)
func (_MyXAuctionV2 *MyXAuctionV2Session) FeeTiers(arg0 *big.Int) (struct {
	Threshold *big.Int
	FeeRate   *big.Int
}, error) {
	return _MyXAuctionV2.Contract.FeeTiers(&_MyXAuctionV2.CallOpts, arg0)
}

// FeeTiers is a free data retrieval call binding the contract method 0x230ed44a.
//
// Solidity: function feeTiers(uint256 ) view returns(uint256 threshold, uint256 feeRate)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) FeeTiers(arg0 *big.Int) (struct {
	Threshold *big.Int
	FeeRate   *big.Int
}, error) {
	return _MyXAuctionV2.Contract.FeeTiers(&_MyXAuctionV2.CallOpts, arg0)
}

// GetAllFeeTiers is a free data retrieval call binding the contract method 0x387296d3.
//
// Solidity: function getAllFeeTiers() view returns((uint256,uint256)[])
func (_MyXAuctionV2 *MyXAuctionV2Caller) GetAllFeeTiers(opts *bind.CallOpts) ([]MyXAuctionV2FeeTier, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "getAllFeeTiers")

	if err != nil {
		return *new([]MyXAuctionV2FeeTier), err
	}

	out0 := *abi.ConvertType(out[0], new([]MyXAuctionV2FeeTier)).(*[]MyXAuctionV2FeeTier)

	return out0, err

}

// GetAllFeeTiers is a free data retrieval call binding the contract method 0x387296d3.
//
// Solidity: function getAllFeeTiers() view returns((uint256,uint256)[])
func (_MyXAuctionV2 *MyXAuctionV2Session) GetAllFeeTiers() ([]MyXAuctionV2FeeTier, error) {
	return _MyXAuctionV2.Contract.GetAllFeeTiers(&_MyXAuctionV2.CallOpts)
}

// GetAllFeeTiers is a free data retrieval call binding the contract method 0x387296d3.
//
// Solidity: function getAllFeeTiers() view returns((uint256,uint256)[])
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) GetAllFeeTiers() ([]MyXAuctionV2FeeTier, error) {
	return _MyXAuctionV2.Contract.GetAllFeeTiers(&_MyXAuctionV2.CallOpts)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _auctionId) view returns((address,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,bool,bool,uint256))
func (_MyXAuctionV2 *MyXAuctionV2Caller) GetAuction(opts *bind.CallOpts, _auctionId *big.Int) (MyXAuctionV2Auction, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "getAuction", _auctionId)

	if err != nil {
		return *new(MyXAuctionV2Auction), err
	}

	out0 := *abi.ConvertType(out[0], new(MyXAuctionV2Auction)).(*MyXAuctionV2Auction)

	return out0, err

}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _auctionId) view returns((address,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,bool,bool,uint256))
func (_MyXAuctionV2 *MyXAuctionV2Session) GetAuction(_auctionId *big.Int) (MyXAuctionV2Auction, error) {
	return _MyXAuctionV2.Contract.GetAuction(&_MyXAuctionV2.CallOpts, _auctionId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _auctionId) view returns((address,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,bool,bool,uint256))
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) GetAuction(_auctionId *big.Int) (MyXAuctionV2Auction, error) {
	return _MyXAuctionV2.Contract.GetAuction(&_MyXAuctionV2.CallOpts, _auctionId)
}

// GetAuctionCount is a free data retrieval call binding the contract method 0xc44e6640.
//
// Solidity: function getAuctionCount() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) GetAuctionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "getAuctionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAuctionCount is a free data retrieval call binding the contract method 0xc44e6640.
//
// Solidity: function getAuctionCount() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) GetAuctionCount() (*big.Int, error) {
	return _MyXAuctionV2.Contract.GetAuctionCount(&_MyXAuctionV2.CallOpts)
}

// GetAuctionCount is a free data retrieval call binding the contract method 0xc44e6640.
//
// Solidity: function getAuctionCount() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) GetAuctionCount() (*big.Int, error) {
	return _MyXAuctionV2.Contract.GetAuctionCount(&_MyXAuctionV2.CallOpts)
}

// GetAuctionSimpleStats is a free data retrieval call binding the contract method 0x0a48a7fe.
//
// Solidity: function getAuctionSimpleStats() view returns(uint256, uint256, uint256, uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) GetAuctionSimpleStats(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "getAuctionSimpleStats")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAuctionSimpleStats is a free data retrieval call binding the contract method 0x0a48a7fe.
//
// Solidity: function getAuctionSimpleStats() view returns(uint256, uint256, uint256, uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) GetAuctionSimpleStats() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MyXAuctionV2.Contract.GetAuctionSimpleStats(&_MyXAuctionV2.CallOpts)
}

// GetAuctionSimpleStats is a free data retrieval call binding the contract method 0x0a48a7fe.
//
// Solidity: function getAuctionSimpleStats() view returns(uint256, uint256, uint256, uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) GetAuctionSimpleStats() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MyXAuctionV2.Contract.GetAuctionSimpleStats(&_MyXAuctionV2.CallOpts)
}

// GetAuctionStats is a free data retrieval call binding the contract method 0x2ecdee5b.
//
// Solidity: function getAuctionStats() view returns(uint256 totalAuctions, uint256 totalBids, uint256 currentPlatformFee, bool isPaused, uint256 activeAuctions)
func (_MyXAuctionV2 *MyXAuctionV2Caller) GetAuctionStats(opts *bind.CallOpts) (struct {
	TotalAuctions      *big.Int
	TotalBids          *big.Int
	CurrentPlatformFee *big.Int
	IsPaused           bool
	ActiveAuctions     *big.Int
}, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "getAuctionStats")

	outstruct := new(struct {
		TotalAuctions      *big.Int
		TotalBids          *big.Int
		CurrentPlatformFee *big.Int
		IsPaused           bool
		ActiveAuctions     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalAuctions = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalBids = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentPlatformFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsPaused = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.ActiveAuctions = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAuctionStats is a free data retrieval call binding the contract method 0x2ecdee5b.
//
// Solidity: function getAuctionStats() view returns(uint256 totalAuctions, uint256 totalBids, uint256 currentPlatformFee, bool isPaused, uint256 activeAuctions)
func (_MyXAuctionV2 *MyXAuctionV2Session) GetAuctionStats() (struct {
	TotalAuctions      *big.Int
	TotalBids          *big.Int
	CurrentPlatformFee *big.Int
	IsPaused           bool
	ActiveAuctions     *big.Int
}, error) {
	return _MyXAuctionV2.Contract.GetAuctionStats(&_MyXAuctionV2.CallOpts)
}

// GetAuctionStats is a free data retrieval call binding the contract method 0x2ecdee5b.
//
// Solidity: function getAuctionStats() view returns(uint256 totalAuctions, uint256 totalBids, uint256 currentPlatformFee, bool isPaused, uint256 activeAuctions)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) GetAuctionStats() (struct {
	TotalAuctions      *big.Int
	TotalBids          *big.Int
	CurrentPlatformFee *big.Int
	IsPaused           bool
	ActiveAuctions     *big.Int
}, error) {
	return _MyXAuctionV2.Contract.GetAuctionStats(&_MyXAuctionV2.CallOpts)
}

// GetAuctionsBatch is a free data retrieval call binding the contract method 0x7a222a4a.
//
// Solidity: function getAuctionsBatch(uint256 _startIndex, uint256 _count) view returns((address,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,bool,bool,uint256)[])
func (_MyXAuctionV2 *MyXAuctionV2Caller) GetAuctionsBatch(opts *bind.CallOpts, _startIndex *big.Int, _count *big.Int) ([]MyXAuctionV2Auction, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "getAuctionsBatch", _startIndex, _count)

	if err != nil {
		return *new([]MyXAuctionV2Auction), err
	}

	out0 := *abi.ConvertType(out[0], new([]MyXAuctionV2Auction)).(*[]MyXAuctionV2Auction)

	return out0, err

}

// GetAuctionsBatch is a free data retrieval call binding the contract method 0x7a222a4a.
//
// Solidity: function getAuctionsBatch(uint256 _startIndex, uint256 _count) view returns((address,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,bool,bool,uint256)[])
func (_MyXAuctionV2 *MyXAuctionV2Session) GetAuctionsBatch(_startIndex *big.Int, _count *big.Int) ([]MyXAuctionV2Auction, error) {
	return _MyXAuctionV2.Contract.GetAuctionsBatch(&_MyXAuctionV2.CallOpts, _startIndex, _count)
}

// GetAuctionsBatch is a free data retrieval call binding the contract method 0x7a222a4a.
//
// Solidity: function getAuctionsBatch(uint256 _startIndex, uint256 _count) view returns((address,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,bool,bool,uint256)[])
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) GetAuctionsBatch(_startIndex *big.Int, _count *big.Int) ([]MyXAuctionV2Auction, error) {
	return _MyXAuctionV2.Contract.GetAuctionsBatch(&_MyXAuctionV2.CallOpts, _startIndex, _count)
}

// GetFeeTierCount is a free data retrieval call binding the contract method 0x6da7c22c.
//
// Solidity: function getFeeTierCount() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) GetFeeTierCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "getFeeTierCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeTierCount is a free data retrieval call binding the contract method 0x6da7c22c.
//
// Solidity: function getFeeTierCount() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) GetFeeTierCount() (*big.Int, error) {
	return _MyXAuctionV2.Contract.GetFeeTierCount(&_MyXAuctionV2.CallOpts)
}

// GetFeeTierCount is a free data retrieval call binding the contract method 0x6da7c22c.
//
// Solidity: function getFeeTierCount() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) GetFeeTierCount() (*big.Int, error) {
	return _MyXAuctionV2.Contract.GetFeeTierCount(&_MyXAuctionV2.CallOpts)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xd02641a0.
//
// Solidity: function getTokenPrice(address _token) view returns(uint256 price, uint8 decimals)
func (_MyXAuctionV2 *MyXAuctionV2Caller) GetTokenPrice(opts *bind.CallOpts, _token common.Address) (struct {
	Price    *big.Int
	Decimals uint8
}, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "getTokenPrice", _token)

	outstruct := new(struct {
		Price    *big.Int
		Decimals uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Decimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetTokenPrice is a free data retrieval call binding the contract method 0xd02641a0.
//
// Solidity: function getTokenPrice(address _token) view returns(uint256 price, uint8 decimals)
func (_MyXAuctionV2 *MyXAuctionV2Session) GetTokenPrice(_token common.Address) (struct {
	Price    *big.Int
	Decimals uint8
}, error) {
	return _MyXAuctionV2.Contract.GetTokenPrice(&_MyXAuctionV2.CallOpts, _token)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xd02641a0.
//
// Solidity: function getTokenPrice(address _token) view returns(uint256 price, uint8 decimals)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) GetTokenPrice(_token common.Address) (struct {
	Price    *big.Int
	Decimals uint8
}, error) {
	return _MyXAuctionV2.Contract.GetTokenPrice(&_MyXAuctionV2.CallOpts, _token)
}

// GetTotalValueLocked is a free data retrieval call binding the contract method 0xb26025aa.
//
// Solidity: function getTotalValueLocked() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) GetTotalValueLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "getTotalValueLocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalValueLocked is a free data retrieval call binding the contract method 0xb26025aa.
//
// Solidity: function getTotalValueLocked() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) GetTotalValueLocked() (*big.Int, error) {
	return _MyXAuctionV2.Contract.GetTotalValueLocked(&_MyXAuctionV2.CallOpts)
}

// GetTotalValueLocked is a free data retrieval call binding the contract method 0xb26025aa.
//
// Solidity: function getTotalValueLocked() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) GetTotalValueLocked() (*big.Int, error) {
	return _MyXAuctionV2.Contract.GetTotalValueLocked(&_MyXAuctionV2.CallOpts)
}

// GetUserCreatedAuctionCount is a free data retrieval call binding the contract method 0x1a8d52a1.
//
// Solidity: function getUserCreatedAuctionCount(address _user) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) GetUserCreatedAuctionCount(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "getUserCreatedAuctionCount", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserCreatedAuctionCount is a free data retrieval call binding the contract method 0x1a8d52a1.
//
// Solidity: function getUserCreatedAuctionCount(address _user) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) GetUserCreatedAuctionCount(_user common.Address) (*big.Int, error) {
	return _MyXAuctionV2.Contract.GetUserCreatedAuctionCount(&_MyXAuctionV2.CallOpts, _user)
}

// GetUserCreatedAuctionCount is a free data retrieval call binding the contract method 0x1a8d52a1.
//
// Solidity: function getUserCreatedAuctionCount(address _user) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) GetUserCreatedAuctionCount(_user common.Address) (*big.Int, error) {
	return _MyXAuctionV2.Contract.GetUserCreatedAuctionCount(&_MyXAuctionV2.CallOpts, _user)
}

// GetUserCreatedAuctions is a free data retrieval call binding the contract method 0x450ed735.
//
// Solidity: function getUserCreatedAuctions(address _user) view returns((address,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,bool,bool,uint256)[])
func (_MyXAuctionV2 *MyXAuctionV2Caller) GetUserCreatedAuctions(opts *bind.CallOpts, _user common.Address) ([]MyXAuctionV2Auction, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "getUserCreatedAuctions", _user)

	if err != nil {
		return *new([]MyXAuctionV2Auction), err
	}

	out0 := *abi.ConvertType(out[0], new([]MyXAuctionV2Auction)).(*[]MyXAuctionV2Auction)

	return out0, err

}

// GetUserCreatedAuctions is a free data retrieval call binding the contract method 0x450ed735.
//
// Solidity: function getUserCreatedAuctions(address _user) view returns((address,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,bool,bool,uint256)[])
func (_MyXAuctionV2 *MyXAuctionV2Session) GetUserCreatedAuctions(_user common.Address) ([]MyXAuctionV2Auction, error) {
	return _MyXAuctionV2.Contract.GetUserCreatedAuctions(&_MyXAuctionV2.CallOpts, _user)
}

// GetUserCreatedAuctions is a free data retrieval call binding the contract method 0x450ed735.
//
// Solidity: function getUserCreatedAuctions(address _user) view returns((address,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,bool,bool,uint256)[])
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) GetUserCreatedAuctions(_user common.Address) ([]MyXAuctionV2Auction, error) {
	return _MyXAuctionV2.Contract.GetUserCreatedAuctions(&_MyXAuctionV2.CallOpts, _user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyXAuctionV2 *MyXAuctionV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyXAuctionV2 *MyXAuctionV2Session) Owner() (common.Address, error) {
	return _MyXAuctionV2.Contract.Owner(&_MyXAuctionV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) Owner() (common.Address, error) {
	return _MyXAuctionV2.Contract.Owner(&_MyXAuctionV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MyXAuctionV2 *MyXAuctionV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MyXAuctionV2 *MyXAuctionV2Session) Paused() (bool, error) {
	return _MyXAuctionV2.Contract.Paused(&_MyXAuctionV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) Paused() (bool, error) {
	return _MyXAuctionV2.Contract.Paused(&_MyXAuctionV2.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) PlatformFee() (*big.Int, error) {
	return _MyXAuctionV2.Contract.PlatformFee(&_MyXAuctionV2.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) PlatformFee() (*big.Int, error) {
	return _MyXAuctionV2.Contract.PlatformFee(&_MyXAuctionV2.CallOpts)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_MyXAuctionV2 *MyXAuctionV2Caller) PriceFeeds(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "priceFeeds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_MyXAuctionV2 *MyXAuctionV2Session) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _MyXAuctionV2.Contract.PriceFeeds(&_MyXAuctionV2.CallOpts, arg0)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _MyXAuctionV2.Contract.PriceFeeds(&_MyXAuctionV2.CallOpts, arg0)
}

// TotalAuctionsCreated is a free data retrieval call binding the contract method 0x18aa08f3.
//
// Solidity: function totalAuctionsCreated() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) TotalAuctionsCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "totalAuctionsCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAuctionsCreated is a free data retrieval call binding the contract method 0x18aa08f3.
//
// Solidity: function totalAuctionsCreated() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) TotalAuctionsCreated() (*big.Int, error) {
	return _MyXAuctionV2.Contract.TotalAuctionsCreated(&_MyXAuctionV2.CallOpts)
}

// TotalAuctionsCreated is a free data retrieval call binding the contract method 0x18aa08f3.
//
// Solidity: function totalAuctionsCreated() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) TotalAuctionsCreated() (*big.Int, error) {
	return _MyXAuctionV2.Contract.TotalAuctionsCreated(&_MyXAuctionV2.CallOpts)
}

// TotalBidsPlaced is a free data retrieval call binding the contract method 0xba8b196e.
//
// Solidity: function totalBidsPlaced() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) TotalBidsPlaced(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "totalBidsPlaced")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBidsPlaced is a free data retrieval call binding the contract method 0xba8b196e.
//
// Solidity: function totalBidsPlaced() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) TotalBidsPlaced() (*big.Int, error) {
	return _MyXAuctionV2.Contract.TotalBidsPlaced(&_MyXAuctionV2.CallOpts)
}

// TotalBidsPlaced is a free data retrieval call binding the contract method 0xba8b196e.
//
// Solidity: function totalBidsPlaced() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) TotalBidsPlaced() (*big.Int, error) {
	return _MyXAuctionV2.Contract.TotalBidsPlaced(&_MyXAuctionV2.CallOpts)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) TotalValueLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "totalValueLocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) TotalValueLocked() (*big.Int, error) {
	return _MyXAuctionV2.Contract.TotalValueLocked(&_MyXAuctionV2.CallOpts)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) TotalValueLocked() (*big.Int, error) {
	return _MyXAuctionV2.Contract.TotalValueLocked(&_MyXAuctionV2.CallOpts)
}

// UseDynamicFee is a free data retrieval call binding the contract method 0xa1127c56.
//
// Solidity: function useDynamicFee() view returns(bool)
func (_MyXAuctionV2 *MyXAuctionV2Caller) UseDynamicFee(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "useDynamicFee")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDynamicFee is a free data retrieval call binding the contract method 0xa1127c56.
//
// Solidity: function useDynamicFee() view returns(bool)
func (_MyXAuctionV2 *MyXAuctionV2Session) UseDynamicFee() (bool, error) {
	return _MyXAuctionV2.Contract.UseDynamicFee(&_MyXAuctionV2.CallOpts)
}

// UseDynamicFee is a free data retrieval call binding the contract method 0xa1127c56.
//
// Solidity: function useDynamicFee() view returns(bool)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) UseDynamicFee() (bool, error) {
	return _MyXAuctionV2.Contract.UseDynamicFee(&_MyXAuctionV2.CallOpts)
}

// UserCreatedAuctionIds is a free data retrieval call binding the contract method 0x9cde2871.
//
// Solidity: function userCreatedAuctionIds(address , uint256 ) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Caller) UserCreatedAuctionIds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "userCreatedAuctionIds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserCreatedAuctionIds is a free data retrieval call binding the contract method 0x9cde2871.
//
// Solidity: function userCreatedAuctionIds(address , uint256 ) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) UserCreatedAuctionIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MyXAuctionV2.Contract.UserCreatedAuctionIds(&_MyXAuctionV2.CallOpts, arg0, arg1)
}

// UserCreatedAuctionIds is a free data retrieval call binding the contract method 0x9cde2871.
//
// Solidity: function userCreatedAuctionIds(address , uint256 ) view returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) UserCreatedAuctionIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MyXAuctionV2.Contract.UserCreatedAuctionIds(&_MyXAuctionV2.CallOpts, arg0, arg1)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_MyXAuctionV2 *MyXAuctionV2Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyXAuctionV2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_MyXAuctionV2 *MyXAuctionV2Session) Version() (string, error) {
	return _MyXAuctionV2.Contract.Version(&_MyXAuctionV2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_MyXAuctionV2 *MyXAuctionV2CallerSession) Version() (string, error) {
	return _MyXAuctionV2.Contract.Version(&_MyXAuctionV2.CallOpts)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0x9600dd83.
//
// Solidity: function approveNFT(address _nftAddress, uint256 _tokenId) returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) ApproveNFT(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "approveNFT", _nftAddress, _tokenId)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0x9600dd83.
//
// Solidity: function approveNFT(address _nftAddress, uint256 _tokenId) returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) ApproveNFT(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.ApproveNFT(&_MyXAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0x9600dd83.
//
// Solidity: function approveNFT(address _nftAddress, uint256 _tokenId) returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) ApproveNFT(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.ApproveNFT(&_MyXAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x742f0a90.
//
// Solidity: function bid(uint256 _auctionId, uint256 _amount, address _token) payable returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) Bid(opts *bind.TransactOpts, _auctionId *big.Int, _amount *big.Int, _token common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "bid", _auctionId, _amount, _token)
}

// Bid is a paid mutator transaction binding the contract method 0x742f0a90.
//
// Solidity: function bid(uint256 _auctionId, uint256 _amount, address _token) payable returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) Bid(_auctionId *big.Int, _amount *big.Int, _token common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.Bid(&_MyXAuctionV2.TransactOpts, _auctionId, _amount, _token)
}

// Bid is a paid mutator transaction binding the contract method 0x742f0a90.
//
// Solidity: function bid(uint256 _auctionId, uint256 _amount, address _token) payable returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) Bid(_auctionId *big.Int, _amount *big.Int, _token common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.Bid(&_MyXAuctionV2.TransactOpts, _auctionId, _amount, _token)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _auctionId) returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) CancelAuction(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "cancelAuction", _auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _auctionId) returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) CancelAuction(_auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.CancelAuction(&_MyXAuctionV2.TransactOpts, _auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _auctionId) returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) CancelAuction(_auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.CancelAuction(&_MyXAuctionV2.TransactOpts, _auctionId)
}

// CancelNFTApproval is a paid mutator transaction binding the contract method 0xdcd4325d.
//
// Solidity: function cancelNFTApproval(address _nftAddress, uint256 _tokenId) returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) CancelNFTApproval(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "cancelNFTApproval", _nftAddress, _tokenId)
}

// CancelNFTApproval is a paid mutator transaction binding the contract method 0xdcd4325d.
//
// Solidity: function cancelNFTApproval(address _nftAddress, uint256 _tokenId) returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) CancelNFTApproval(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.CancelNFTApproval(&_MyXAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// CancelNFTApproval is a paid mutator transaction binding the contract method 0xdcd4325d.
//
// Solidity: function cancelNFTApproval(address _nftAddress, uint256 _tokenId) returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) CancelNFTApproval(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.CancelNFTApproval(&_MyXAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// CancelUserAuction is a paid mutator transaction binding the contract method 0x3fe3b2ce.
//
// Solidity: function cancelUserAuction(uint256 _auctionId) returns(bool)
func (_MyXAuctionV2 *MyXAuctionV2Transactor) CancelUserAuction(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "cancelUserAuction", _auctionId)
}

// CancelUserAuction is a paid mutator transaction binding the contract method 0x3fe3b2ce.
//
// Solidity: function cancelUserAuction(uint256 _auctionId) returns(bool)
func (_MyXAuctionV2 *MyXAuctionV2Session) CancelUserAuction(_auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.CancelUserAuction(&_MyXAuctionV2.TransactOpts, _auctionId)
}

// CancelUserAuction is a paid mutator transaction binding the contract method 0x3fe3b2ce.
//
// Solidity: function cancelUserAuction(uint256 _auctionId) returns(bool)
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) CancelUserAuction(_auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.CancelUserAuction(&_MyXAuctionV2.TransactOpts, _auctionId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, uint256 _startPrice, uint256 _startTime, uint256 _endTime) returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Transactor) CreateAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _startPrice *big.Int, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "createAuction", _nftAddress, _tokenId, _startPrice, _startTime, _endTime)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, uint256 _startPrice, uint256 _startTime, uint256 _endTime) returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2Session) CreateAuction(_nftAddress common.Address, _tokenId *big.Int, _startPrice *big.Int, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.CreateAuction(&_MyXAuctionV2.TransactOpts, _nftAddress, _tokenId, _startPrice, _startTime, _endTime)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, uint256 _startPrice, uint256 _startTime, uint256 _endTime) returns(uint256)
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) CreateAuction(_nftAddress common.Address, _tokenId *big.Int, _startPrice *big.Int, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.CreateAuction(&_MyXAuctionV2.TransactOpts, _nftAddress, _tokenId, _startPrice, _startTime, _endTime)
}

// EndAuctionAndClaimNFT is a paid mutator transaction binding the contract method 0x2b48d9a3.
//
// Solidity: function endAuctionAndClaimNFT(uint256 _auctionId) returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) EndAuctionAndClaimNFT(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "endAuctionAndClaimNFT", _auctionId)
}

// EndAuctionAndClaimNFT is a paid mutator transaction binding the contract method 0x2b48d9a3.
//
// Solidity: function endAuctionAndClaimNFT(uint256 _auctionId) returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) EndAuctionAndClaimNFT(_auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.EndAuctionAndClaimNFT(&_MyXAuctionV2.TransactOpts, _auctionId)
}

// EndAuctionAndClaimNFT is a paid mutator transaction binding the contract method 0x2b48d9a3.
//
// Solidity: function endAuctionAndClaimNFT(uint256 _auctionId) returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) EndAuctionAndClaimNFT(_auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.EndAuctionAndClaimNFT(&_MyXAuctionV2.TransactOpts, _auctionId)
}

// ForceEndAuctionAndClaimNFT is a paid mutator transaction binding the contract method 0x85f63a19.
//
// Solidity: function forceEndAuctionAndClaimNFT(uint256 _auctionId) returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) ForceEndAuctionAndClaimNFT(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "forceEndAuctionAndClaimNFT", _auctionId)
}

// ForceEndAuctionAndClaimNFT is a paid mutator transaction binding the contract method 0x85f63a19.
//
// Solidity: function forceEndAuctionAndClaimNFT(uint256 _auctionId) returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) ForceEndAuctionAndClaimNFT(_auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.ForceEndAuctionAndClaimNFT(&_MyXAuctionV2.TransactOpts, _auctionId)
}

// ForceEndAuctionAndClaimNFT is a paid mutator transaction binding the contract method 0x85f63a19.
//
// Solidity: function forceEndAuctionAndClaimNFT(uint256 _auctionId) returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) ForceEndAuctionAndClaimNFT(_auctionId *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.ForceEndAuctionAndClaimNFT(&_MyXAuctionV2.TransactOpts, _auctionId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) Initialize() (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.Initialize(&_MyXAuctionV2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) Initialize() (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.Initialize(&_MyXAuctionV2.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) InitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "initializeV2")
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) InitializeV2() (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.InitializeV2(&_MyXAuctionV2.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) InitializeV2() (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.InitializeV2(&_MyXAuctionV2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) Pause() (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.Pause(&_MyXAuctionV2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) Pause() (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.Pause(&_MyXAuctionV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.RenounceOwnership(&_MyXAuctionV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.RenounceOwnership(&_MyXAuctionV2.TransactOpts)
}

// SetDynamicFeeEnabled is a paid mutator transaction binding the contract method 0x7948c325.
//
// Solidity: function setDynamicFeeEnabled(bool _enabled) returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) SetDynamicFeeEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "setDynamicFeeEnabled", _enabled)
}

// SetDynamicFeeEnabled is a paid mutator transaction binding the contract method 0x7948c325.
//
// Solidity: function setDynamicFeeEnabled(bool _enabled) returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) SetDynamicFeeEnabled(_enabled bool) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.SetDynamicFeeEnabled(&_MyXAuctionV2.TransactOpts, _enabled)
}

// SetDynamicFeeEnabled is a paid mutator transaction binding the contract method 0x7948c325.
//
// Solidity: function setDynamicFeeEnabled(bool _enabled) returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) SetDynamicFeeEnabled(_enabled bool) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.SetDynamicFeeEnabled(&_MyXAuctionV2.TransactOpts, _enabled)
}

// SetFeeTiers is a paid mutator transaction binding the contract method 0x1d57042c.
//
// Solidity: function setFeeTiers(uint256[] _thresholds, uint256[] _feeRates) returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) SetFeeTiers(opts *bind.TransactOpts, _thresholds []*big.Int, _feeRates []*big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "setFeeTiers", _thresholds, _feeRates)
}

// SetFeeTiers is a paid mutator transaction binding the contract method 0x1d57042c.
//
// Solidity: function setFeeTiers(uint256[] _thresholds, uint256[] _feeRates) returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) SetFeeTiers(_thresholds []*big.Int, _feeRates []*big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.SetFeeTiers(&_MyXAuctionV2.TransactOpts, _thresholds, _feeRates)
}

// SetFeeTiers is a paid mutator transaction binding the contract method 0x1d57042c.
//
// Solidity: function setFeeTiers(uint256[] _thresholds, uint256[] _feeRates) returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) SetFeeTiers(_thresholds []*big.Int, _feeRates []*big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.SetFeeTiers(&_MyXAuctionV2.TransactOpts, _thresholds, _feeRates)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _fee) returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) SetPlatformFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "setPlatformFee", _fee)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _fee) returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) SetPlatformFee(_fee *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.SetPlatformFee(&_MyXAuctionV2.TransactOpts, _fee)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _fee) returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) SetPlatformFee(_fee *big.Int) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.SetPlatformFee(&_MyXAuctionV2.TransactOpts, _fee)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address _token, address _priceFeed) returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) SetPriceFeed(opts *bind.TransactOpts, _token common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "setPriceFeed", _token, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address _token, address _priceFeed) returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) SetPriceFeed(_token common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.SetPriceFeed(&_MyXAuctionV2.TransactOpts, _token, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address _token, address _priceFeed) returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) SetPriceFeed(_token common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.SetPriceFeed(&_MyXAuctionV2.TransactOpts, _token, _priceFeed)
}

// SetPriceFeeds is a paid mutator transaction binding the contract method 0x92aeb53a.
//
// Solidity: function setPriceFeeds(address[] _tokens, address[] _priceFeeds) returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) SetPriceFeeds(opts *bind.TransactOpts, _tokens []common.Address, _priceFeeds []common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "setPriceFeeds", _tokens, _priceFeeds)
}

// SetPriceFeeds is a paid mutator transaction binding the contract method 0x92aeb53a.
//
// Solidity: function setPriceFeeds(address[] _tokens, address[] _priceFeeds) returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) SetPriceFeeds(_tokens []common.Address, _priceFeeds []common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.SetPriceFeeds(&_MyXAuctionV2.TransactOpts, _tokens, _priceFeeds)
}

// SetPriceFeeds is a paid mutator transaction binding the contract method 0x92aeb53a.
//
// Solidity: function setPriceFeeds(address[] _tokens, address[] _priceFeeds) returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) SetPriceFeeds(_tokens []common.Address, _priceFeeds []common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.SetPriceFeeds(&_MyXAuctionV2.TransactOpts, _tokens, _priceFeeds)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.TransferOwnership(&_MyXAuctionV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.TransferOwnership(&_MyXAuctionV2.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MyXAuctionV2 *MyXAuctionV2Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyXAuctionV2.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MyXAuctionV2 *MyXAuctionV2Session) Unpause() (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.Unpause(&_MyXAuctionV2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MyXAuctionV2 *MyXAuctionV2TransactorSession) Unpause() (*types.Transaction, error) {
	return _MyXAuctionV2.Contract.Unpause(&_MyXAuctionV2.TransactOpts)
}

// MyXAuctionV2AuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the MyXAuctionV2 contract.
type MyXAuctionV2AuctionCancelledIterator struct {
	Event *MyXAuctionV2AuctionCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2AuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2AuctionCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2AuctionCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2AuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2AuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2AuctionCancelled represents a AuctionCancelled event raised by the MyXAuctionV2 contract.
type MyXAuctionV2AuctionCancelled struct {
	AuctionId         *big.Int
	CancelledBy       common.Address
	Bidder            common.Address
	PaymentToken      common.Address
	RefundAmount      *big.Int
	RefundAmountValue *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x10ec610e9b6f8628e9a57f51bfda4adcb5db4c761791045d0f85214b681c7c48.
//
// Solidity: event AuctionCancelled(uint256 indexed auctionId, address indexed cancelledBy, address bidder, address indexed paymentToken, uint256 refundAmount, uint256 refundAmountValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterAuctionCancelled(opts *bind.FilterOpts, auctionId []*big.Int, cancelledBy []common.Address, paymentToken []common.Address) (*MyXAuctionV2AuctionCancelledIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var cancelledByRule []interface{}
	for _, cancelledByItem := range cancelledBy {
		cancelledByRule = append(cancelledByRule, cancelledByItem)
	}

	var paymentTokenRule []interface{}
	for _, paymentTokenItem := range paymentToken {
		paymentTokenRule = append(paymentTokenRule, paymentTokenItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "AuctionCancelled", auctionIdRule, cancelledByRule, paymentTokenRule)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2AuctionCancelledIterator{contract: _MyXAuctionV2.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x10ec610e9b6f8628e9a57f51bfda4adcb5db4c761791045d0f85214b681c7c48.
//
// Solidity: event AuctionCancelled(uint256 indexed auctionId, address indexed cancelledBy, address bidder, address indexed paymentToken, uint256 refundAmount, uint256 refundAmountValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2AuctionCancelled, auctionId []*big.Int, cancelledBy []common.Address, paymentToken []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var cancelledByRule []interface{}
	for _, cancelledByItem := range cancelledBy {
		cancelledByRule = append(cancelledByRule, cancelledByItem)
	}

	var paymentTokenRule []interface{}
	for _, paymentTokenItem := range paymentToken {
		paymentTokenRule = append(paymentTokenRule, paymentTokenItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "AuctionCancelled", auctionIdRule, cancelledByRule, paymentTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2AuctionCancelled)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionCancelled is a log parse operation binding the contract event 0x10ec610e9b6f8628e9a57f51bfda4adcb5db4c761791045d0f85214b681c7c48.
//
// Solidity: event AuctionCancelled(uint256 indexed auctionId, address indexed cancelledBy, address bidder, address indexed paymentToken, uint256 refundAmount, uint256 refundAmountValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseAuctionCancelled(log types.Log) (*MyXAuctionV2AuctionCancelled, error) {
	event := new(MyXAuctionV2AuctionCancelled)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2AuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the MyXAuctionV2 contract.
type MyXAuctionV2AuctionCreatedIterator struct {
	Event *MyXAuctionV2AuctionCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2AuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2AuctionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2AuctionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2AuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2AuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2AuctionCreated represents a AuctionCreated event raised by the MyXAuctionV2 contract.
type MyXAuctionV2AuctionCreated struct {
	AuctionId  *big.Int
	Creator    common.Address
	NftAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0x8777bed2a899ba8843de663a1f6ed1c48d071cc8bde08e2488b59c00c1993f76.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed creator, address indexed nftAddress, uint256 tokenId)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterAuctionCreated(opts *bind.FilterOpts, auctionId []*big.Int, creator []common.Address, nftAddress []common.Address) (*MyXAuctionV2AuctionCreatedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "AuctionCreated", auctionIdRule, creatorRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2AuctionCreatedIterator{contract: _MyXAuctionV2.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0x8777bed2a899ba8843de663a1f6ed1c48d071cc8bde08e2488b59c00c1993f76.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed creator, address indexed nftAddress, uint256 tokenId)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2AuctionCreated, auctionId []*big.Int, creator []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "AuctionCreated", auctionIdRule, creatorRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2AuctionCreated)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionCreated is a log parse operation binding the contract event 0x8777bed2a899ba8843de663a1f6ed1c48d071cc8bde08e2488b59c00c1993f76.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed creator, address indexed nftAddress, uint256 tokenId)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseAuctionCreated(log types.Log) (*MyXAuctionV2AuctionCreated, error) {
	event := new(MyXAuctionV2AuctionCreated)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2AuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the MyXAuctionV2 contract.
type MyXAuctionV2AuctionEndedIterator struct {
	Event *MyXAuctionV2AuctionEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2AuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2AuctionEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2AuctionEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2AuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2AuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2AuctionEnded represents a AuctionEnded event raised by the MyXAuctionV2 contract.
type MyXAuctionV2AuctionEnded struct {
	AuctionId    *big.Int
	Winner       common.Address
	FinalBid     *big.Int
	Seller       common.Address
	PaymentToken common.Address
	BidValue     *big.Int
	MinBidValue  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0x4805050c098aacbeda24d1234b31e570f6a2236ed167fd9516d8e87e92688399.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 finalBid, address seller, address paymentToken, uint256 bidValue, uint256 minBidValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterAuctionEnded(opts *bind.FilterOpts, auctionId []*big.Int, winner []common.Address) (*MyXAuctionV2AuctionEndedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "AuctionEnded", auctionIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2AuctionEndedIterator{contract: _MyXAuctionV2.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0x4805050c098aacbeda24d1234b31e570f6a2236ed167fd9516d8e87e92688399.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 finalBid, address seller, address paymentToken, uint256 bidValue, uint256 minBidValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2AuctionEnded, auctionId []*big.Int, winner []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "AuctionEnded", auctionIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2AuctionEnded)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionEnded is a log parse operation binding the contract event 0x4805050c098aacbeda24d1234b31e570f6a2236ed167fd9516d8e87e92688399.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 finalBid, address seller, address paymentToken, uint256 bidValue, uint256 minBidValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseAuctionEnded(log types.Log) (*MyXAuctionV2AuctionEnded, error) {
	event := new(MyXAuctionV2AuctionEnded)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2AuctionForceEndedIterator is returned from FilterAuctionForceEnded and is used to iterate over the raw logs and unpacked data for AuctionForceEnded events raised by the MyXAuctionV2 contract.
type MyXAuctionV2AuctionForceEndedIterator struct {
	Event *MyXAuctionV2AuctionForceEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2AuctionForceEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2AuctionForceEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2AuctionForceEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2AuctionForceEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2AuctionForceEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2AuctionForceEnded represents a AuctionForceEnded event raised by the MyXAuctionV2 contract.
type MyXAuctionV2AuctionForceEnded struct {
	AuctionId *big.Int
	EndedBy   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionForceEnded is a free log retrieval operation binding the contract event 0x77e5df52a34e08f2c8bf4ea9bea6de1ef93675f92c17e6a246847ab529d385ac.
//
// Solidity: event AuctionForceEnded(uint256 indexed auctionId, address indexed endedBy)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterAuctionForceEnded(opts *bind.FilterOpts, auctionId []*big.Int, endedBy []common.Address) (*MyXAuctionV2AuctionForceEndedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var endedByRule []interface{}
	for _, endedByItem := range endedBy {
		endedByRule = append(endedByRule, endedByItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "AuctionForceEnded", auctionIdRule, endedByRule)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2AuctionForceEndedIterator{contract: _MyXAuctionV2.contract, event: "AuctionForceEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionForceEnded is a free log subscription operation binding the contract event 0x77e5df52a34e08f2c8bf4ea9bea6de1ef93675f92c17e6a246847ab529d385ac.
//
// Solidity: event AuctionForceEnded(uint256 indexed auctionId, address indexed endedBy)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchAuctionForceEnded(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2AuctionForceEnded, auctionId []*big.Int, endedBy []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var endedByRule []interface{}
	for _, endedByItem := range endedBy {
		endedByRule = append(endedByRule, endedByItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "AuctionForceEnded", auctionIdRule, endedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2AuctionForceEnded)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "AuctionForceEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionForceEnded is a log parse operation binding the contract event 0x77e5df52a34e08f2c8bf4ea9bea6de1ef93675f92c17e6a246847ab529d385ac.
//
// Solidity: event AuctionForceEnded(uint256 indexed auctionId, address indexed endedBy)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseAuctionForceEnded(log types.Log) (*MyXAuctionV2AuctionForceEnded, error) {
	event := new(MyXAuctionV2AuctionForceEnded)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "AuctionForceEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2BidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the MyXAuctionV2 contract.
type MyXAuctionV2BidPlacedIterator struct {
	Event *MyXAuctionV2BidPlaced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2BidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2BidPlaced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2BidPlaced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2BidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2BidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2BidPlaced represents a BidPlaced event raised by the MyXAuctionV2 contract.
type MyXAuctionV2BidPlaced struct {
	AuctionId    *big.Int
	Bidder       common.Address
	Amount       *big.Int
	PaymentToken common.Address
	BidCount     *big.Int
	Timestamp    *big.Int
	BidValue     *big.Int
	MinBidder    common.Address
	MinBidValue  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x8c875a887f1bc54cb0fe2951468ca736412e5429164f585257ed15372e1d6b2e.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount, address indexed paymentToken, uint256 bidCount, uint256 timestamp, uint256 bidValue, address minBidder, uint256 minBidValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterBidPlaced(opts *bind.FilterOpts, auctionId []*big.Int, bidder []common.Address, paymentToken []common.Address) (*MyXAuctionV2BidPlacedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var paymentTokenRule []interface{}
	for _, paymentTokenItem := range paymentToken {
		paymentTokenRule = append(paymentTokenRule, paymentTokenItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "BidPlaced", auctionIdRule, bidderRule, paymentTokenRule)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2BidPlacedIterator{contract: _MyXAuctionV2.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x8c875a887f1bc54cb0fe2951468ca736412e5429164f585257ed15372e1d6b2e.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount, address indexed paymentToken, uint256 bidCount, uint256 timestamp, uint256 bidValue, address minBidder, uint256 minBidValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2BidPlaced, auctionId []*big.Int, bidder []common.Address, paymentToken []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var paymentTokenRule []interface{}
	for _, paymentTokenItem := range paymentToken {
		paymentTokenRule = append(paymentTokenRule, paymentTokenItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "BidPlaced", auctionIdRule, bidderRule, paymentTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2BidPlaced)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "BidPlaced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBidPlaced is a log parse operation binding the contract event 0x8c875a887f1bc54cb0fe2951468ca736412e5429164f585257ed15372e1d6b2e.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount, address indexed paymentToken, uint256 bidCount, uint256 timestamp, uint256 bidValue, address minBidder, uint256 minBidValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseBidPlaced(log types.Log) (*MyXAuctionV2BidPlaced, error) {
	event := new(MyXAuctionV2BidPlaced)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2BidValueTooLowIterator is returned from FilterBidValueTooLow and is used to iterate over the raw logs and unpacked data for BidValueTooLow events raised by the MyXAuctionV2 contract.
type MyXAuctionV2BidValueTooLowIterator struct {
	Event *MyXAuctionV2BidValueTooLow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2BidValueTooLowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2BidValueTooLow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2BidValueTooLow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2BidValueTooLowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2BidValueTooLowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2BidValueTooLow represents a BidValueTooLow event raised by the MyXAuctionV2 contract.
type MyXAuctionV2BidValueTooLow struct {
	AuctionId    *big.Int
	Bidder       common.Address
	Amount       *big.Int
	PaymentToken common.Address
	BidCount     *big.Int
	Timestamp    *big.Int
	BidValue     *big.Int
	MinBidder    common.Address
	MinBidValue  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBidValueTooLow is a free log retrieval operation binding the contract event 0xe814151408fcc11f8069f77262418f2fdd23e91da33bd56d4c27c2d959521f59.
//
// Solidity: event BidValueTooLow(uint256 indexed auctionId, address indexed bidder, uint256 amount, address indexed paymentToken, uint256 bidCount, uint256 timestamp, uint256 bidValue, address minBidder, uint256 minBidValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterBidValueTooLow(opts *bind.FilterOpts, auctionId []*big.Int, bidder []common.Address, paymentToken []common.Address) (*MyXAuctionV2BidValueTooLowIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var paymentTokenRule []interface{}
	for _, paymentTokenItem := range paymentToken {
		paymentTokenRule = append(paymentTokenRule, paymentTokenItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "BidValueTooLow", auctionIdRule, bidderRule, paymentTokenRule)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2BidValueTooLowIterator{contract: _MyXAuctionV2.contract, event: "BidValueTooLow", logs: logs, sub: sub}, nil
}

// WatchBidValueTooLow is a free log subscription operation binding the contract event 0xe814151408fcc11f8069f77262418f2fdd23e91da33bd56d4c27c2d959521f59.
//
// Solidity: event BidValueTooLow(uint256 indexed auctionId, address indexed bidder, uint256 amount, address indexed paymentToken, uint256 bidCount, uint256 timestamp, uint256 bidValue, address minBidder, uint256 minBidValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchBidValueTooLow(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2BidValueTooLow, auctionId []*big.Int, bidder []common.Address, paymentToken []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var paymentTokenRule []interface{}
	for _, paymentTokenItem := range paymentToken {
		paymentTokenRule = append(paymentTokenRule, paymentTokenItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "BidValueTooLow", auctionIdRule, bidderRule, paymentTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2BidValueTooLow)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "BidValueTooLow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBidValueTooLow is a log parse operation binding the contract event 0xe814151408fcc11f8069f77262418f2fdd23e91da33bd56d4c27c2d959521f59.
//
// Solidity: event BidValueTooLow(uint256 indexed auctionId, address indexed bidder, uint256 amount, address indexed paymentToken, uint256 bidCount, uint256 timestamp, uint256 bidValue, address minBidder, uint256 minBidValue)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseBidValueTooLow(log types.Log) (*MyXAuctionV2BidValueTooLow, error) {
	event := new(MyXAuctionV2BidValueTooLow)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "BidValueTooLow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2DynamicFeeEnabledIterator is returned from FilterDynamicFeeEnabled and is used to iterate over the raw logs and unpacked data for DynamicFeeEnabled events raised by the MyXAuctionV2 contract.
type MyXAuctionV2DynamicFeeEnabledIterator struct {
	Event *MyXAuctionV2DynamicFeeEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2DynamicFeeEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2DynamicFeeEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2DynamicFeeEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2DynamicFeeEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2DynamicFeeEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2DynamicFeeEnabled represents a DynamicFeeEnabled event raised by the MyXAuctionV2 contract.
type MyXAuctionV2DynamicFeeEnabled struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDynamicFeeEnabled is a free log retrieval operation binding the contract event 0x5f0dd08ee0bed9364680dd75e895a44ad73286aa34423ee7531fe7841f4e76d1.
//
// Solidity: event DynamicFeeEnabled(bool enabled)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterDynamicFeeEnabled(opts *bind.FilterOpts) (*MyXAuctionV2DynamicFeeEnabledIterator, error) {

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "DynamicFeeEnabled")
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2DynamicFeeEnabledIterator{contract: _MyXAuctionV2.contract, event: "DynamicFeeEnabled", logs: logs, sub: sub}, nil
}

// WatchDynamicFeeEnabled is a free log subscription operation binding the contract event 0x5f0dd08ee0bed9364680dd75e895a44ad73286aa34423ee7531fe7841f4e76d1.
//
// Solidity: event DynamicFeeEnabled(bool enabled)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchDynamicFeeEnabled(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2DynamicFeeEnabled) (event.Subscription, error) {

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "DynamicFeeEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2DynamicFeeEnabled)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "DynamicFeeEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDynamicFeeEnabled is a log parse operation binding the contract event 0x5f0dd08ee0bed9364680dd75e895a44ad73286aa34423ee7531fe7841f4e76d1.
//
// Solidity: event DynamicFeeEnabled(bool enabled)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseDynamicFeeEnabled(log types.Log) (*MyXAuctionV2DynamicFeeEnabled, error) {
	event := new(MyXAuctionV2DynamicFeeEnabled)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "DynamicFeeEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2FeeTierUpdatedIterator is returned from FilterFeeTierUpdated and is used to iterate over the raw logs and unpacked data for FeeTierUpdated events raised by the MyXAuctionV2 contract.
type MyXAuctionV2FeeTierUpdatedIterator struct {
	Event *MyXAuctionV2FeeTierUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2FeeTierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2FeeTierUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2FeeTierUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2FeeTierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2FeeTierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2FeeTierUpdated represents a FeeTierUpdated event raised by the MyXAuctionV2 contract.
type MyXAuctionV2FeeTierUpdated struct {
	TierIndex *big.Int
	Threshold *big.Int
	FeeRate   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeTierUpdated is a free log retrieval operation binding the contract event 0x050378e8543487fde7e7bd39a32579544cecfec6b7bd3f81082141eb9c255c48.
//
// Solidity: event FeeTierUpdated(uint256 indexed tierIndex, uint256 threshold, uint256 feeRate)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterFeeTierUpdated(opts *bind.FilterOpts, tierIndex []*big.Int) (*MyXAuctionV2FeeTierUpdatedIterator, error) {

	var tierIndexRule []interface{}
	for _, tierIndexItem := range tierIndex {
		tierIndexRule = append(tierIndexRule, tierIndexItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "FeeTierUpdated", tierIndexRule)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2FeeTierUpdatedIterator{contract: _MyXAuctionV2.contract, event: "FeeTierUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeTierUpdated is a free log subscription operation binding the contract event 0x050378e8543487fde7e7bd39a32579544cecfec6b7bd3f81082141eb9c255c48.
//
// Solidity: event FeeTierUpdated(uint256 indexed tierIndex, uint256 threshold, uint256 feeRate)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchFeeTierUpdated(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2FeeTierUpdated, tierIndex []*big.Int) (event.Subscription, error) {

	var tierIndexRule []interface{}
	for _, tierIndexItem := range tierIndex {
		tierIndexRule = append(tierIndexRule, tierIndexItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "FeeTierUpdated", tierIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2FeeTierUpdated)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "FeeTierUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeTierUpdated is a log parse operation binding the contract event 0x050378e8543487fde7e7bd39a32579544cecfec6b7bd3f81082141eb9c255c48.
//
// Solidity: event FeeTierUpdated(uint256 indexed tierIndex, uint256 threshold, uint256 feeRate)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseFeeTierUpdated(log types.Log) (*MyXAuctionV2FeeTierUpdated, error) {
	event := new(MyXAuctionV2FeeTierUpdated)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "FeeTierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MyXAuctionV2 contract.
type MyXAuctionV2InitializedIterator struct {
	Event *MyXAuctionV2Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2Initialized represents a Initialized event raised by the MyXAuctionV2 contract.
type MyXAuctionV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*MyXAuctionV2InitializedIterator, error) {

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2InitializedIterator{contract: _MyXAuctionV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2Initialized) (event.Subscription, error) {

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2Initialized)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseInitialized(log types.Log) (*MyXAuctionV2Initialized, error) {
	event := new(MyXAuctionV2Initialized)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2NFTApprovalCancelledIterator is returned from FilterNFTApprovalCancelled and is used to iterate over the raw logs and unpacked data for NFTApprovalCancelled events raised by the MyXAuctionV2 contract.
type MyXAuctionV2NFTApprovalCancelledIterator struct {
	Event *MyXAuctionV2NFTApprovalCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2NFTApprovalCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2NFTApprovalCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2NFTApprovalCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2NFTApprovalCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2NFTApprovalCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2NFTApprovalCancelled represents a NFTApprovalCancelled event raised by the MyXAuctionV2 contract.
type MyXAuctionV2NFTApprovalCancelled struct {
	Owner      common.Address
	NftAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTApprovalCancelled is a free log retrieval operation binding the contract event 0x508327cabf1ba4e8ea2d0b51016f5a153adac8f237e5b675759b86e9c0fc00de.
//
// Solidity: event NFTApprovalCancelled(address indexed owner, address indexed nftAddress, uint256 tokenId)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterNFTApprovalCancelled(opts *bind.FilterOpts, owner []common.Address, nftAddress []common.Address) (*MyXAuctionV2NFTApprovalCancelledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "NFTApprovalCancelled", ownerRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2NFTApprovalCancelledIterator{contract: _MyXAuctionV2.contract, event: "NFTApprovalCancelled", logs: logs, sub: sub}, nil
}

// WatchNFTApprovalCancelled is a free log subscription operation binding the contract event 0x508327cabf1ba4e8ea2d0b51016f5a153adac8f237e5b675759b86e9c0fc00de.
//
// Solidity: event NFTApprovalCancelled(address indexed owner, address indexed nftAddress, uint256 tokenId)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchNFTApprovalCancelled(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2NFTApprovalCancelled, owner []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "NFTApprovalCancelled", ownerRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2NFTApprovalCancelled)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "NFTApprovalCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTApprovalCancelled is a log parse operation binding the contract event 0x508327cabf1ba4e8ea2d0b51016f5a153adac8f237e5b675759b86e9c0fc00de.
//
// Solidity: event NFTApprovalCancelled(address indexed owner, address indexed nftAddress, uint256 tokenId)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseNFTApprovalCancelled(log types.Log) (*MyXAuctionV2NFTApprovalCancelled, error) {
	event := new(MyXAuctionV2NFTApprovalCancelled)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "NFTApprovalCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2NFTApprovedIterator is returned from FilterNFTApproved and is used to iterate over the raw logs and unpacked data for NFTApproved events raised by the MyXAuctionV2 contract.
type MyXAuctionV2NFTApprovedIterator struct {
	Event *MyXAuctionV2NFTApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2NFTApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2NFTApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2NFTApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2NFTApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2NFTApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2NFTApproved represents a NFTApproved event raised by the MyXAuctionV2 contract.
type MyXAuctionV2NFTApproved struct {
	Owner      common.Address
	NftAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTApproved is a free log retrieval operation binding the contract event 0x515fe62e6a28ce5c73f42b5955b1c53631fa49fe144c9dc8eadd9b548d6713cb.
//
// Solidity: event NFTApproved(address indexed owner, address indexed nftAddress, uint256 tokenId)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterNFTApproved(opts *bind.FilterOpts, owner []common.Address, nftAddress []common.Address) (*MyXAuctionV2NFTApprovedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "NFTApproved", ownerRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2NFTApprovedIterator{contract: _MyXAuctionV2.contract, event: "NFTApproved", logs: logs, sub: sub}, nil
}

// WatchNFTApproved is a free log subscription operation binding the contract event 0x515fe62e6a28ce5c73f42b5955b1c53631fa49fe144c9dc8eadd9b548d6713cb.
//
// Solidity: event NFTApproved(address indexed owner, address indexed nftAddress, uint256 tokenId)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchNFTApproved(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2NFTApproved, owner []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "NFTApproved", ownerRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2NFTApproved)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "NFTApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTApproved is a log parse operation binding the contract event 0x515fe62e6a28ce5c73f42b5955b1c53631fa49fe144c9dc8eadd9b548d6713cb.
//
// Solidity: event NFTApproved(address indexed owner, address indexed nftAddress, uint256 tokenId)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseNFTApproved(log types.Log) (*MyXAuctionV2NFTApproved, error) {
	event := new(MyXAuctionV2NFTApproved)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "NFTApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MyXAuctionV2 contract.
type MyXAuctionV2OwnershipTransferredIterator struct {
	Event *MyXAuctionV2OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2OwnershipTransferred represents a OwnershipTransferred event raised by the MyXAuctionV2 contract.
type MyXAuctionV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MyXAuctionV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2OwnershipTransferredIterator{contract: _MyXAuctionV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2OwnershipTransferred)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseOwnershipTransferred(log types.Log) (*MyXAuctionV2OwnershipTransferred, error) {
	event := new(MyXAuctionV2OwnershipTransferred)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MyXAuctionV2 contract.
type MyXAuctionV2PausedIterator struct {
	Event *MyXAuctionV2Paused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2Paused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2Paused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2Paused represents a Paused event raised by the MyXAuctionV2 contract.
type MyXAuctionV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterPaused(opts *bind.FilterOpts) (*MyXAuctionV2PausedIterator, error) {

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2PausedIterator{contract: _MyXAuctionV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2Paused) (event.Subscription, error) {

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2Paused)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParsePaused(log types.Log) (*MyXAuctionV2Paused, error) {
	event := new(MyXAuctionV2Paused)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2PlatformFeeUpdatedIterator is returned from FilterPlatformFeeUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeUpdated events raised by the MyXAuctionV2 contract.
type MyXAuctionV2PlatformFeeUpdatedIterator struct {
	Event *MyXAuctionV2PlatformFeeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2PlatformFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2PlatformFeeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2PlatformFeeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2PlatformFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2PlatformFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2PlatformFeeUpdated represents a PlatformFeeUpdated event raised by the MyXAuctionV2 contract.
type MyXAuctionV2PlatformFeeUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeUpdated is a free log retrieval operation binding the contract event 0xd347e206f25a89b917fc9482f1a2d294d749baa4dc9bde7fb495ee11fe491643.
//
// Solidity: event PlatformFeeUpdated(uint256 oldFee, uint256 newFee)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterPlatformFeeUpdated(opts *bind.FilterOpts) (*MyXAuctionV2PlatformFeeUpdatedIterator, error) {

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2PlatformFeeUpdatedIterator{contract: _MyXAuctionV2.contract, event: "PlatformFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeUpdated is a free log subscription operation binding the contract event 0xd347e206f25a89b917fc9482f1a2d294d749baa4dc9bde7fb495ee11fe491643.
//
// Solidity: event PlatformFeeUpdated(uint256 oldFee, uint256 newFee)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchPlatformFeeUpdated(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2PlatformFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2PlatformFeeUpdated)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePlatformFeeUpdated is a log parse operation binding the contract event 0xd347e206f25a89b917fc9482f1a2d294d749baa4dc9bde7fb495ee11fe491643.
//
// Solidity: event PlatformFeeUpdated(uint256 oldFee, uint256 newFee)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParsePlatformFeeUpdated(log types.Log) (*MyXAuctionV2PlatformFeeUpdated, error) {
	event := new(MyXAuctionV2PlatformFeeUpdated)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2TotalValueLockedUpdatedIterator is returned from FilterTotalValueLockedUpdated and is used to iterate over the raw logs and unpacked data for TotalValueLockedUpdated events raised by the MyXAuctionV2 contract.
type MyXAuctionV2TotalValueLockedUpdatedIterator struct {
	Event *MyXAuctionV2TotalValueLockedUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2TotalValueLockedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2TotalValueLockedUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2TotalValueLockedUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2TotalValueLockedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2TotalValueLockedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2TotalValueLockedUpdated represents a TotalValueLockedUpdated event raised by the MyXAuctionV2 contract.
type MyXAuctionV2TotalValueLockedUpdated struct {
	NewTVL          *big.Int
	TotalBidsPlaced *big.Int
	Change          *big.Int
	IsIncrease      bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTotalValueLockedUpdated is a free log retrieval operation binding the contract event 0xfe9b183fcbd8416e8d2ddcc6180ff8e7a33c363b7191deccac34f694c4faab29.
//
// Solidity: event TotalValueLockedUpdated(uint256 newTVL, uint256 totalBidsPlaced, uint256 change, bool isIncrease)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterTotalValueLockedUpdated(opts *bind.FilterOpts) (*MyXAuctionV2TotalValueLockedUpdatedIterator, error) {

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "TotalValueLockedUpdated")
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2TotalValueLockedUpdatedIterator{contract: _MyXAuctionV2.contract, event: "TotalValueLockedUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalValueLockedUpdated is a free log subscription operation binding the contract event 0xfe9b183fcbd8416e8d2ddcc6180ff8e7a33c363b7191deccac34f694c4faab29.
//
// Solidity: event TotalValueLockedUpdated(uint256 newTVL, uint256 totalBidsPlaced, uint256 change, bool isIncrease)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchTotalValueLockedUpdated(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2TotalValueLockedUpdated) (event.Subscription, error) {

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "TotalValueLockedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2TotalValueLockedUpdated)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "TotalValueLockedUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTotalValueLockedUpdated is a log parse operation binding the contract event 0xfe9b183fcbd8416e8d2ddcc6180ff8e7a33c363b7191deccac34f694c4faab29.
//
// Solidity: event TotalValueLockedUpdated(uint256 newTVL, uint256 totalBidsPlaced, uint256 change, bool isIncrease)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseTotalValueLockedUpdated(log types.Log) (*MyXAuctionV2TotalValueLockedUpdated, error) {
	event := new(MyXAuctionV2TotalValueLockedUpdated)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "TotalValueLockedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyXAuctionV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MyXAuctionV2 contract.
type MyXAuctionV2UnpausedIterator struct {
	Event *MyXAuctionV2Unpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyXAuctionV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXAuctionV2Unpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyXAuctionV2Unpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyXAuctionV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXAuctionV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXAuctionV2Unpaused represents a Unpaused event raised by the MyXAuctionV2 contract.
type MyXAuctionV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*MyXAuctionV2UnpausedIterator, error) {

	logs, sub, err := _MyXAuctionV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MyXAuctionV2UnpausedIterator{contract: _MyXAuctionV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MyXAuctionV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _MyXAuctionV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXAuctionV2Unpaused)
				if err := _MyXAuctionV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MyXAuctionV2 *MyXAuctionV2Filterer) ParseUnpaused(log types.Log) (*MyXAuctionV2Unpaused, error) {
	event := new(MyXAuctionV2Unpaused)
	if err := _MyXAuctionV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
