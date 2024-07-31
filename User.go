package taskapi

import "gorm.io/gorm"

type RegisterParam struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
}

type LoginParam struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type User struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primarykey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserService interface {
	Register(body *RegisterParam) (*User, error)
	Login(body *LoginParam) (*User, error)
}
