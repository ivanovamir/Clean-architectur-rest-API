package dto

type RegisterUser struct {
	Name     string `json:"name"`
	Inn      string `json:"inn"`
	Email    string `json:"mail"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Code     string `json:"code" binding:"required"`
}

type UserAuth struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
