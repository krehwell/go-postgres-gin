package controllers

import (
	"rest-api-practice/database"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db database.Database
}

func InitializeNewControllerWithDB(db database.Database) Controller {
	return Controller{db: db}
}

func (c Controller) GetUsers(ctx *gin.Context) {
	users := c.db.GetUsers()
	ctx.JSON(200, users)
}

func (c Controller) CreateUser(ctx *gin.Context) {

}
