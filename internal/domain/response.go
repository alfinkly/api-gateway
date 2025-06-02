package domain

type RegisterResponse struct {
	Username string `json:"username" validate:"required,min=3,max=64"`
	Name     string `json:"name" validate:"required,min=3,max=64"`
	Surname  string `json:"surname" validate:"required,min=3,max=64"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
