package controllers

import (
	"net/http"

	"github.com/axelsomerseth/varsity-dev-challenge/backend/db"
	"github.com/axelsomerseth/varsity-dev-challenge/backend/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var postInput struct {
		UserID uint   `validate:"required"`
		Post   string `validate:"required,max=140"`
	}

	if err := c.BindJSON(&postInput); err != nil {
		return
	}

	post := models.Post{
		UserID: postInput.UserID,
		Post:   postInput.Post,
	}

	if r := db.Connection.Create(&post); r.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": r.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, &post)
}

func GetPost(c *gin.Context) {

}
