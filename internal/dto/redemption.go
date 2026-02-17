package dto

import (
	"time"

	"github.com/gift-redemption/internal/model"
)

type RedemptionRequest struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

type RedemptionResponse struct {
	RedemptionID uint   `json:"redemption_id"`
	GiftID       uint   `json:"gift_id"`
	GiftName     string `json:"gift_name"`
	Quantity     int    `json:"quantity"`
	TotalPoint   int    `json:"total_point"`
	RedeemedAt   string `json:"redeemed_at"`
}

func ToRedemptionResponse(r model.Redemption, giftName string) RedemptionResponse {
	return RedemptionResponse{
		RedemptionID: r.ID,
		GiftID:       r.GiftID,
		GiftName:     giftName,
		Quantity:     r.Quantity,
		TotalPoint:   r.TotalPoint,
		RedeemedAt:   r.RedeemedAt.Format(time.RFC3339),
	}
}