package service

import (
	"go-api-sample/domain/entity"
	"go-api-sample/infrastructure/repository"
)

type (
	CategoryService interface {
		GetOrCreate(category *entity.Category) (*entity.Category, error)
	}
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *categoryService) GetOrCreate(category *entity.Category) (*entity.Category, error) {
	return s.categoryRepository.GetOrCreate(category)
}
