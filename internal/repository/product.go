package repository

import (
	"fmt"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAllProductsByParams(params *dto.Params) ([]dto.Product, error) {
	var products []dto.Product
	var err error
	fmt.Println(params.Limit)
	fmt.Println(params.Page)
	if params.Limit == 0 {
		err = r.db.Select(&products, "SELECT id, category_id, title, price, image  FROM product")

	} else {
		err = r.db.Select(&products, "SELECT id, category_id, title, price, image  FROM product LIMIT $1 OFFSET $2", params.Limit, params.Page-1)
	}
	return products, err
}

func (r *ProductRepository) GetProductDetail(id string) (*dto.ProductInformation, error) {
	return nil, nil
}

func (r *ProductRepository) ParseUrlParams() []int {
	//var allCategoriesIds []models.Category
	//
	//var categoriesIds []int
	//r.db.Model(&allCategoriesIds).Pluck("Id", &categoriesIds)
	//return categoriesIds
	return nil
}
