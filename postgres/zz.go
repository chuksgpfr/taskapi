package postgres

// import (
// 	"errors"
// 	"fmt"
// 	"regexp"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	taskapi "github.com/chuksgpfr/task-api"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// // Mocking the HashPassword function
// var mockHashPassword = func(password string) (string, error) {
// 	return "hashedpassword", nil
// }

// func TestRegisterz(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	require.NoError(t, err)
// 	defer db.Close()

// 	gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: db,
// 	}), &gorm.Config{})
// 	require.NoError(t, err)

// 	service := &DbService{DB: gormDB}

// 	t.Run("User already exists", func(t *testing.T) {
// 		email := "test@example.com"
// body := &taskapi.RegisterParam{
// 	Email:     email,
// 	Password:  "password",
// 	FirstName: "John",
// 	LastName:  "Doe",
// }

// 		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."email" = $1 AND "users"."deleted_at" IS NULL LIMIT $2`)).
// 			WithArgs(email, 1).
// 			WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).AddRow(1, email))

// 		user, err := service.Register(body)

// 		assert.Nil(t, user)
// 		assert.EqualError(t, err, "User with this email already exist")
// 		assert.NoError(t, mock.ExpectationsWereMet())
// 	})

// 	t.Run("Failed to create new account - Take error", func(t *testing.T) {
// 		email := "newuser@example.com"
// 		password := "testpassword"

// 		// hashedPassword, _ := pkg.HashPassword(password)
// 		body := &taskapi.RegisterParam{
// 			Email:     email,
// 			Password:  password,
// 			FirstName: "Jane",
// 			LastName:  "Doe",
// 		}

// 		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."email" = $1 AND "users"."deleted_at" IS NULL LIMIT $2`)).
// 			WithArgs(email, 1).
// 			WillReturnError(gorm.ErrRecordNotFound)

// 		mock.ExpectBegin()
// 		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("created_at","updated_at","deleted_at","first_name","last_name","email","password") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id","id"`)).
// 			WithArgs(
// 				sqlmock.AnyArg(), // for created_at
// 				sqlmock.AnyArg(), // for updated_at
// 				nil,              // for deleted_at
// 				"Jane",
// 				"Doe",
// 				email,
// 				sqlmock.AnyArg(),
// 			).
// 			WillReturnError(errors.New("insert error"))

// 		mock.ExpectRollback()

// 		user, err := service.Register(body)

// 		assert.Nil(t, user)
// 		assert.EqualError(t, err, "Failed to create new account")
// 		assert.NoError(t, mock.ExpectationsWereMet())
// 	})

// 	t.Run("Successful registration", func(t *testing.T) {
// 		email := "newuser@example.com"
// 		password := "testpassword"

// 		// hashedPassword, _ := pkg.HashPassword(password)
// 		body := &taskapi.RegisterParam{
// 			Email:     email,
// 			Password:  password,
// 			FirstName: "Jane",
// 			LastName:  "Doe",
// 		}

// 		mock.ExpectBegin()
// 		mock.ExpectQuery(`^INSERT INTO users ("first_name","last_name","email","password") VALUES($1,$2,$3,$4)`).WithArgs(
// 			"Jane",
// 			"Doe",
// 			email,
// 			sqlmock.AnyArg(),
// 		).WillReturnError(nil)
// 		mock.ExpectCommit()

// 		if err := gormDB.Create(body).Error; err != nil {
// 			t.Fatalf("Failed to insert user: %v", err)
// 		}

// 		user, err := service.Register(body)
// 		fmt.Println("WWW ", user, err)

// 		assert.NotNil(t, user)
// 		assert.NoError(t, err)
// 		assert.NoError(t, mock.ExpectationsWereMet())
// 	})

// }
