package routers

import (
	"rest-api-practice/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(address string, c controllers.Controller) *gin.Engine {
	r := gin.Default()

	r.GET("/users", c.GetUsers)
	r.GET("/user/:id")
	r.POST("/user", c.CreateUser)
	r.PUT("/user/:id")
	r.DELETE("/user/:id")

	return r
}
