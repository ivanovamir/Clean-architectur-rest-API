package service

import (
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) GetAllCategories() ([]dto.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *CategoryService) GetCategoriesById(id *dto.CategoryParams) ([]*dto.Category, error) {
	return s.repo.GetCategoriesById(id)
}
