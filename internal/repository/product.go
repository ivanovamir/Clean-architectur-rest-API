package repository

import (
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

func (r *ProductRepository) GetAllProductsByParams(params *dto.Params) ([]*dto.Product, error) {
	//var allProducts []*models.Products
	//offset := (params.Page - 1) * params.Limit
	//if err := r.db.Where("category IN (?)", params.CatId).Where("can_to_view",
	//	true).Select("Id", "Category", "Title", "Vendor_code",
	//	"Code1c", "Description", "Price", "Image_original", "Image128",
	//	"Image432").Limit(params.Limit).Offset(offset).Find(&allProducts).Error; err != nil {
	//	return nil, err
	//} else {
	//	var resAlready []*dto.Product
	//	for a := range allProducts {
	//		//TODO implement products photos methods in other module
	//		result := &dto.Product{
	//			Id:            fmt.Sprint(allProducts[a].Id),
	//			Category:      fmt.Sprint(allProducts[a].Category),
	//			Title:         allProducts[a].Title,
	//			VendorCode:    allProducts[a].VendorCode,
	//			Code1c:        allProducts[a].Code1c,
	//			Description:   allProducts[a].Description,
	//			Price:         fmt.Sprint(allProducts[a].Price),
	//			ImageOriginal: internal.ProductPhotoChecker(allProducts[a].ImageOriginal),
	//			Image128:      internal.ProductPhotoChecker(allProducts[a].Image128),
	//			Image432:      internal.ProductPhotoChecker(allProducts[a].Image432),
	//		}
	//		resAlready = append(resAlready, result)
	//	}
	//	return resAlready, nil
	//}
	return nil, nil

}

func (r *ProductRepository) GetProductDetail(id string) (*dto.ProductInformation, error) {

	//rand.Seed(time.Now().Unix())
	//
	//var productInfo *models.Products
	//
	//if err := r.db.Where("id=?", id).Select("Id", "Category",
	//	"Title", "Vendor_code", "Code1c", "Description", "Price", "Image_original",
	//	"Image128", "Image432").First(&productInfo).Error; err != nil {
	//	return nil, err
	//} else {
	//	var idsTrue []int
	//
	//	var relativesProducts []*models.Products
	//
	//	r.db.Model(&[]models.Products{}).Where("category = ?",
	//		&productInfo.Category).Not("id = (?)",
	//		&productInfo.Id).Pluck("id", &idsTrue)
	//
	//	var idsArray []int
	//
	//	if idsTrue == nil {
	//		idsArray = []int{}
	//	} else {
	//		for i := 0; i <= 4; i++ {
	//			randomIndex := rand.Intn(len(idsTrue))
	//			pick := idsTrue[randomIndex]
	//			idsArray = append(idsArray, pick)
	//		}
	//	}
	//
	//	r.db.Select("Id", "Category", "Title", "Vendor_code",
	//		"Code1c", "Description", "Price", "Image_original",
	//		"Image128", "Image432").Find(&relativesProducts, idsArray)
	//
	//	productInfoDto := &dto.ProductInformation{
	//		Id:                fmt.Sprint(productInfo.Id),
	//		Category:          fmt.Sprint(productInfo.Category),
	//		Title:             productInfo.Title,
	//		VendorCode:        productInfo.VendorCode,
	//		Code1c:            productInfo.Code1c,
	//		Quantity:          "1",
	//		Description:       productInfo.Description,
	//		Price:             fmt.Sprint(productInfo.Price),
	//		ImageOriginal:     internal.ProductPhotoChecker(productInfo.ImageOriginal),
	//		Image128:          internal.ProductPhotoChecker(productInfo.Image128),
	//		Image432:          internal.ProductPhotoChecker(productInfo.Image432),
	//		RelativesProducts: internal.RelativesProduct(relativesProducts),
	//	}
	//
	//	return productInfoDto, nil
	//}
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
