package factories

import (
	"github.com/gin-gonic/gin"

	"go-micro-blog/controllers"
)

func NewEngine(controller controllers.PostsController) *gin.Engine {
	engine := gin.Default()

	engine.GET("/posts", controller.GetPosts)
	engine.GET("/users/:user/posts", controller.GetPostsByUser)
	engine.GET("/posts/:id", controller.GetPostById)
	engine.POST("/posts", controller.AddPost)

	return engine
}
