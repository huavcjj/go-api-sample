package repository

import (
	"go-api-sample/domain/entity"

	"gorm.io/gorm"
)

type (
	CategoryRepository interface {
		GetOrCreate(category *entity.Category) (*entity.Category, error)
	}
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) GetOrCreate(category *entity.Category) (*entity.Category, error) {
	var getOrCreateCategory entity.Category
	if err := r.db.FirstOrCreate(&getOrCreateCategory, category).Error; err != nil {
		return nil, err
	}
	return &getOrCreateCategory, nil
}
