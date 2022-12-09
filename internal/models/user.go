package models

import "time"

type User struct {
	Id           uint
	CreatedAt    time.Time
	Inn          string
	Email        string
	PasswordHash string
	Phone        string
}
