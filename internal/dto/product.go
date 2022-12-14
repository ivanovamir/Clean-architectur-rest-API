package dto

type Product struct {
	Id         string `json:"id" db:"id"`
	CategoryId int    `json:"category_id" db:"category_id"`
	Title      string `json:"title" db:"title"`
	Price      string `json:"price" db:"price"`
	Image      string `json:"image" db:"image"`
}

type ProductInformation struct {
	Id          string `json:"id" db:"id"`
	CategoryId  int    `json:"category_id" db:"category_id"`
	Title       string `json:"title" db:"title"`
	VendorCode  string `json:"vendor_code" db:"vendor_code"`
	Description string `json:"description" db:"description"`
	Price       string `json:"price" db:"price"`
	Image       string `json:"image" db:"image"`
}
