package gin

import (
	"github.com/chuksgpfr/task-api/pkg"
	"github.com/chuksgpfr/task-api/postgres"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewServer(config pkg.Configuration, DB *gorm.DB) *gin.Engine {
	gin.SetMode(config.GIN_MODE)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORS())

	service := postgres.DbService{
		DB: DB,
	}

	route := Route{
		Service: service,
		Router:  router,
		DB:      DB,
		Config:  config,
	}

	Routes(&route)
	return router
}
