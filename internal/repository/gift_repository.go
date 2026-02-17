package repository

import (
	"errors"

	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/pkg/apperror"
	"gorm.io/gorm"
)

type GiftFilter struct {
	Page    int
	Limit   int
	SortBy  string // "created_at" | "avg_rating"
	SortDir string // "asc" | "desc"
}

type GiftRepository interface {
	FindAll(filter GiftFilter) ([]model.Gift, int64, error)
	FindByID(id uint) (*model.Gift, error)
	Create(gift *model.Gift) error
	Update(gift *model.Gift) error
	Delete(id uint) error
	// DeductStock reduces stock atomically inside an existing transaction
	DeductStock(tx *gorm.DB, giftID uint, qty int) error
	UpdateRatingStats(tx *gorm.DB, giftID uint) error
}

type giftRepository struct {
	db *gorm.DB
}

func NewGiftRepository(db *gorm.DB) GiftRepository {
	return &giftRepository{db}
}

func (r *giftRepository) FindAll(filter GiftFilter) ([]model.Gift, int64, error) {
	var gifts []model.Gift
	var total int64

	query := r.db.Model(&model.Gift{})

	// count before pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sortBy := "created_at"
	if filter.SortBy == "avg_rating" {
		sortBy = "avg_rating"
	}

	sortDir := "DESC"
	if filter.SortDir == "asc" {
		sortDir = "ASC"
	}

	offset := (filter.Page - 1) * filter.Limit

	err := query.
		Order(sortBy + " " + sortDir).
		Limit(filter.Limit).
		Offset(offset).
		Find(&gifts).Error

	return gifts, total, err
}

func (r *giftRepository) FindByID(id uint) (*model.Gift, error) {
	var gift model.Gift
	err := r.db.First(&gift, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return &gift, err
}

func (r *giftRepository) Create(gift *model.Gift) error {
	return r.db.Create(gift).Error
}

func (r *giftRepository) Update(gift *model.Gift) error {
	return r.db.Save(gift).Error
}

func (r *giftRepository) Delete(id uint) error {
	result := r.db.Delete(&model.Gift{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperror.ErrNotFound
	}
	return nil
}

// DeductStock uses SELECT FOR UPDATE to prevent race condition on stock
func (r *giftRepository) DeductStock(tx *gorm.DB, giftID uint, qty int) error {
	var gift model.Gift

	err := tx.Set("gorm:query_option", "FOR UPDATE").
		First(&gift, giftID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.ErrNotFound
	}
	if err != nil {
		return err
	}

	if gift.Stock < qty {
		return apperror.ErrInsufficientStock
	}

	return tx.Model(&gift).Update("stock", gift.Stock-qty).Error
}

// UpdateRatingStats recalculates avg_rating and total_reviews from ratings table
func (r *giftRepository) UpdateRatingStats(tx *gorm.DB, giftID uint) error {
	return tx.Exec(`
		UPDATE gifts
		SET avg_rating   = (SELECT COALESCE(AVG(score), 0) FROM ratings WHERE gift_id = ?),
		    total_reviews = (SELECT COUNT(*) FROM ratings WHERE gift_id = ?),
		    updated_at   = NOW()
		WHERE id = ?
	`, giftID, giftID, giftID).Error
}
