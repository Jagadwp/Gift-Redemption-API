package model

import (
	"math"
	"time"

	"gorm.io/gorm"
)

type Gift struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"not null" json:"name"`
	Description  string         `json:"description"`
	Point        int            `gorm:"not null;default:0" json:"point"`
	Stock        int            `gorm:"not null;default:0" json:"stock"`
	ImageURL     string         `json:"image_url"`
	IsNew        bool           `gorm:"default:false" json:"is_new"`
	IsBestSeller bool           `gorm:"default:false" json:"is_best_seller"`
	AvgRating    float64        `gorm:"default:0" json:"avg_rating"`
	TotalReviews int            `gorm:"default:0" json:"total_reviews"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// StarRating returns avg_rating rounded to nearest 0.5
// e.g. 3.2 → 3.0, 3.6 → 3.5, 3.9 → 4.0
func (g *Gift) StarRating() float64 {
	return math.Round(g.AvgRating*2) / 2
}

func (g *Gift) InStock() bool {
	return g.Stock > 0
}