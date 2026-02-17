package service

import (
	"math"

	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/pkg/response"
	"github.com/gift-redemption/internal/repository"
)

type GiftService interface {
	GetAll(query dto.PaginationQuery) ([]dto.GiftResponse, *response.Pagination, error)
	GetByID(id uint) (*dto.GiftResponse, error)
	Create(req dto.CreateGiftRequest) (*dto.GiftResponse, error)
	Update(id uint, req dto.UpdateGiftRequest) (*dto.GiftResponse, error)
	Patch(id uint, req dto.PatchGiftRequest) (*dto.GiftResponse, error)
	Delete(id uint) error
}

type giftService struct {
	giftRepo repository.GiftRepository
}

func NewGiftService(giftRepo repository.GiftRepository) GiftService {
	return &giftService{giftRepo}
}

func (s *giftService) GetAll(query dto.PaginationQuery) ([]dto.GiftResponse, *response.Pagination, error) {
	query.Normalize()

	filter := repository.GiftFilter{
		Page:    query.Page,
		Limit:   query.Limit,
		SortBy:  query.SortBy,
		SortDir: query.SortDir,
	}

	gifts, total, err := s.giftRepo.FindAll(filter)
	if err != nil {
		return nil, nil, err
	}

	result := make([]dto.GiftResponse, len(gifts))
	for i, g := range gifts {
		result[i] = dto.ToGiftResponse(g)
	}

	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))
	pagination := &response.Pagination{
		CurrentPage: query.Page,
		PerPage:     query.Limit,
		Total:       total,
		TotalPages:  totalPages,
	}

	return result, pagination, nil
}

func (s *giftService) GetByID(id uint) (*dto.GiftResponse, error) {
	gift, err := s.giftRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	res := dto.ToGiftResponse(*gift)
	return &res, nil
}

func (s *giftService) Create(req dto.CreateGiftRequest) (*dto.GiftResponse, error) {
	gift := &model.Gift{
		Name:         req.Name,
		Description:  req.Description,
		Point:        req.Point,
		Stock:        req.Stock,
		ImageURL:     req.ImageURL,
		IsNew:        req.IsNew,
		IsBestSeller: req.IsBestSeller,
	}

	if err := s.giftRepo.Create(gift); err != nil {
		return nil, err
	}

	res := dto.ToGiftResponse(*gift)
	return &res, nil
}

func (s *giftService) Update(id uint, req dto.UpdateGiftRequest) (*dto.GiftResponse, error) {
	gift, err := s.giftRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	gift.Name = req.Name
	gift.Description = req.Description
	gift.Point = req.Point
	gift.Stock = req.Stock
	gift.ImageURL = req.ImageURL
	gift.IsNew = req.IsNew
	gift.IsBestSeller = req.IsBestSeller

	if err := s.giftRepo.Update(gift); err != nil {
		return nil, err
	}

	res := dto.ToGiftResponse(*gift)
	return &res, nil
}

func (s *giftService) Patch(id uint, req dto.PatchGiftRequest) (*dto.GiftResponse, error) {
	gift, err := s.giftRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// only update fields that are explicitly provided
	if req.Name != nil {
		gift.Name = *req.Name
	}
	if req.Description != nil {
		gift.Description = *req.Description
	}
	if req.Point != nil {
		gift.Point = *req.Point
	}
	if req.Stock != nil {
		gift.Stock = *req.Stock
	}
	if req.ImageURL != nil {
		gift.ImageURL = *req.ImageURL
	}
	if req.IsNew != nil {
		gift.IsNew = *req.IsNew
	}
	if req.IsBestSeller != nil {
		gift.IsBestSeller = *req.IsBestSeller
	}

	if err := s.giftRepo.Update(gift); err != nil {
		return nil, err
	}

	res := dto.ToGiftResponse(*gift)
	return &res, nil
}

func (s *giftService) Delete(id uint) error {
	return s.giftRepo.Delete(id)
}
