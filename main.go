package main

import (
	"go-micro-blog/factories"
)

func main() {
	cfg := factories.NewConfig("./config/config.yml")
	repository := factories.NewPostsRepository(cfg.Database)
	controller := factories.NewPostsController(repository)
	engine := factories.NewEngine(controller)
	engine.Run(":" + cfg.Server.Port)
}
