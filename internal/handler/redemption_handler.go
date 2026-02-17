package handler

import (
	"errors"

	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/middleware"
	"github.com/gift-redemption/internal/pkg/apperror"
	"github.com/gift-redemption/internal/pkg/response"
	"github.com/gift-redemption/internal/service"
	"github.com/gin-gonic/gin"
)

type RedemptionHandler struct {
	redemptionService service.RedemptionService
}

func NewRedemptionHandler(redemptionService service.RedemptionService) *RedemptionHandler {
	return &RedemptionHandler{redemptionService}
}

// RedeemGift godoc
// @Summary      Redeem a gift
// @Description  Redeem a gift item. Stock must be available. Supports quantity > 1.
// @Tags         Gifts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                     true  "Gift ID"
// @Param        body  body      dto.RedemptionRequest   true  "Redemption data"
// @Success      201   {object}  response.envelope{data=dto.RedemptionResponse}
// @Failure      400   {object}  response.envelope
// @Failure      404   {object}  response.envelope
// @Failure      422   {object}  response.envelope  "Insufficient stock"
// @Router       /gifts/{id}/redeem [post]
func (h *RedemptionHandler) Redeem(c *gin.Context) {
	giftID, err := parseID(c, "id")
	if err != nil {
		return
	}

	var req dto.RedemptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body", err.Error())
		return
	}

	userID := middleware.GetUserID(c)

	result, err := h.redemptionService.Redeem(userID, giftID, req)
	if err != nil {
		switch {
		case errors.Is(err, apperror.ErrNotFound):
			response.NotFound(c, "gift not found")
		case errors.Is(err, apperror.ErrInsufficientStock):
			response.UnprocessableEntity(c, "insufficient stock", nil)
		default:
			response.InternalServerError(c, "failed to redeem gift")
		}
		return
	}

	response.Created(c, "gift redeemed successfully", result)
}

// RateGift godoc
// @Summary      Rate a gift
// @Description  Give a rating (1-5) to a gift that has been redeemed. One rating per redemption.
// @Tags         Gifts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int               true  "Gift ID"
// @Param        body  body      dto.RatingRequest  true  "Rating data"
// @Success      201   {object}  response.envelope{data=dto.RatingResponse}
// @Failure      400   {object}  response.envelope
// @Failure      422   {object}  response.envelope  "Not redeemed or already rated"
// @Router       /gifts/{id}/rating [post]
func (h *RedemptionHandler) Rate(c *gin.Context) {
	giftID, err := parseID(c, "id")
	if err != nil {
		return
	}

	var req dto.RatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body", err.Error())
		return
	}

	userID := middleware.GetUserID(c)

	result, err := h.redemptionService.Rate(userID, giftID, req)
	if err != nil {
		switch {
		case errors.Is(err, apperror.ErrNotFound):
			response.NotFound(c, "gift not found")
		case errors.Is(err, apperror.ErrNotRedeemed):
			response.UnprocessableEntity(c, "you have not redeemed this gift or have already rated it", nil)
		case errors.Is(err, apperror.ErrAlreadyRated):
			response.UnprocessableEntity(c, "you have already rated this redemption", nil)
		default:
			response.InternalServerError(c, "failed to submit rating")
		}
		return
	}

	response.Created(c, "rating submitted successfully", result)
}
