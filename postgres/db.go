package postgres

import (
	"log"
	"os"
	"time"

	taskapi "github.com/chuksgpfr/task-api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
)

func dbClient(postgresDSN string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{
		TranslateError:                           true,
		QueryFields:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
	return db
}

func NewDbClient(postgresDSN string) *gorm.DB {
	client := dbClient(postgresDSN)

	client.AutoMigrate(&taskapi.Task{})
	client.AutoMigrate(&taskapi.User{})
	return client
}

type DbService struct {
	DB *gorm.DB
}
