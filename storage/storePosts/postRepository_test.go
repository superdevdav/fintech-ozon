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
