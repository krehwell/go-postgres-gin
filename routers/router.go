package routers

import (
	"rest-api-practice/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(address string, c controllers.Controller) *gin.Engine {
	r := gin.Default()

	r.GET("/person", c.GetPerson)
	r.GET("/person/:id")
	r.POST("/person", c.CreatePerson)
	r.PUT("/person/:id")
	r.DELETE("/person/:id")

	return r
}
