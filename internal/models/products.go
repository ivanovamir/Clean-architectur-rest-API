package models

const (
	domain = "https://api.kron-x.ru/data/photos/"
)

type Products struct {
	Id            uint
	Category      int64
	Title         string
	VendorCode    string
	Code1c        string
	Description   string
	Price         float32
	ImageOriginal string
	Image128      string
	Image432      string
}
