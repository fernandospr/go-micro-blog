package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqlPostsRepository struct {
	db        *gorm.DB
	Dialector func() gorm.Dialector
}

func (repository *SqlPostsRepository) Init() {
	database, err := gorm.Open(repository.Dialector(), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Post{})
	if err != nil {
		return
	}

	repository.db = database
}

func (repository *SqlPostsRepository) Posts() *[]Post {
	var posts []Post
	repository.db.Find(&posts)
	return &posts
}

func (repository *SqlPostsRepository) AddPost(newPost *Post) (*Post, error) {
	if err := repository.db.Create(newPost).Error; err != nil {
		return nil, fmt.Errorf("models: %v", err)
	}

	return newPost, nil
}

func (repository *SqlPostsRepository) FilterPostsByUser(user string) *[]Post {
	postsFromUser := []Post{}
	repository.db.Where("user = ?", user).Find(&postsFromUser)
	return &postsFromUser
}

func (repository *SqlPostsRepository) FindPostById(id uint) *Post {
	var post Post
	if err := repository.db.Where("id = ?", id).First(&post).Error; err != nil {
		return nil
	}
	return &post
}

func SqliteDialector() gorm.Dialector {
	dsn := "test.db"
	return sqlite.Open(dsn)
}

func MySqlDialector() gorm.Dialector {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname"
	return mysql.Open(dsn)
}
