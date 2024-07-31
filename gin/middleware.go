package gin

import (
	"net/http"
	"strings"

	taskapi "github.com/chuksgpfr/task-api"
	"github.com/chuksgpfr/task-api/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Method", "POST, PUT, PATCH, DELETE, GET, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func Auth(db *gorm.DB, config pkg.Configuration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse("no authorization header passed for user"))
			return
		}

		tokenStr := strings.Split(authHeader, " ")[1]

		userId, _ := pkg.ValidateToken(tokenStr, config.LoginSymmetricKey)

		if userId == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse("unauthorized user"))
			return
		}

		var user *taskapi.User
		err := db.Take(&user, &taskapi.User{ID: userId}).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse("unauthorized user, login to continue"))
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
