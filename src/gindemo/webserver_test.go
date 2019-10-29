package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)



func getRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestPermissionChange(t *testing.T) {


	// create a response recorder
	response := httptest.NewRecorder()

	router := getRouter()

	permChangePath := "/api/auth/v1/permission/resource/:id"

	//authPermission := router.Group("/api/auth/v1/permission")

	router.POST(permChangePath, PermissionChange)

	// get param
	req, _ := http.NewRequest(http.MethodPost, permChangePath, nil)
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Fail()
	}

}
