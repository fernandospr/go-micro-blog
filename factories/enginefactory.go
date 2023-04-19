package factories

import (
	"go-micro-blog/controllers"

	"github.com/gin-gonic/gin"
)

func NewEngine(controller controllers.PostsController) *gin.Engine {
	engine := gin.Default()

	engine.GET("/posts", controller.GetPosts)
	engine.GET("/users/:user/posts", controller.GetPostsByUser)
	engine.GET("/posts/:id", controller.GetPostsById)
	engine.POST("/posts", controller.AddPost)

	return engine
}
