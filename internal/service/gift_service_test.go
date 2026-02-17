package service

import (
	"testing"
	"time"

	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/pkg/apperror"
	"github.com/gift-redemption/internal/repository"
	"github.com/gift-redemption/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGiftService_GetAll_Success(t *testing.T) {
	mockGiftRepo := new(mocks.MockGiftRepository)
	giftService := NewGiftService(mockGiftRepo)

	gifts := []model.Gift{
		{
			ID:           1,
			Name:         "Gift 1",
			Point:        100,
			Stock:        10,
			AvgRating:    4.5,
			TotalReviews: 10,
			CreatedAt:    time.Now(),
		},
		{
			ID:           2,
			Name:         "Gift 2",
			Point:        200,
			Stock:        5,
			AvgRating:    3.8,
			TotalReviews: 5,
			CreatedAt:    time.Now(),
		},
	}

	filter := repository.GiftFilter{
		Page:    1,
		Limit:   10,
		SortBy:  "created_at",
		SortDir: "desc",
	}

	mockGiftRepo.On("FindAll", filter).Return(gifts, int64(2), nil)

	query := dto.PaginationQuery{
		Page:    1,
		Limit:   10,
		SortBy:  "created_at",
		SortDir: "desc",
	}

	result, pagination, err := giftService.GetAll(query)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Gift 1", result[0].Name)
	assert.Equal(t, 4.5, result[0].StarRating) // 4.5 rounded to nearest 0.5
	assert.NotNil(t, pagination)
	assert.Equal(t, 1, pagination.CurrentPage)
	assert.Equal(t, int64(2), pagination.Total)
	mockGiftRepo.AssertExpectations(t)
}

func TestGiftService_GetByID_Success(t *testing.T) {
	mockGiftRepo := new(mocks.MockGiftRepository)
	giftService := NewGiftService(mockGiftRepo)

	gift := &model.Gift{
		ID:           1,
		Name:         "Test Gift",
		Point:        100,
		Stock:        10,
		AvgRating:    3.6, // should round to 3.5
		TotalReviews: 10,
		CreatedAt:    time.Now(),
	}

	mockGiftRepo.On("FindByID", uint(1)).Return(gift, nil)

	result, err := giftService.GetByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Gift", result.Name)
	assert.Equal(t, 3.6, result.AvgRating)
	assert.Equal(t, 3.5, result.StarRating) // rounded
	assert.True(t, result.InStock)
	mockGiftRepo.AssertExpectations(t)
}

func TestGiftService_GetByID_NotFound(t *testing.T) {
	mockGiftRepo := new(mocks.MockGiftRepository)
	giftService := NewGiftService(mockGiftRepo)

	mockGiftRepo.On("FindByID", uint(999)).Return(nil, apperror.ErrNotFound)

	result, err := giftService.GetByID(999)

	assert.Error(t, err)
	assert.Equal(t, apperror.ErrNotFound, err)
	assert.Nil(t, result)
	mockGiftRepo.AssertExpectations(t)
}

func TestGiftService_Create_Success(t *testing.T) {
	mockGiftRepo := new(mocks.MockGiftRepository)
	giftService := NewGiftService(mockGiftRepo)

	req := dto.CreateGiftRequest{
		Name:         "New Gift",
		Description:  "Test description",
		Point:        100,
		Stock:        10,
		IsNew:        true,
		IsBestSeller: false,
	}

	mockGiftRepo.On("Create", mock.MatchedBy(func(g *model.Gift) bool {
		return g.Name == req.Name && g.Point == req.Point
	})).Return(nil)

	result, err := giftService.Create(req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Name, result.Name)
	mockGiftRepo.AssertExpectations(t)
}

func TestGiftService_Patch_Success(t *testing.T) {
	mockGiftRepo := new(mocks.MockGiftRepository)
	giftService := NewGiftService(mockGiftRepo)

	existingGift := &model.Gift{
		ID:           1,
		Name:         "Original Name",
		Description:  "Original Desc",
		Point:        100,
		Stock:        10,
		IsNew:        false,
		IsBestSeller: false,
	}

	newStock := 20
	isNew := true
	req := dto.PatchGiftRequest{
		Stock: &newStock,
		IsNew: &isNew,
	}

	mockGiftRepo.On("FindByID", uint(1)).Return(existingGift, nil)
	mockGiftRepo.On("Update", existingGift).Return(nil)

	result, err := giftService.Patch(1, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Original Name", result.Name) // unchanged
	assert.Equal(t, 20, result.Stock)             // changed
	assert.True(t, result.IsNew)                  // changed
	mockGiftRepo.AssertExpectations(t)
}

func TestGiftService_StarRating_Rounding(t *testing.T) {
	tests := []struct {
		name       string
		avgRating  float64
		wantRating float64
	}{
		{"3.2 rounds to 3.0", 3.2, 3.0},
		{"3.6 rounds to 3.5", 3.6, 3.5},
		{"3.9 rounds to 4.0", 3.9, 4.0},
		{"4.25 rounds to 4.5", 4.25, 4.5},
		{"4.74 rounds to 4.5", 4.74, 4.5},
		{"4.75 rounds to 5.0", 4.75, 5.0},
	}

	mockGiftRepo := new(mocks.MockGiftRepository)
	giftService := NewGiftService(mockGiftRepo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gift := &model.Gift{
				ID:        1,
				Name:      "Test",
				AvgRating: tt.avgRating,
			}

			mockGiftRepo.On("FindByID", uint(1)).Return(gift, nil).Once()

			result, err := giftService.GetByID(1)

			assert.NoError(t, err)
			assert.Equal(t, tt.wantRating, result.StarRating)
		})
	}

	mockGiftRepo.AssertExpectations(t)
}
