package controllers

import (
	"fmt"
	"rest-api-practice/database"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db database.Database
}

func InitializeNewControllerWithDB(db database.Database) Controller {
	return Controller{db: db}
}

func (c Controller) GetPerson(ctx *gin.Context) {
	users := c.db.GetUsers()
	fmt.Println(users)
	ctx.JSON(200, users)
}

func (c Controller) CreatePerson(ctx *gin.Context) {

}
