package models

import "fmt"

type InMemoryPostsRepository struct {
	posts []Post
}

func (repository *InMemoryPostsRepository) Posts() *[]Post {
	return &repository.posts
}

func (repository *InMemoryPostsRepository) Init() {
	repository.posts = []Post{
		{ID: 1, User: "user1", Text: "Hello World"},
		{ID: 2, User: "user2", Text: "Hola Mundo"},
		{ID: 3, User: "user1", Text: "Other message from user"},
		{ID: 4, User: "user2", Text: "Hello World"},
	}
}

func (repository *InMemoryPostsRepository) AddPost(newPost *Post) (*Post, error) {
	if repository.postExists(newPost) {
		return nil, fmt.Errorf("models: post with id %d already exists", newPost.ID)
	}
	repository.posts = append(repository.posts, *newPost)
	return newPost, nil
}

func (repository *InMemoryPostsRepository) postExists(post *Post) bool {
	for _, p := range repository.posts {
		if p.ID == post.ID {
			return true
		}
	}
	return false
}

func (repository *InMemoryPostsRepository) FilterPostsByUser(user string) *[]Post {
	postsFromUser := []Post{}
	for _, p := range repository.posts {
		if p.User == user {
			postsFromUser = append(postsFromUser, p)
		}
	}
	return &postsFromUser
}

func (repository *InMemoryPostsRepository) FindPostById(id uint) *Post {
	for _, p := range repository.posts {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
