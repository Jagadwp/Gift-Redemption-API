package handler

import (
	"errors"

	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/pkg/apperror"
	"github.com/gift-redemption/internal/pkg/response"
	"github.com/gift-redemption/internal/service"
	"github.com/gin-gonic/gin"
)

type GiftHandler struct {
	giftService service.GiftService
}

func NewGiftHandler(giftService service.GiftService) *GiftHandler {
	return &GiftHandler{giftService}
}

// GetGifts godoc
// @Summary      Get all gifts
// @Description  Returns paginated list of gifts with sorting options
// @Tags         Gifts
// @Produce      json
// @Security     BearerAuth
// @Param        page      query     int     false  "Page number (default: 1)"
// @Param        limit     query     int     false  "Items per page (default: 10, max: 100)"
// @Param        sort_by   query     string  false  "Sort field: created_at | avg_rating (default: created_at)"
// @Param        sort_dir  query     string  false  "Sort direction: asc | desc (default: desc)"
// @Success      200       {object}  response.envelope{data=[]dto.GiftResponse}
// @Failure      500       {object}  response.envelope
// @Router       /gifts [get]
func (h *GiftHandler) GetAll(c *gin.Context) {
	var query dto.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.BadRequest(c, "invalid query parameters", err.Error())
		return
	}

	gifts, pagination, err := h.giftService.GetAll(query)
	if err != nil {
		response.InternalServerError(c, "failed to fetch gifts")
		return
	}

	response.SuccessPaginated(c, "gifts retrieved successfully", gifts, pagination)
}

// GetGift godoc
// @Summary      Get gift by ID
// @Description  Returns a single gift with star rating
// @Tags         Gifts
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Gift ID"
// @Success      200  {object}  response.envelope{data=dto.GiftResponse}
// @Failure      404  {object}  response.envelope
// @Router       /gifts/{id} [get]
func (h *GiftHandler) GetByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	gift, err := h.giftService.GetByID(id)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			response.NotFound(c, "gift not found")
			return
		}
		response.InternalServerError(c, "failed to fetch gift")
		return
	}

	response.Success(c, "gift retrieved successfully", gift)
}

// CreateGift godoc
// @Summary      Create gift
// @Description  Create a new gift item (admin only)
// @Tags         Gifts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      dto.CreateGiftRequest  true  "Gift data"
// @Success      201   {object}  response.envelope{data=dto.GiftResponse}
// @Failure      400   {object}  response.envelope
// @Router       /gifts [post]
func (h *GiftHandler) Create(c *gin.Context) {
	var req dto.CreateGiftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body", err.Error())
		return
	}

	gift, err := h.giftService.Create(req)
	if err != nil {
		response.InternalServerError(c, "failed to create gift")
		return
	}

	response.Created(c, "gift created successfully", gift)
}

// UpdateGift godoc
// @Summary      Update gift
// @Description  Full update of a gift item (admin only)
// @Tags         Gifts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                    true  "Gift ID"
// @Param        body  body      dto.UpdateGiftRequest  true  "Gift data"
// @Success      200   {object}  response.envelope{data=dto.GiftResponse}
// @Failure      400   {object}  response.envelope
// @Failure      404   {object}  response.envelope
// @Router       /gifts/{id} [put]
func (h *GiftHandler) Update(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	var req dto.UpdateGiftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body", err.Error())
		return
	}

	gift, err := h.giftService.Update(id, req)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			response.NotFound(c, "gift not found")
			return
		}
		response.InternalServerError(c, "failed to update gift")
		return
	}

	response.Success(c, "gift updated successfully", gift)
}

// PatchGift godoc
// @Summary      Patch gift
// @Description  Partial update of a gift item (admin only)
// @Tags         Gifts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                   true  "Gift ID"
// @Param        body  body      dto.PatchGiftRequest  true  "Partial gift data"
// @Success      200   {object}  response.envelope{data=dto.GiftResponse}
// @Failure      400   {object}  response.envelope
// @Failure      404   {object}  response.envelope
// @Router       /gifts/{id} [patch]
func (h *GiftHandler) Patch(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	var req dto.PatchGiftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body", err.Error())
		return
	}

	gift, err := h.giftService.Patch(id, req)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			response.NotFound(c, "gift not found")
			return
		}
		response.InternalServerError(c, "failed to patch gift")
		return
	}

	response.Success(c, "gift updated successfully", gift)
}

// DeleteGift godoc
// @Summary      Delete gift
// @Description  Soft delete a gift item (admin only)
// @Tags         Gifts
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Gift ID"
// @Success      200  {object}  response.envelope
// @Failure      404  {object}  response.envelope
// @Router       /gifts/{id} [delete]
func (h *GiftHandler) Delete(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	if err := h.giftService.Delete(id); err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			response.NotFound(c, "gift not found")
			return
		}
		response.InternalServerError(c, "failed to delete gift")
		return
	}

	response.Success(c, "gift deleted successfully", nil)
}
