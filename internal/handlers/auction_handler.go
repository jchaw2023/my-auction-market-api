package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"my-auction-market-api/internal/jwt"
	"my-auction-market-api/internal/models"
	"my-auction-market-api/internal/page"
	"my-auction-market-api/internal/response"
	"my-auction-market-api/internal/services"
)

type AuctionHandler struct {
	service *services.AuctionService
}

func NewAuctionHandler(auctionService *services.AuctionService) *AuctionHandler {
	return &AuctionHandler{
		service: auctionService,
	}
}

// List godoc
// @Summary      List auctions
// @Description  Get a paginated list of auctions
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number" default(1)
// @Param        pageSize  query     int     false  "Page size" default(10)
// @Success      200       {object}  response.Response
// @Failure      400       {object}  response.Response
// @Failure      500       {object}  response.Response
// @Router       /auctions [get]
func (h *AuctionHandler) List(c *gin.Context) {
	var query page.PageQuery
	if err := query.Bind(c); err != nil {
		return
	}

	auctions, total, err := h.service.List(query)
	if err != nil {
		response.Error(c, err)
		return
	}

	pageData := page.NewPageData(query.Page, query.PageSize, total, auctions)
	response.Success(c, pageData)
}

// ListPublic godoc
// @Summary      List public auctions (for home page)
// @Description  Get a paginated list of public auctions, sorted by status (active first, then ended) and time
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number" default(1)
// @Param        pageSize  query     int     false  "Page size" default(10)
// @Param        status    query     string  false  "Filter by status (active, ended, all)" default(all)
// @Success      200       {object}  response.Response
// @Failure      400       {object}  response.Response
// @Failure      500       {object}  response.Response
// @Router       /auctions/public [get]
func (h *AuctionHandler) ListPublic(c *gin.Context) {
	var query page.PageQuery
	if err := query.Bind(c); err != nil {
		return
	}

	// 获取状态筛选参数
	statusFilter := c.DefaultQuery("status", "all")

	auctions, total, err := h.service.ListPublic(query, statusFilter)
	if err != nil {
		response.Error(c, err)
		return
	}

	pageData := page.NewPageData(query.Page, query.PageSize, total, auctions)
	response.Success(c, pageData)
}

// GetByID godoc
// @Summary      Get auction by ID
// @Description  Get a single auction record by AuctionID (string, basic auction information)
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Auction ID (string)"
// @Success      200  {object}  response.Response{data=models.Auction}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /auctions/{id} [get]
func (h *AuctionHandler) GetByID(c *gin.Context) {
	auctionID := c.Param("id")
	if auctionID == "" {
		response.BadRequest(c, "auction id is required")
		return
	}

	auction, err := h.service.GetByID(auctionID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, auction)
}

// GetDetailByID godoc
// @Summary      Get auction detail by ID
// @Description  Get detailed auction information by ID, including seller wallet address (no full user info or bids)
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Auction ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /auctions/{id}/detail [get]
func (h *AuctionHandler) GetDetailByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid auction id")
		return
	}

	auctionDetail, err := h.service.GetDetailByID(id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, auctionDetail)
}

// Create godoc
// @Summary      Create a new auction
// @Description  Create a new auction with the provided information
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        auction  body      models.AuctionPayload  true  "Auction payload"
// @Success      201      {object}  response.Response
// @Failure      400      {object}  response.Response
// @Failure      401      {object}  response.Response
// @Failure      500      {object}  response.Response
// @Security     BearerAuth
// @Router       /auctions [post]
func (h *AuctionHandler) Create(c *gin.Context) {
	var payload models.AuctionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.Error(err)
		return
	}

	user, err := jwt.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	auction, err := h.service.Create(user.ID, payload)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, auction)
}

// GetUserAuctions godoc
// @Summary      Get user auctions
// @Description  Get auctions created by the current user
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        page      query     int      false  "Page number" default(1)
// @Param        pageSize  query     int      false  "Page size" default(10)
// @Param        status    query     []string false  "Filter by status (pending, active, ended, cancelled)" collectionFormat(multi)
// @Success      200       {object}  response.Response
// @Failure      400       {object}  response.Response
// @Failure      401       {object}  response.Response
// @Failure      500       {object}  response.Response
// @Security     BearerAuth
// @Router       /auctions/my [get]
func (h *AuctionHandler) GetUserAuctions(c *gin.Context) {
	user, err := jwt.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	var query page.PageQuery
	if err := query.Bind(c); err != nil {
		return
	}

	// 获取 status 查询参数（支持多个值）
	statuses := c.QueryArray("status")

	auctions, total, err := h.service.GetUserAuctions(user.ID, query, statuses)
	if err != nil {
		response.Error(c, err)
		return
	}

	pageData := page.NewPageData(query.Page, query.PageSize, total, auctions)
	response.Success(c, pageData)
}

// Update godoc
// @Summary      Update auction
// @Description  Update an existing auction (only pending auctions can be updated)
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        id      path      string                     true  "Auction ID (string)"
// @Param        auction body      models.UpdateAuctionPayload true  "Update auction payload"
// @Success      200     {object}  response.Response
// @Failure      400     {object}  response.Response
// @Failure      401     {object}  response.Response
// @Failure      403     {object}  response.Response
// @Failure      404     {object}  response.Response
// @Failure      500     {object}  response.Response
// @Security     BearerAuth
// @Router       /auctions/{id} [put]
func (h *AuctionHandler) Update(c *gin.Context) {
	auctionID := c.Param("id")
	if auctionID == "" {
		response.BadRequest(c, "auction id is required")
		return
	}

	var payload models.UpdateAuctionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, err := jwt.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	auction, err := h.service.Update(user.ID, auctionID, payload)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, auction)
}

// ConvertToUSD godoc
// @Summary      Convert amount to USD
// @Description  Convert token amount to USD value using Chainlink price oracle
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        request  body      models.ConvertToUSDPayload  true  "Convert to USD payload"
// @Success      200      {object}  response.Response
// @Failure      400      {object}  response.Response
// @Failure      500      {object}  response.Response
// @Router       /auctions/convert-to-usd [post]
func (h *AuctionHandler) ConvertToUSD(c *gin.Context) {
	var payload models.ConvertToUSDPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.service.ConvertTokenAmountToUSD(payload.Token, payload.Amount)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, result)
}

// GetTokenPrice godoc
// @Summary      Get token price
// @Description  Get token price in USD using Chainlink price oracle
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        token   path      string  true  "Token address (0x0 for ETH)"
// @Success      200     {object}  response.Response{data=models.TokenPriceResponse}
// @Failure      400     {object}  response.Response
// @Failure      500     {object}  response.Response
// @Router       /auctions/token-price/{token} [get]
func (h *AuctionHandler) GetTokenPrice(c *gin.Context) {
	tokenAddress := c.Param("token")
	if tokenAddress == "" {
		response.BadRequest(c, "token address is required")
		return
	}

	result, err := h.service.GetTokenPrice(tokenAddress)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, result)
}

// ListNFTs godoc
// @Summary      List NFTs in auctions
// @Description  Get a paginated list of unique NFTs from all auctions
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number" default(1)
// @Param        pageSize  query     int     false  "Page size" default(10)
// @Success      200       {object}  response.Response
// @Failure      400       {object}  response.Response
// @Failure      500       {object}  response.Response
// @Router       /auctions/nfts [get]
func (h *AuctionHandler) ListNFTs(c *gin.Context) {
	var query page.PageQuery
	if err := query.Bind(c); err != nil {
		return
	}

	nfts, total, err := h.service.ListNFTs(query)
	if err != nil {
		response.Error(c, err)
		return
	}

	pageData := page.NewPageData(query.Page, query.PageSize, total, nfts)
	response.Success(c, pageData)
}

// GetSupportedTokens godoc
// @Summary      Get supported tokens
// @Description  Get list of tokens supported by the platform
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.Response
// @Failure      500       {object}  response.Response
// @Router       /auctions/supported-tokens [get]
func (h *AuctionHandler) GetSupportedTokens(c *gin.Context) {
	tokens, err := h.service.GetSupportedTokens()
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, tokens)
}

// CheckNFTApproval godoc
// @Summary      Check NFT approval status
// @Description  Check if a specific NFT token is approved for the platform contract
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        request  body      models.CheckNFTApprovalPayload  true  "Check NFT approval payload"
// @Success      200      {object}  response.Response
// @Failure      400      {object}  response.Response
// @Failure      500      {object}  response.Response
// @Router       /auctions/check-nft-approval [post]
func (h *AuctionHandler) CheckNFTApproval(c *gin.Context) {
	var payload models.CheckNFTApprovalPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	isApproved, err := h.service.CheckNFTApproval(payload.NFTAddress, payload.TokenID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, gin.H{
		"nftAddress": payload.NFTAddress,
		"tokenId":    payload.TokenID,
		"approved":   isApproved,
	})
}

// GetUserAuctionHistory godoc
// @Summary      Get user auction history
// @Description  Get auction history records created by the current user with simplified fields
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        page      query     int      false  "Page number" default(1)
// @Param        pageSize  query     int      false  "Page size" default(10)
// @Param        status    query     []string false  "Filter by status (pending, active, ended, cancelled)" collectionFormat(multi)
// @Success      200       {object}  response.Response
// @Failure      400       {object}  response.Response
// @Failure      401       {object}  response.Response
// @Failure      500       {object}  response.Response
// @Security     BearerAuth
// @Router       /auctions/my/history [get]
func (h *AuctionHandler) GetUserAuctionHistory(c *gin.Context) {
	user, err := jwt.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	var query page.PageQuery
	if err := query.Bind(c); err != nil {
		return
	}

	// 获取 status 查询参数（支持多个值）
	statuses := c.QueryArray("status")

	history, total, err := h.service.GetUserAuctionHistory(user.ID, query, statuses)
	if err != nil {
		response.Error(c, err)
		return
	}

	pageData := page.NewPageData(query.Page, query.PageSize, total, history)
	response.Success(c, pageData)
}

// Cancel godoc
// @Summary      Cancel auction
// @Description  Cancel an existing auction (only pending or active auctions can be cancelled)
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Auction ID (string)"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      403  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Security     BearerAuth
// @Router       /auctions/{id}/cancel [post]
func (h *AuctionHandler) Cancel(c *gin.Context) {
	auctionID := c.Param("id")
	if auctionID == "" {
		response.BadRequest(c, "auction id is required")
		return
	}

	user, err := jwt.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	auction, err := h.service.Cancel(user.ID, auctionID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, auction)
}

// GetAuctionSimpleStats godoc
// @Summary      Get auction simple statistics
// @Description  Get simple statistics from contract (total auctions, total bids, platform fee, total value locked)
// @Tags         auctions
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response{data=models.AuctionSimpleStatsResponse}
// @Failure      500  {object}  response.Response
// @Router       /auctions/stats [get]
func (h *AuctionHandler) GetAuctionSimpleStats(c *gin.Context) {
	stats, err := h.service.GetAuctionSimpleStats()
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, stats)
}
