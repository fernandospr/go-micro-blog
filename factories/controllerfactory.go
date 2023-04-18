package factories

import (
	"go-micro-blog/controllers"
	"go-micro-blog/models"
)

func NewPostsController(repository models.PostsRepository) controllers.PostsController {
	controller := controllers.JsonPostsController{Repository: repository}
	return &controller
}
