package models

type Category struct {
	Id       int64
	Title    string
	Image    string
	Products []Products
}
