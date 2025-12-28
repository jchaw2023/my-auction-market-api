package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"my-auction-market-api/internal/page"
	"my-auction-market-api/internal/response"
	"my-auction-market-api/internal/services"
)

type BidHandler struct {
	bidService     *services.BidService
	auctionService *services.AuctionService
}

func NewBidHandler(bidService *services.BidService, auctionService *services.AuctionService) *BidHandler {
	return &BidHandler{
		bidService:     bidService,
		auctionService: auctionService,
	}
}

// GetBidsByAuctionID godoc
// @Summary      Get bids by auction ID
// @Description  Get all bids for a specific auction (only wallet address, no full user info)
// @Tags         bids
// @Accept       json
// @Produce      json
// @Param        id         path      string  true   "Auction ID"
// @Param        page       query     int     false  "Page number" default(1)
// @Param        pageSize   query     int     false  "Page size" default(10)
// @Success      200        {object}  response.Response
// @Failure      400        {object}  response.Response
// @Failure      500        {object}  response.Response
// @Router       /auctions/{id}/bids [get]
func (h *BidHandler) GetBidsByAuctionID(c *gin.Context) {
	auctionID := c.Param("id")
	if auctionID == "" {
		response.BadRequest(c, "auction id is required")
		return
	}

	var query page.PageQuery
	if err := query.Bind(c); err != nil {
		return
	}

	bidDetails, total, err := h.bidService.GetBidsByAuctionID(auctionID, query)
	if err != nil {
		response.Error(c, err)
		return
	}

	pageData := page.NewPageData(query.Page, query.PageSize, total, bidDetails)
	response.Success(c, pageData)
}

// GetByID godoc
// @Summary      Get bid by ID
// @Description  Get a single bid by ID
// @Tags         bids
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Bid ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /bids/{id} [get]
func (h *BidHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid bid id")
		return
	}

	bid, err := h.bidService.GetByID(id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, bid)
}
