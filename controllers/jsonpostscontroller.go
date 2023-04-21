package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-micro-blog/models"
)

type JsonPostsController struct {
	Repository models.PostsRepository
}

func (pc *JsonPostsController) GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, pc.Repository.Posts())
}

func (pc *JsonPostsController) GetPostsByUser(c *gin.Context) {
	user := c.Param("user")

	postsFromUser := pc.Repository.FilterPostsByUser(user)

	c.JSON(http.StatusOK, postsFromUser)
}

func (pc *JsonPostsController) GetPostById(c *gin.Context) {
	id := c.Param("id")

	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	} else {
		post := pc.Repository.FindPostById(uint(id64))
		if post == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		} else {
			c.JSON(http.StatusOK, post)
		}
	}
}

func (pc *JsonPostsController) AddPost(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, e := pc.Repository.AddPost(&newPost); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPost)
}
