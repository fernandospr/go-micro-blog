package main

import (
	"go-micro-blog/models"
	"strconv"

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

func getPostsById(c *gin.Context) {
	id := c.Param("id")

	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	} else {
		post := models.FindPostById(uint(id64))
		if post == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		} else {
			c.JSON(http.StatusOK, post)
		}
	}
}

func addPost(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, e := models.AddPost(&newPost); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPost)
}

func main() {
	models.Init()

	r := gin.Default()

	r.GET("/posts", getPosts)
	r.GET("/users/:user/posts", getPostsByUser)
	r.GET("/posts/:id", getPostsById)
	r.POST("/posts", addPost)

	r.Run()
}
