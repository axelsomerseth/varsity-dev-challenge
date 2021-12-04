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
	postID := c.Param("postId")

	var post models.Post
	if r := db.Connection.Preload("User").First(&post, postID); r.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": r.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, &post)
}

func ListPost(c *gin.Context) {
	postID := c.Param("postId")

	var posts []models.Post
	if r := db.Connection.Preload("User").Select(&posts, postID); r.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": r.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &posts,
	})
}
