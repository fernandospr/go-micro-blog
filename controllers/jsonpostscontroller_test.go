package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-micro-blog/models"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Posts() *[]models.Post {
	args := r.Called()
	return args.Get(0).(*[]models.Post)
}

func (r *repositoryMock) AddPost(newPost *models.Post) (*models.Post, error) {
	args := r.Called(newPost)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (r *repositoryMock) FilterPostsByUser(user string) *[]models.Post {
	args := r.Called(user)
	return args.Get(0).(*[]models.Post)
}

func (r *repositoryMock) FindPostById(id uint) *models.Post {
	args := r.Called(id)
	return args.Get(0).(*models.Post)
}

func Test_GetPosts_ShouldReturnPostsFromRepository(t *testing.T) {
	posts := &[]models.Post{
		{
			ID:   1234,
			User: "TestUser",
			Text: "Hello world!",
		},
	}
	repositoryMockObj := new(repositoryMock)
	repositoryMockObj.On("Posts").Return(posts)
	controller := JsonPostsController{Repository: repositoryMockObj}
	engine := gin.Default()
	engine.GET("/posts", controller.GetPosts)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[{\"id\":1234,\"user\":\"TestUser\",\"text\":\"Hello world!\"}]", w.Body.String())
}

func Test_GetPostById_ShouldReturnPostFromRepository(t *testing.T) {
	post := &models.Post{
		ID:   1234,
		User: "TestUser",
		Text: "Hello world!",
	}
	repositoryMockObj := new(repositoryMock)
	repositoryMockObj.On("FindPostById", uint(1234)).Return(post)
	controller := JsonPostsController{Repository: repositoryMockObj}
	engine := gin.Default()
	engine.GET("/posts/:id", controller.GetPostById)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts/1234", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"id\":1234,\"user\":\"TestUser\",\"text\":\"Hello world!\"}", w.Body.String())
}

func Test_GetPostById_ShouldReturnNotFound_WhenRepositoryDoesNotFindPost(t *testing.T) {
	var noPost *models.Post
	repositoryMockObj := new(repositoryMock)
	repositoryMockObj.On("FindPostById", mock.Anything).Return(noPost)
	controller := JsonPostsController{Repository: repositoryMockObj}
	engine := gin.Default()
	engine.GET("/posts/:id", controller.GetPostById)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts/5678", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}
