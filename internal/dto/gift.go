package dto

import (
	"time"

	"github.com/gift-redemption/internal/model"
)

type CreateGiftRequest struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	Point        int    `json:"point" binding:"required,min=1"`
	Stock        int    `json:"stock" binding:"min=0"`
	ImageURL     string `json:"image_url"`
	IsNew        bool   `json:"is_new"`
	IsBestSeller bool   `json:"is_best_seller"`
}

type UpdateGiftRequest struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	Point        int    `json:"point" binding:"required,min=1"`
	Stock        int    `json:"stock" binding:"min=0"`
	ImageURL     string `json:"image_url"`
	IsNew        bool   `json:"is_new"`
	IsBestSeller bool   `json:"is_best_seller"`
}

type PatchGiftRequest struct {
	Name         *string `json:"name"`
	Description  *string `json:"description"`
	Point        *int    `json:"point"`
	Stock        *int    `json:"stock"`
	ImageURL     *string `json:"image_url"`
	IsNew        *bool   `json:"is_new"`
	IsBestSeller *bool   `json:"is_best_seller"`
}

type GiftResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Point        int     `json:"point"`
	Stock        int     `json:"stock"`
	ImageURL     string  `json:"image_url"`
	IsNew        bool    `json:"is_new"`
	IsBestSeller bool    `json:"is_best_seller"`
	AvgRating    float64 `json:"avg_rating"`
	StarRating   float64 `json:"star_rating"`
	TotalReviews int     `json:"total_reviews"`
	InStock      bool    `json:"in_stock"`
	CreatedAt    string  `json:"created_at"`
}

func ToGiftResponse(g model.Gift) GiftResponse {
	return GiftResponse{
		ID:           g.ID,
		Name:         g.Name,
		Description:  g.Description,
		Point:        g.Point,
		Stock:        g.Stock,
		ImageURL:     g.ImageURL,
		IsNew:        g.IsNew,
		IsBestSeller: g.IsBestSeller,
		AvgRating:    g.AvgRating,
		StarRating:   g.StarRating(),
		TotalReviews: g.TotalReviews,
		InStock:      g.InStock(),
		CreatedAt:    g.CreatedAt.Format(time.RFC3339),
	}
}
