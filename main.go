package main

import (
	"go-micro-blog/config"
	"go-micro-blog/controllers"
	"go-micro-blog/factories"

	"github.com/gin-gonic/gin"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	var cfg config.Config
	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		panic("Missing config.yml")
	}

	r := gin.Default()

	repository := factories.NewRepository(cfg.Database)
	c := controllers.PostsController{Repository: repository}

	r.GET("/posts", c.GetPosts)
	r.GET("/users/:user/posts", c.GetPostsByUser)
	r.GET("/posts/:id", c.GetPostsById)
	r.POST("/posts", c.AddPost)

	r.Run(":" + cfg.Server.Port)
}
