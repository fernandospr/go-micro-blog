package main

import (
	"go-micro-blog/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func getPosts(c *gin.Context) {
	c.JSON(http.StatusOK, models.Posts())
}

func getPostsByUser(c *gin.Context) {
	user := c.Param("user")

	postsFromUser := models.FilterPostsByUser(user)

	c.JSON(http.StatusOK, postsFromUser)
}

func addPost(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.AddPost(newPost)

	c.JSON(http.StatusCreated, newPost)
}

func main() {
	models.Init()

	r := gin.Default()

	r.GET("/posts", getPosts)
	r.GET("/posts/:user", getPostsByUser)
	r.POST("/posts", addPost)

	r.Run()
}
