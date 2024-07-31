package gin

import (
	"github.com/chuksgpfr/task-api/pkg"
	"github.com/chuksgpfr/task-api/postgres"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Route struct {
	Service postgres.DbService
	Router  *gin.Engine
	DB      *gorm.DB
	Config  pkg.Configuration
}

func Routes(a *Route) {
	taskHandler := TaskHandler{
		TaskService: &a.Service,
		Config:      a.Config,
	}

	userHandler := UserHandler{
		UserService: &a.Service,
		Config:      a.Config,
	}

	v1 := a.Router.Group("v1")
	{
		authGroup := v1.Group("auth")
		{
			authGroup.POST("/register", userHandler.Register)
			authGroup.POST("/login", userHandler.Login)
		}

		taskGroup := v1.Group("task", Auth(a.DB, a.Config))
		{
			taskGroup.POST("/", taskHandler.CreateTask)
			taskGroup.GET("/", taskHandler.GetTasks)

			taskGroup.GET("/:slug", taskHandler.GetTask)
			taskGroup.PATCH("/:slug", taskHandler.UpdateTask)
			taskGroup.POST("/:slug", taskHandler.CompleteTask)
		}

	}

}
