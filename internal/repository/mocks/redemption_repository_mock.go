package mocks

import (
	"github.com/gift-redemption/internal/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockRedemptionRepository struct {
	mock.Mock
}

func (m *MockRedemptionRepository) Create(tx *gorm.DB, redemption *model.Redemption) error {
	args := m.Called(tx, redemption)
	return args.Error(0)
}

func (m *MockRedemptionRepository) FindByUserAndGift(userID, giftID uint) (*model.Redemption, error) {
	args := m.Called(userID, giftID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Redemption), args.Error(1)
}

func (m *MockRedemptionRepository) FindUnratedByUserAndGift(userID, giftID uint) (*model.Redemption, error) {
	args := m.Called(userID, giftID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Redemption), args.Error(1)
}

type MockRatingRepository struct {
	mock.Mock
}

func (m *MockRatingRepository) Create(tx *gorm.DB, rating *model.Rating) error {
	args := m.Called(tx, rating)
	return args.Error(0)
}

func (m *MockRatingRepository) ExistsByRedemption(redemptionID uint) (bool, error) {
	args := m.Called(redemptionID)
	return args.Bool(0), args.Error(1)
}
