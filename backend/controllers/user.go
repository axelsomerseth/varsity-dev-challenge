package controllers

import (
	"net/http"

	"github.com/axelsomerseth/varsity-dev-challenge/backend/db"
	"github.com/axelsomerseth/varsity-dev-challenge/backend/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userInput struct {
		Username string `json:"username" validate:"required,max=20"`
		Email    string `json:"email" validate:"required,max=255"`
		Subject  string `json:"subject" validate:"required"`
		Picture  string `json:"picture" validate:"url"`
	}

	if err := c.BindJSON(&userInput); err != nil {
		return
	}

	user := models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Subject:  userInput.Subject,
		Picture:  userInput.Picture,
	}

	if r := db.Connection.Create(&user); r.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": r.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func GetUser(c *gin.Context) {
	userID := c.Param("userId")

	var user models.User
	if r := db.Connection.First(&user, userID); r.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": r.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}
