package storePosts

import (
	"fintech-app/graph/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тестирование создания поста
func TestAddPost(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewStorePosts(db)
	repo := store.PostRepository()

	// Создание тестового пользователя
	query := `INSERT INTO users (name, email) VALUES('name_test', 'email_test');`
	_, err := db.Exec(query)
	require.NoError(t, err)

	post := &model.Post{
		Title:               "Test Post",
		Description:         "This is a test post",
		Author:              &model.User{ID: "1"},
		URL:                 "https://test_example.com",
		CreatedAt:           time.Now().Format(time.RFC3339),
		PermissionToComment: true,
	}

	err = repo.AddPost(post)
	require.NoError(t, err)
	assert.NotEmpty(t, post.ID)

	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM posts WHERE id = ?`, post.ID).Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

// Тестирование получения всех записей
func TestGetAllPosts(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewStorePosts(db)
	repo := store.PostRepository()

	// Создание тестовых пользователей
	query := `INSERT INTO users (name, email) VALUES('name1', 'email1');`
	_, err := db.Exec(query)
	require.NoError(t, err)

	query = `INSERT INTO users (name, email) VALUES('name2', 'email2');`
	_, err = db.Exec(query)
	require.NoError(t, err)

	// Создание тестовых записей
	post1 := &model.Post{
		Title:               "Test Post 1",
		Description:         "This is a test post 1",
		Author:              &model.User{ID: "1"},
		URL:                 "https://test_example1.com",
		CreatedAt:           time.Now().Format(time.RFC3339),
		PermissionToComment: true,
	}
	err = repo.AddPost(post1)
	require.NoError(t, err)
	assert.NotEmpty(t, post1.ID)

	post2 := &model.Post{
		Title:               "Test Post 2",
		Description:         "This is a test post 2",
		Author:              &model.User{ID: "2"},
		URL:                 "https://test_example2.com",
		CreatedAt:           time.Now().Format(time.RFC3339),
		PermissionToComment: true,
	}
	err = repo.AddPost(post2)
	require.NoError(t, err)
	assert.NotEmpty(t, post2.ID)

	// Сравнение
	expectedPosts := []*model.Post{post1, post2}
	actualPosts, err := repo.GetAllPosts()
	require.NoError(t, err)
	assert.Equal(t, expectedPosts, actualPosts)
}
