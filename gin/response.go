package gin

import (
	"github.com/gin-gonic/gin"
)

func ErrorResponse(msg string) *gin.H {
	return &gin.H{"message": msg}
}

func SuccessResponse(msg string, data interface{}) *gin.H {
	return &gin.H{"message": msg, "data": data}
}
