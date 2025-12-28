package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	appJWT "my-auction-market-api/internal/jwt"
	"my-auction-market-api/internal/models"
	"my-auction-market-api/internal/page"
	"my-auction-market-api/internal/response"
	"my-auction-market-api/internal/services"
)

type NFTHandler struct {
	service *services.NFTService
}

func NewNFTHandler(nftService *services.NFTService) *NFTHandler {
	return &NFTHandler{
		service: nftService,
	}
}

// GetOwnedNFTs godoc
// @Summary      Get owned NFTs
// @Description  Get NFTs owned by the current user from a specific contract address
// @Tags         nft
// @Accept       json
// @Produce      json
// @Param        contractAddress  query     string  false  "NFT contract address (optional, uses default if not provided)"
// @Param        page             query     int     false  "Page number" default(1)
// @Param        pageSize         query     int     false  "Page size" default(10)
// @Success      200              {object}  response.Response
// @Failure      400              {object}  response.Response
// @Failure      401              {object}  response.Response
// @Failure      500              {object}  response.Response
// @Security     BearerAuth
// @Router       /nfts/owned [get]
func (h *NFTHandler) GetOwnedNFTs(c *gin.Context) {
	user, err := appJWT.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	contractAddress := c.Query("contractAddress")

	var query page.PageQuery
	if err := query.Bind(c); err != nil {
		return
	}

	// TODO: 实现业务逻辑
	// nfts, total, err := h.service.GetOwnedNFTs(user.ID, contractAddress, query)
	// if err != nil {
	// 	response.Error(c, err)
	// 	return
	// }

	// pageData := page.NewPageData(query.Page, query.PageSize, total, nfts)
	// response.Success(c, pageData)

	response.Success(c, gin.H{
		"message":         "GetOwnedNFTs - not implemented yet",
		"userId":          user.ID,
		"contractAddress": contractAddress,
	})
}

// SyncNFTs godoc
// @Summary      Sync user NFTs
// @Description  Synchronize all NFTs from blockchain to database for the current user
// @Tags         nft
// @Accept       json
// @Produce      json
// @Success      200              {object}  response.Response{data=models.NFTSyncResult}
// @Failure      400              {object}  response.Response
// @Failure      401              {object}  response.Response
// @Failure      500              {object}  response.Response
// @Security     BearerAuth
// @Router       /nfts/sync [post]
func (h *NFTHandler) SyncNFTs(c *gin.Context) {
	user, err := appJWT.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}
	result, err := h.service.SyncNFTs(user.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, result)
}

// GetSyncStatus godoc
// @Summary      Get NFT sync status
// @Description  Get the synchronization status of user's NFTs
// @Tags         nft
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Security     BearerAuth
// @Router       /nfts/sync/status [get]
func (h *NFTHandler) GetSyncStatus(c *gin.Context) {
	user, err := appJWT.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	// TODO: 实现业务逻辑
	// status, err := h.service.GetSyncStatus(user.ID)
	// if err != nil {
	// 	response.Error(c, err)
	// 	return
	// }

	// response.Success(c, status)

	response.Success(c, gin.H{
		"message": "GetSyncStatus - not implemented yet",
		"userId":  user.ID,
	})
}

// GetNFTByID godoc
// @Summary      Get NFT by ID
// @Description  Get a single NFT by ID from database
// @Tags         nft
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "NFT ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Security     BearerAuth
// @Router       /nfts/{id} [get]
func (h *NFTHandler) GetNFTByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid nft id")
		return
	}

	// TODO: 实现业务逻辑
	// nft, err := h.service.GetByID(id)
	// if err != nil {
	// 	response.NotFound(c, err.Error())
	// 	return
	// }

	// response.Success(c, nft)

	response.Success(c, gin.H{
		"message": "GetNFTByID - not implemented yet",
		"id":      id,
	})
}

// VerifyOwnership godoc
// @Summary      Verify NFT ownership
// @Description  Verify if the current user owns a specific NFT
// @Tags         nft
// @Accept       json
// @Produce      json
// @Param        verify  body      models.NFTOwnershipVerifyPayload  true  "NFT ownership verify payload"
// @Success      200     {object}  response.Response
// @Failure      400     {object}  response.Response
// @Failure      401     {object}  response.Response
// @Failure      500     {object}  response.Response
// @Security     BearerAuth
// @Router       /nfts/verify [post]
func (h *NFTHandler) VerifyOwnership(c *gin.Context) {
	user, err := appJWT.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	var payload models.NFTOwnershipVerifyPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.Error(err)
		return
	}

	// TODO: 实现业务逻辑
	// isValid, err := h.service.VerifyOwnership(user.ID, payload)
	// if err != nil {
	// 	response.Error(c, err)
	// 	return
	// }

	// response.Success(c, gin.H{
	// 	"isValid": isValid,
	// })

	response.Success(c, gin.H{
		"message": "VerifyOwnership - not implemented yet",
		"userId":  user.ID,
		"payload": payload,
	})
}

// GetMyNFTs godoc
// @Summary      Get my NFTs
// @Description  Get all NFTs owned by the current user from database
// @Tags         nft
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number" default(1)
// @Param        pageSize  query     int     false  "Page size" default(10)
// @Param        contractAddress  query     string  false  "Filter by contract address"
// @Success      200       {object}  response.Response
// @Failure      400       {object}  response.Response
// @Failure      401       {object}  response.Response
// @Failure      500       {object}  response.Response
// @Security     BearerAuth
// @Router       /nfts/my [get]
func (h *NFTHandler) GetMyNFTs(c *gin.Context) {
	user, err := appJWT.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	contractAddress := c.Query("contractAddress")

	var query page.PageQuery
	if err := query.Bind(c); err != nil {
		return
	}

	// TODO: 实现业务逻辑
	// nfts, total, err := h.service.GetMyNFTs(user.ID, contractAddress, query)
	// if err != nil {
	// 	response.Error(c, err)
	// 	return
	// }

	// pageData := page.NewPageData(query.Page, query.PageSize, total, nfts)
	// response.Success(c, pageData)

	response.Success(c, gin.H{
		"message":         "GetMyNFTs - not implemented yet",
		"userId":          user.ID,
		"contractAddress": contractAddress,
	})
}

// GetMyNFTsList godoc
// @Summary      Get my NFTs list (no pagination)
// @Description  Get all NFTs owned by the current user from database without pagination (all contracts)
// @Tags         nft
// @Accept       json
// @Produce      json
// @Param        status    query     string  false  "Filter by status (holding, selling, sold, transfered, all)" default(all)
// @Success      200       {object}  response.Response
// @Failure      400       {object}  response.Response
// @Failure      401       {object}  response.Response
// @Failure      500       {object}  response.Response
// @Security     BearerAuth
// @Router       /nfts/my/list [get]
func (h *NFTHandler) GetMyNFTsList(c *gin.Context) {
	user, err := appJWT.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	// 获取状态筛选参数
	statusFilter := c.DefaultQuery("status", "all")

	ownerships, err := h.service.GetMyNFTsList(user.ID, statusFilter)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, ownerships)
}

// GetMyNFTOwnershipByNFTID godoc
// @Summary      Get my NFT ownership by NFT ID
// @Description  Get a single NFT ownership record by nftId for the current user
// @Tags         nft
// @Accept       json
// @Produce      json
// @Param        nftId   path      string  true  "NFT ID"
// @Success      200     {object}  response.Response
// @Failure      400     {object}  response.Response
// @Failure      401     {object}  response.Response
// @Failure      404     {object}  response.Response
// @Failure      500     {object}  response.Response
// @Security     BearerAuth
// @Router       /nfts/my/ownership/{nftId} [get]
func (h *NFTHandler) GetMyNFTOwnershipByNFTID(c *gin.Context) {
	user, err := appJWT.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	nftID := c.Param("nftId")
	if nftID == "" {
		response.BadRequest(c, "nftId is required")
		return
	}

	ownership, err := h.service.GetMyNFTOwnershipByNFTID(user.ID, nftID)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, ownership)
}
