package service

import (
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/repository"
)

type Service struct {
	Category
	Product
	User
}

type Category interface {
	GetAllCategories() ([]dto.Category, error)
	GetCategoriesById(id *dto.CategoryParams) ([]*dto.Category, error)
}

type Product interface {
	GetAllProductsByParams(params *dto.Params) ([]*dto.Product, error)
	GetProductDetail(id string) (*dto.ProductInformation, error)
	ParseUrlParams() []int
}

type User interface {
	ParseToken(accessToken string) (int, error)
	RegisterUser(userDTO *dto.RegisterUser) (string, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Category: NewCategoryService(repo),
		Product:  NewProductService(repo),
		User:     NewUserService(repo),
	}
}
