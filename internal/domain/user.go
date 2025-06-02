package domain

type User struct {
	Id       int32  `validate:"gte=0"`
	Username string `validate:"required,min=3,max=32"`
	Name     string `validate:"required,min=2,max=64"`
	Surname  string `validate:"required,min=2,max=64"`
	Password string `validate:"required,min=8"`
	Email    string `validate:"required,email"`
}
