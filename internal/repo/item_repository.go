package repo

import (
	"github.com/naufal225/go-simple-login-crud-api/internal/model"
	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *model.Item) error
	FindByID(id string) (*model.Item, error)
	FindByUserID(userID string) ([]model.Item, error)
	Update(item *model.Item) error
	Delete(id string) error
}

type itemRepository struct {
	db *gorm.DB
}

func (r *itemRepository) Create(item *model.Item) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) FindByID(id string) (*model.Item, error) {
	var item model.Item
	err := r.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) FindByUserID(id string) ([]model.Item, error) {
	var items []model.Item
	err := r.db.Where("user_id = ?", id).Find(&items).Error
	return items, err
}

func (r *itemRepository) Update(item *model.Item) error {
	return r.db.Save(&item).Error
}

func (r *itemRepository) Delete(id string) error {
	return r.db.Delete(&model.Item{}, "id = ?", id).Error
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}
