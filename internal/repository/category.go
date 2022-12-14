package repository

import (
	"fmt"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetAllCategories() ([]dto.Category, error) {
	var allCategories []dto.Category
	query := fmt.Sprint("SELECT id, title, image FROM category")
	err := r.db.Select(&allCategories, query)
	return allCategories, err
}

func (r *CategoryRepository) GetCategoriesById(id *dto.CategoryParams) ([]dto.Category, error) {
	var category []dto.Category
	err := r.db.Select(&category, "SELECT id, title, image FROM category WHERE id = $1", id.ID)
	return category, err
}
