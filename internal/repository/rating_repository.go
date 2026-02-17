package repository

import (
	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/pkg/apperror"
	"gorm.io/gorm"
)

type RatingRepository interface {
	Create(tx *gorm.DB, rating *model.Rating) error
	ExistsByRedemption(redemptionID uint) (bool, error)
}

type ratingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) RatingRepository {
	return &ratingRepository{db}
}

func (r *ratingRepository) Create(tx *gorm.DB, rating *model.Rating) error {
	err := tx.Create(rating).Error
	if err != nil && isDuplicateError(err) {
		return apperror.ErrAlreadyRated
	}
	return err
}

func (r *ratingRepository) ExistsByRedemption(redemptionID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.Rating{}).
		Where("redemption_id = ?", redemptionID).
		Count(&count).Error
	return count > 0, err
}
