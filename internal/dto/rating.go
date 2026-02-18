package dto

import (
	"math"

	"github.com/gift-redemption/internal/model"
)

type RatingRequest struct {
	Score float64 `json:"score" binding:"required,min=1,max=5"`
}

type RatingResponse struct {
	ID         uint    `json:"id"`
	GiftID     uint    `json:"gift_id"`
	GiftName   string  `json:"gift_name"`
	Score      float64 `json:"score"`
	AvgRating  float64 `json:"avg_rating"`
	StarRating float64 `json:"star_rating"`
}

// StarRating returns avg_rating rounded to nearest 0.5
// e.g. 3.2 → 3.0, 3.6 → 3.5, 3.9 → 4.0
func RoundToHalf(v float64) float64 {
	return math.Round(v*2) / 2
}

func ToRatingResponse(r model.Rating, gift model.Gift) RatingResponse {
	return RatingResponse{
		ID:         r.ID,
		GiftID:     gift.ID,
		GiftName:   gift.Name,
		Score:      RoundToHalf(r.Score),
		AvgRating:  gift.AvgRating,
		StarRating: RoundToHalf(gift.AvgRating),
	}
}
