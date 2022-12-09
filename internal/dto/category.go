package dto

type Category struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type CategoryParams struct {
	ID string
}
