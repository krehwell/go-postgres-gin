package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"rest-api-practice/database"
	"rest-api-practice/models"
	"strconv"

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

func (c Controller) GetUserById(ctx *gin.Context) {
	_id := ctx.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("id should be a number"))
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
	_id := ctx.Param("id")
	id, parseIdError := strconv.Atoi(_id)
	if parseIdError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("id should be a number"))
		return
	}

	user, getUserError := c.db.GetUserById(id)
	userWithUpdatedData := user
	if getUserError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, getUserError)
		return
	}

	ctx.ShouldBindJSON(&userWithUpdatedData)
	if err := c.db.UpdateUserWithNewUserData(user, userWithUpdatedData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("one user updated")
	fmt.Println("prev data", user)
	fmt.Println("updated data", userWithUpdatedData)
}
