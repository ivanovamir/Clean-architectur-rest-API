package repository

import (
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
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
	RegisterUser(userDTO *dto.RegisterUser) (int, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Category: NewCategoryRepository(db),
		Product:  NewProductRepository(db),
		User:     NewUserRepository(db),
	}
}
