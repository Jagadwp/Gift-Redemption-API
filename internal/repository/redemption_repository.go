package repository

import (
	"errors"
	"time"

	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/pkg/apperror"
	"gorm.io/gorm"
)

type RedemptionRepository interface {
	Create(tx *gorm.DB, redemption *model.Redemption) error
	FindByUserAndGift(userID, giftID uint) (*model.Redemption, error)
	FindUnratedByUserAndGift(userID, giftID uint) (*model.Redemption, error)
}

type redemptionRepository struct {
	db *gorm.DB
}

func NewRedemptionRepository(db *gorm.DB) RedemptionRepository {
	return &redemptionRepository{db}
}

func (r *redemptionRepository) Create(tx *gorm.DB, redemption *model.Redemption) error {
	redemption.RedeemedAt = time.Now()
	return tx.Create(redemption).Error
}

func (r *redemptionRepository) FindByUserAndGift(userID, giftID uint) (*model.Redemption, error) {
	var redemption model.Redemption
	err := r.db.
		Where("user_id = ? AND gift_id = ?", userID, giftID).
		First(&redemption).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return &redemption, err
}

// FindUnratedByUserAndGift returns a redemption that has not been rated yet
func (r *redemptionRepository) FindUnratedByUserAndGift(userID, giftID uint) (*model.Redemption, error) {
	var redemption model.Redemption
	err := r.db.
		Where("user_id = ? AND gift_id = ?", userID, giftID).
		Where("id NOT IN (SELECT redemption_id FROM ratings WHERE user_id = ? AND gift_id = ?)", userID, giftID).
		First(&redemption).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotRedeemed
	}
	return &redemption, err
}
