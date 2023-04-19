package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddPost(t *testing.T) {
	repo := InMemoryPostsRepository{}

	post := Post{
		ID:   1234,
		User: "TestUser",
		Text: "Hello world!",
	}
	repo.AddPost(&post)

	assert.Contains(t, *repo.Posts(), post)
}

func Test_AddPost_ShouldReturnError_WhenAddingPostWithExistingId(t *testing.T) {
	repo := InMemoryPostsRepository{}

	post := Post{
		ID:   1234,
		User: "TestUser",
		Text: "Hello world!",
	}
	repo.AddPost(&post)
	p, err := repo.AddPost(&post)

	assert.Nil(t, p)
	assert.NotNil(t, err)
}

func Test_FindPostById(t *testing.T) {
	repo := InMemoryPostsRepository{}
	post := Post{
		ID:   1234,
		User: "TestUser",
		Text: "Hello world!",
	}
	repo.AddPost(&post)

	postFound := repo.FindPostById(1234)

	assert.Equal(t, post, *postFound)
}

func Test_FindPostById_ShouldReturnNil_WhenIdNotFound(t *testing.T) {
	repo := InMemoryPostsRepository{}

	postFound := repo.FindPostById(1234)

	assert.Nil(t, postFound)
}

func Test_FilterPostsByUser(t *testing.T) {
	repo := InMemoryPostsRepository{}
	testUserPosts := []Post{
		{
			ID:   1234,
			User: "TestUser",
			Text: "Hello world 1!",
		},
		{
			ID:   5678,
			User: "TestUser",
			Text: "Hello world 2!",
		},
	}
	otherUserPosts := []Post{
		{
			ID:   9012,
			User: "OtherUser",
			Text: "Hello world!",
		},
	}
	allPosts := append(testUserPosts, otherUserPosts...)
	for _, p := range allPosts {
		repo.AddPost(&p)
	}

	postsFound := repo.FilterPostsByUser("TestUser")

	assert.Equal(t, testUserPosts, *postsFound)
}

func Test_FilterPostsByUser_ShouldReturnEmptySlice_WhenNoPostsFromUser(t *testing.T) {
	repo := InMemoryPostsRepository{}

	postsFound := repo.FilterPostsByUser("TestUser")

	assert.Empty(t, *postsFound)
}
