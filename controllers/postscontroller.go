package controllers

import (
	"github.com/gin-gonic/gin"
)

type PostsController interface {
	GetPosts(c *gin.Context)
	GetPostsByUser(c *gin.Context)
	GetPostById(c *gin.Context)
	AddPost(c *gin.Context)
}
