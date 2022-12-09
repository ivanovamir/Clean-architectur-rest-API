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
	if err != nil {
		return nil, err
	}
	return allCategories, nil
}

func (r *CategoryRepository) GetCategoriesById(id *dto.CategoryParams) ([]*dto.Category, error) {
	//var category *models.Category
	//if err := r.db.Where("id=?", id.ID).First(&category).Error; err != nil {
	//	return nil, err
	//} else {
	//	var resAlready []*dto.Category
	//	result := &dto.Category{
	//		ID:         fmt.Sprint(category.Id),
	//		Title:      category.Title,
	//		Image:      category.Image,
	//		CountItems: "0",
	//	}
	//	resAlready = append(resAlready, result)
	//	return resAlready, nil
	//}
	return nil, nil
}
