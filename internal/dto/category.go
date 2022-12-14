package dto

type Category struct {
	ID    string `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Image string `json:"image" db:"image"`
}

type CategoryParams struct {
	ID string
}
