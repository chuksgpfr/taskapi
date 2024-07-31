package main

// import (
// 	"testing"

// 	"github.com/gin-gonic/gin"
// )

// func SetUpRouter() *gin.Engine {
// 	router := gin.Default()
// 	return router
// }

// func TestHomepageHandler(t *testing.T) {
// 	mockResponse := `{"message":"Welcome"}`
// 	r := SetUpRouter()
// 	r.GET("/", HomepageHandler)
// 	req, _ := http.NewRequest("GET", "/", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	responseData, _ := io.ReadAll(w.Body)
// 	assert.Equal(t, mockResponse, string(responseData))
// 	assert.Equal(t, http.StatusOK, w.Code)
// }
