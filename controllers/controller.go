package controllers

import (
	"net/http"
	"rest-api-practice/database"
	"rest-api-practice/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db database.Database
}

func parseIdParamsWithAbortion(ctx *gin.Context) int {
	id, parseIdError := strconv.Atoi(ctx.Param("id"))
	if parseIdError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, struct{ Message string }{"ID must be a number"})
		return -1
	}
	return id
}

func InitializeNewControllerWithDB(db database.Database) Controller {
	return Controller{db: db}
}

func (c Controller) GetUsers(ctx *gin.Context) {
	users := c.db.GetUsers()
	ctx.JSON(200, users)
}

func (c Controller) GetUserById(ctx *gin.Context) {
	id := parseIdParamsWithAbortion(ctx)
	if id == -1 {
		return
	}

	user, err := c.db.GetUserById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, user)
}

func (c Controller) CreateUser(ctx *gin.Context) {
	newUser := models.User{}
	ctx.ShouldBindJSON(&newUser)
	c.db.CreateUser(newUser)
}

func (c Controller) UpdateUserById(ctx *gin.Context) {
	id := parseIdParamsWithAbortion(ctx)
	if id == -1 {
		return
	}

	user, getUserError := c.db.GetUserById(id)
	userWithUpdatedData := user
	if getUserError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, getUserError)
		return
	}

	ctx.ShouldBindJSON(&userWithUpdatedData)
	c.db.UpdateUserWithNewUserData(user, userWithUpdatedData)
}

func (c Controller) DeleteUserById(ctx *gin.Context) {
	id := parseIdParamsWithAbortion(ctx)
	if id == -1 {
		return
	}

	c.db.DeleteUserById(id)
}
