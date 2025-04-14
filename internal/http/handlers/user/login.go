package user

import "oop/internal/lib/api/response"

type Request struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}

type Response struct {
	response.Response
	UserName   string `json:"user_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	RegisterAt string `json:"register_at"`
}
