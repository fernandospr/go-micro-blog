package models

type Post struct {
	User string `json:"user"`
	Text string `json:"text"`
}

var posts []Post

func Posts() *[]Post {
	return &posts
}

func Init() {
	posts = []Post{
		{User: "user1", Text: "Hello World"},
		{User: "user2", Text: "Hola Mundo"},
		{User: "user1", Text: "Other message from user"},
		{User: "user2", Text: "Hello World"},
	}
}

func AddPost(newPost Post) {
	posts = append(posts, newPost)
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
