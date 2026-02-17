package model

import "time"

type Redemption struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `gorm:"not null;index" json:"user_id"`
	GiftID     uint      `gorm:"not null;index" json:"gift_id"`
	Quantity   int       `gorm:"not null;default:1" json:"quantity"`
	TotalPoint int       `gorm:"not null" json:"total_point"`
	RedeemedAt time.Time `json:"redeemed_at"`
	CreatedAt  time.Time `json:"created_at"`

	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Gift *Gift `gorm:"foreignKey:GiftID" json:"gift,omitempty"`
}

type Rating struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"not null;index" json:"user_id"`
	GiftID       uint      `gorm:"not null;index" json:"gift_id"`
	RedemptionID uint      `gorm:"not null" json:"redemption_id"`
	Score        float64   `gorm:"not null" json:"score"` // 1–5
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	User       *User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Gift       *Gift       `gorm:"foreignKey:GiftID" json:"gift,omitempty"`
	Redemption *Redemption `gorm:"foreignKey:RedemptionID" json:"redemption,omitempty"`
}
