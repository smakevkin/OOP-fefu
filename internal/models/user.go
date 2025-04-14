package models

import "time"

type User struct {
	UserName   string    `json:"user_name"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	RegisterAt time.Time `json:"register_at"`
}
