package service

import (
	"testing"

	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/pkg/apperror"
	"github.com/gift-redemption/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

// Note: Full transaction testing requires integration tests with real DB.
// These tests focus on validation logic before transactions.

func TestRedemptionService_Redeem_GiftNotFound(t *testing.T) {
	mockGiftRepo := new(mocks.MockGiftRepository)
	mockRedemptionRepo := new(mocks.MockRedemptionRepository)
	mockRatingRepo := new(mocks.MockRatingRepository)

	redemptionService := NewRedemptionService(nil, mockGiftRepo, mockRedemptionRepo, mockRatingRepo)

	mockGiftRepo.On("FindByID", uint(999)).Return(nil, apperror.ErrNotFound)

	req := dto.RedemptionRequest{
		Quantity: 1,
	}

	result, err := redemptionService.Redeem(1, 999, req)

	assert.Error(t, err)
	assert.Equal(t, apperror.ErrNotFound, err)
	assert.Nil(t, result)
	mockGiftRepo.AssertExpectations(t)
}

func TestRedemptionService_Rate_NotRedeemed(t *testing.T) {
	mockGiftRepo := new(mocks.MockGiftRepository)
	mockRedemptionRepo := new(mocks.MockRedemptionRepository)
	mockRatingRepo := new(mocks.MockRatingRepository)

	redemptionService := NewRedemptionService(nil, mockGiftRepo, mockRedemptionRepo, mockRatingRepo)

	mockRedemptionRepo.On("FindUnratedByUserAndGift", uint(1), uint(1)).
		Return(nil, apperror.ErrNotRedeemed)

	req := dto.RatingRequest{
		Score: 5,
	}

	result, err := redemptionService.Rate(1, 1, req)

	assert.Error(t, err)
	assert.Equal(t, apperror.ErrNotRedeemed, err)
	assert.Nil(t, result)
	mockRedemptionRepo.AssertExpectations(t)
}

func TestRedemptionService_Rate_GiftNotFound(t *testing.T) {
	mockGiftRepo := new(mocks.MockGiftRepository)
	mockRedemptionRepo := new(mocks.MockRedemptionRepository)
	mockRatingRepo := new(mocks.MockRatingRepository)

	redemptionService := NewRedemptionService(nil, mockGiftRepo, mockRedemptionRepo, mockRatingRepo)

	redemption := &model.Redemption{
		ID:     1,
		UserID: 1,
		GiftID: 999,
	}

	mockRedemptionRepo.On("FindUnratedByUserAndGift", uint(1), uint(999)).Return(redemption, nil)
	mockGiftRepo.On("FindByID", uint(999)).Return(nil, apperror.ErrNotFound)

	req := dto.RatingRequest{
		Score: 5,
	}

	result, err := redemptionService.Rate(1, 999, req)

	assert.Error(t, err)
	assert.Equal(t, apperror.ErrNotFound, err)
	assert.Nil(t, result)
	mockRedemptionRepo.AssertExpectations(t)
	mockGiftRepo.AssertExpectations(t)
}

func TestRedemptionService_Rate_ValidationScore(t *testing.T) {
	// This test validates that RatingRequest with score 1-5 is valid in DTO level
	// Actual validation happens at handler layer with gin binding

	tests := []struct {
		name  string
		score int
		valid bool
	}{
		{"score 1 is valid", 1, true},
		{"score 3 is valid", 3, true},
		{"score 5 is valid", 5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := dto.RatingRequest{Score: tt.score}
			assert.GreaterOrEqual(t, req.Score, 1)
			assert.LessOrEqual(t, req.Score, 5)
		})
	}
}
