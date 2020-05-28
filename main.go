package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-admin/middleware"
	"resource-admin/utils/log"
)

func test(c *gin.Context) {
	log.Info("hello, world")
	c.String(http.StatusOK, "hello, world")
}

func main() {

	router := gin.Default()

	router.Use(middleware.Cors())
	router.Use(middleware.ValidateUser())

	router.GET("/test", test)

	router.Run("127.0.0.1:7776")

}
