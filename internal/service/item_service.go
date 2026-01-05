package service

import (
	"errors"

	"github.com/naufal225/go-simple-login-crud-api/internal/model"
	"github.com/naufal225/go-simple-login-crud-api/internal/repo"
)

type ItemService interface {
	Create(userID, name, sku string, price, stock int) (*model.Item, error)
	List(userID string) ([]model.Item, error)
	Update(userID, itemID, name string, price, stock int) error
	Delete(userID, itemID string) error
}

type itemService struct {
	itemRepo repo.ItemRepository
} 

func (s *itemService) Create(userID, name, sku string, price, stock int) (*model.Item, error) {
	item := &model.Item{
		UserID: userID,
		Name: name,
		SKU: sku,
		Price: price,
		Stock: stock,
	}

	if err := s.itemRepo.Create(item); err != nil {
		return  nil, err
	}

	return item, nil
}

func (s *itemService) List(userID string) ([]model.Item, error) {
	return s.itemRepo.FindByUserID(userID)
}

func (s *itemService) Update(userID, itemID, name string, price, stock int) error {
	item, err := s.itemRepo.FindByID(itemID)
	if err != nil {
		return err
	}

	if item.UserID != userID {
		return errors.New("unauthorized")
	}

	item.Name = name
	item.Price = price
	item.Stock = stock

	return s.itemRepo.Update(item)
}

func (s *itemService) Delete(userID, itemID string) error {
	item, err := s.itemRepo.FindByID(itemID)
	if err != nil {
		return  err
	}

	if item.UserID != userID {
		return  errors.New("unauthorized")
	}

	return s.itemRepo.Delete(itemID)
}

func NewItemService(itemRepo repo.ItemRepository) ItemService {
	return &itemService{itemRepo: itemRepo}
}