package postgres

import (
	"errors"
	"testing"

	taskapi "github.com/chuksgpfr/task-api"
	mock_taskapi "github.com/chuksgpfr/task-api/mocks"
	"github.com/chuksgpfr/task-api/pkg"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDbService_Register(t *testing.T) {
	// Initialize the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock instance of the DbServiceInterface
	mockService := mock_taskapi.NewMockUserService(ctrl)

	// ctx := context.Background()

	t.Run("Successful registration", func(t *testing.T) {

		// Define the input and expected output
		email := "newuser@example.com"
		password := "testpassword"
		hashedPassword, _ := pkg.HashPassword(password)
		body := &taskapi.RegisterParam{
			Email:     email,
			Password:  password,
			FirstName: "Jane",
			LastName:  "Doe",
		}
		expectedUser := &taskapi.User{
			Email:     email,
			FirstName: "Jane",
			LastName:  "Doe",
			Password:  hashedPassword,
		}

		// Set up the expectation
		mockService.EXPECT().Register(body).Return(expectedUser, nil)

		// Call the Register function
		user, err := mockService.Register(body)

		// Assertions to verify the results
		assert.NotNil(t, user)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("Failed to create new account", func(t *testing.T) {
		// Define the input and expected output
		email := "newuser@example.com"
		password := "testpassword"
		// hashedPassword, _ := pkg.HashPassword(password)
		body := &taskapi.RegisterParam{
			Email:     email,
			Password:  password,
			FirstName: "Jane",
			LastName:  "Doe",
		}

		// Set up the expectation
		expectedError := errors.New("Failed to create new account")
		mockService.EXPECT().Register(body).Return(nil, expectedError)

		// Call the Register function
		user, err := mockService.Register(body)

		// Assertions to verify the results
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestDbService_Login(t *testing.T) {
	// Initialize the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock instance of the DbServiceInterface
	mockService := mock_taskapi.NewMockUserService(ctrl)

	t.Run("failed validator", func(t *testing.T) {
		// Define the input and expected output
		email := "newuser@example.com"
		// password := "testpassword"
		// hashedPassword, _ := pkg.HashPassword(password)
		body := &taskapi.LoginParam{
			Email: email,
		}

		// Set up the expectation
		expectedError := errors.New("password is required")
		mockService.EXPECT().Login(body).Return(nil, expectedError)

		user, err := mockService.Login(body)
		// Assertions to verify the results
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
