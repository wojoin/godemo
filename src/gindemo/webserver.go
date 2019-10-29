package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func PermissionChange(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

func main() {
	log.Println("OK")

	router := gin.Default()

	authPermission := router.Group("/api/test/auth/v1/permission")
	{
		// resource
		//authPermission.GET("/resources", ResourceList)
		//authPermission.GET("/resource/:id", ResourceInfo)
		//authPermission.DELETE("/resource/:id", ResourceDelete)
		//authPermission.POST("/resource", ResourceCreate)
		//authPermission.POST("/verification", PermVerify)
		authPermission.GET("/resource", PermissionChange)
	}

	var systemroles []string
	log.Println(systemroles)

	if systemroles == nil {
		log.Println("-----------")
	}

	router.Run(":9090")
}
