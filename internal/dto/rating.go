package dto

import "github.com/gift-redemption/internal/model"

type RatingRequest struct {
	Score int `json:"score" binding:"required,min=1,max=5"`
}

type RatingResponse struct {
	ID         uint    `json:"id"`
	GiftID     uint    `json:"gift_id"`
	GiftName   string  `json:"gift_name"`
	Score      int     `json:"score"`
	AvgRating  float64 `json:"avg_rating"`
	StarRating float64 `json:"star_rating"`
}

func ToRatingResponse(r model.Rating, gift model.Gift) RatingResponse {
	return RatingResponse{
		ID:         r.ID,
		GiftID:     gift.ID,
		GiftName:   gift.Name,
		Score:      r.Score,
		AvgRating:  gift.AvgRating,
		StarRating: gift.StarRating(),
	}
}
