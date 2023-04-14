package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	User string
	Text string
}

var posts []Post

func getPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

func addPost(c *gin.Context) {
	var newPost Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	posts = append(posts, newPost)

	c.JSON(http.StatusCreated, newPost)
}

func main() {
	r := gin.Default()

	posts = []Post{
		{User: "user1", Text: "Hello World"},
		{User: "user2", Text: "Hola Mundo"},
		{User: "user1", Text: "Other message from user"},
		{User: "user2", Text: "Hello World"},
	}

	r.GET("/posts", getPosts)
	r.POST("/posts", addPost)

	r.Run()
}
