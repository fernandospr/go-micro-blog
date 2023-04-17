package main

import (
	"go-micro-blog/controllers"
	"go-micro-blog/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.Init()

	r := gin.Default()

	r.GET("/posts", controllers.GetPosts)
	r.GET("/users/:user/posts", controllers.GetPostsByUser)
	r.GET("/posts/:id", controllers.GetPostsById)
	r.POST("/posts", controllers.AddPost)

	r.Run()
}
