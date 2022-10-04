package controllers

import (
	"rest-api-practice/database"
	"rest-api-practice/models"

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
	newUser := models.User{}
	ctx.ShouldBindJSON(&newUser)
	c.db.CreateUser(newUser)
}
