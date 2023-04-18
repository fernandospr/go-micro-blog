package main

import (
	"go-micro-blog/factories"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := factories.NewConfig("./config/config.yml")
	repository := factories.NewPostsRepository(cfg.Database)
	controller := factories.NewPostsController(repository)

	r := gin.Default()

	r.GET("/posts", controller.GetPosts)
	r.GET("/users/:user/posts", controller.GetPostsByUser)
	r.GET("/posts/:id", controller.GetPostsById)
	r.POST("/posts", controller.AddPost)

	r.Run(":" + cfg.Server.Port)
}
