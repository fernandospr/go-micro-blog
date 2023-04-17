package models

type PostsRepository interface {
	Posts() *[]Post
	AddPost(newPost *Post) (*Post, error)
	FilterPostsByUser(user string) *[]Post
	FindPostById(id uint) *Post
}
