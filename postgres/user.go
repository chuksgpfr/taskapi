package postgres

import (
	"errors"
	"fmt"

	taskapi "github.com/chuksgpfr/task-api"
	"github.com/chuksgpfr/task-api/pkg"
	"gorm.io/gorm"
)

func (d *DbService) Register(body *taskapi.RegisterParam) (*taskapi.User, error) {
	var user *taskapi.User
	err := d.DB.Take(&user, &taskapi.User{Email: body.Email}).Error

	fmt.Println("1", err)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("Failed to create new account")
	}

	if err == nil {
		return nil, errors.New("User with this email already exist")
	}

	hashedPassword, _ := pkg.HashPassword(body.Password)

	err = d.DB.Create(&taskapi.User{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  hashedPassword,
	}).Error

	fmt.Println("2 ", err)

	if err != nil {
		return nil, errors.New("Failed to create new account")
	}

	err = d.DB.Take(&user, &taskapi.User{Email: body.Email}).Error
	if err != nil {
		return nil, errors.New("Failed to retrieve created user")
	}

	if err != nil {
		return nil, errors.New("Failed to create new account")
	}

	return user, nil
}

func (d *DbService) Login(body *taskapi.LoginParam) (*taskapi.User, error) {
	var user *taskapi.User
	err := d.DB.Take(&user, &taskapi.User{Email: body.Email}).Error

	if err != nil {
		return nil, errors.New("Invalid email and password combination")
	}

	err = pkg.ComparePassword(user.Password, body.Password)

	if err != nil {
		return nil, errors.New("Invalid email and password combination")
	}

	return user, nil
}
