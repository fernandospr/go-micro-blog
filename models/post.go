package models

import "fmt"

type Post struct {
	ID   uint   `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
}

var posts []Post

func Posts() *[]Post {
	return &posts
}

func Init() {
	posts = []Post{
		{ID: 1, User: "user1", Text: "Hello World"},
		{ID: 2, User: "user2", Text: "Hola Mundo"},
		{ID: 3, User: "user1", Text: "Other message from user"},
		{ID: 4, User: "user2", Text: "Hello World"},
	}
}

func AddPost(newPost *Post) (*Post, error) {
	if postExists(newPost) {
		return nil, fmt.Errorf("models: post with id %d already exists", newPost.ID)
	}
	posts = append(posts, *newPost)
	return newPost, nil
}

func postExists(post *Post) bool {
	for _, p := range posts {
		if p.ID == post.ID {
			return true
		}
	}
	return false
}

func FilterPostsByUser(user string) *[]Post {
	postsFromUser := []Post{}
	for _, p := range posts {
		if p.User == user {
			postsFromUser = append(postsFromUser, p)
		}
	}
	return &postsFromUser
}
