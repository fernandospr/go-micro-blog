package main

import (
	"go-micro-blog/controllers"
	"go-micro-blog/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repository := models.SqlPostsRepository{}
	repository.Init()
	c := controllers.PostsController{Repository: &repository}

	r.GET("/posts", c.GetPosts)
	r.GET("/users/:user/posts", c.GetPostsByUser)
	r.GET("/posts/:id", c.GetPostsById)
	r.POST("/posts", c.AddPost)

	r.Run()
}
