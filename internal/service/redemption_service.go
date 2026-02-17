package service

import (
	"fmt"
	"time"

	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/repository"
	"gorm.io/gorm"
)

type RedemptionService interface {
	Redeem(userID, giftID uint, req dto.RedemptionRequest) (*dto.RedemptionResponse, error)
	Rate(userID, giftID uint, req dto.RatingRequest) (*dto.RatingResponse, error)
}

type redemptionService struct {
	db             *gorm.DB
	giftRepo       repository.GiftRepository
	redemptionRepo repository.RedemptionRepository
	ratingRepo     repository.RatingRepository
}

func NewRedemptionService(
	db *gorm.DB,
	giftRepo repository.GiftRepository,
	redemptionRepo repository.RedemptionRepository,
	ratingRepo repository.RatingRepository,
) RedemptionService {
	return &redemptionService{db, giftRepo, redemptionRepo, ratingRepo}
}

func (s *redemptionService) Redeem(userID, giftID uint, req dto.RedemptionRequest) (*dto.RedemptionResponse, error) {
	// check gift exists before opening transaction
	gift, err := s.giftRepo.FindByID(giftID)
	if err != nil {
		return nil, err
	}

	var redemption *model.Redemption

	err = repository.WithTransaction(s.db, func(tx *gorm.DB) error {
		// deduct stock with row lock inside transaction
		if err := s.giftRepo.DeductStock(tx, giftID, req.Quantity); err != nil {
			return err
		}

		redemption = &model.Redemption{
			UserID:     userID,
			GiftID:     giftID,
			Quantity:   req.Quantity,
			TotalPoint: gift.Point * req.Quantity,
		}

		return s.redemptionRepo.Create(tx, redemption)
	})

	if err != nil {
		return nil, err
	}

	return &dto.RedemptionResponse{
		RedemptionID: redemption.ID,
		GiftID:       gift.ID,
		GiftName:     gift.Name,
		Quantity:     redemption.Quantity,
		TotalPoint:   redemption.TotalPoint,
		RedeemedAt:   redemption.RedeemedAt.Format(time.RFC3339),
	}, nil
}

func (s *redemptionService) Rate(userID, giftID uint, req dto.RatingRequest) (*dto.RatingResponse, error) {
	// validate user has an unrated redemption for this gift
	redemption, err := s.redemptionRepo.FindUnratedByUserAndGift(userID, giftID)
	if err != nil {
		return nil, err
	}

	gift, err := s.giftRepo.FindByID(giftID)
	if err != nil {
		return nil, err
	}

	err = repository.WithTransaction(s.db, func(tx *gorm.DB) error {
		rating := &model.Rating{
			UserID:       userID,
			GiftID:       giftID,
			RedemptionID: redemption.ID,
			Score:        req.Score,
		}

		if err := s.ratingRepo.Create(tx, rating); err != nil {
			return fmt.Errorf("create rating: %w", err)
		}

		// recalculate avg_rating and total_reviews in gifts table
		return s.giftRepo.UpdateRatingStats(tx, giftID)
	})

	if err != nil {
		return nil, err
	}

	// fetch updated gift for fresh avg_rating
	updatedGift, err := s.giftRepo.FindByID(giftID)
	if err != nil {
		return nil, err
	}

	return &dto.RatingResponse{
		GiftID:     gift.ID,
		GiftName:   gift.Name,
		Score:      req.Score,
		AvgRating:  updatedGift.AvgRating,
		StarRating: updatedGift.StarRating(),
	}, nil
}
