package pkg

import (
	"fmt"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/models"
)

const (
	domain = "https://api.kron-x.ru/data/photos/"
)

func ProductPhotoChecker(photo string) string {
	if photo == "" {
		return ""
	} else {
		return domain + photo
	}
}

func RelativesProduct(product []*models.Products) []*dto.Product {
	var productDTO []*dto.Product

	for x := range product {
		result := &dto.Product{
			Id:            fmt.Sprint(product[x].Id),
			Category:      fmt.Sprint(product[x].Category),
			Title:         product[x].Title,
			VendorCode:    product[x].VendorCode,
			Code1c:        product[x].Code1c,
			Description:   product[x].Description,
			Price:         fmt.Sprint(product[x].Price),
			ImageOriginal: ProductPhotoChecker(product[x].ImageOriginal),
			Image128:      ProductPhotoChecker(product[x].Image128),
			Image432:      ProductPhotoChecker(product[x].Image432),
		}
		productDTO = append(productDTO, result)
	}
	return productDTO
}
