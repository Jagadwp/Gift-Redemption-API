package mocks

import (
	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/repository"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockGiftRepository struct {
	mock.Mock
}

func (m *MockGiftRepository) FindAll(filter repository.GiftFilter) ([]model.Gift, int64, error) {
	args := m.Called(filter)
	return args.Get(0).([]model.Gift), args.Get(1).(int64), args.Error(2)
}

func (m *MockGiftRepository) FindByID(id uint) (*model.Gift, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Gift), args.Error(1)
}

func (m *MockGiftRepository) Create(gift *model.Gift) error {
	args := m.Called(gift)
	return args.Error(0)
}

func (m *MockGiftRepository) Update(gift *model.Gift) error {
	args := m.Called(gift)
	return args.Error(0)
}

func (m *MockGiftRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockGiftRepository) DeductStock(tx *gorm.DB, giftID uint, qty int) error {
	args := m.Called(tx, giftID, qty)
	return args.Error(0)
}

func (m *MockGiftRepository) UpdateRatingStats(tx *gorm.DB, giftID uint) error {
	args := m.Called(tx, giftID)
	return args.Error(0)
}
