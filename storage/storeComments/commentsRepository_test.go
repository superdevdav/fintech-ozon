package storeComments

import (
	"fintech-app/graph/model"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тестирование создания комментария
func TestAddComment(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewStoreComments(db)
	repo := store.CommentsRepository()

	// Создание тестового пользователя
	query := `INSERT INTO users (name, email) VALUES('name_test', 'email_test');`
	_, err := db.Exec(query)
	require.NoError(t, err)

	// Создание тестовой записи
	var post_id string
	query = `INSERT INTO posts (title, description, author_id, url, created_at, permission_to_comment) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;`
	_ = db.QueryRow(query, "Test post", "This is a test post", "1", "https://test_example.com", time.Now().Format(time.RFC3339), "true").Scan(&post_id)

	// Создание тестового комментария
	comment := &model.Comment{
		Description: "Comment content",
		Author: &model.User{
			ID:    "2",
			Name:  "name_test",
			Email: "email_test",
		},
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err = repo.AddComment(comment, post_id)
	require.NoError(t, err)
}

// Тестирование создания комментария >= 2000 символов
func TestAddCommentMore2000Symbols(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewStoreComments(db)
	repo := store.CommentsRepository()

	// Создание тестового пользователя
	query := `INSERT INTO users (name, email) VALUES('name_test', 'email_test');`
	_, err := db.Exec(query)
	require.NoError(t, err)

	// Создание тестовой записи
	var post_id string
	query = `INSERT INTO posts (title, description, author_id, url, created_at, permission_to_comment) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;`
	_ = db.QueryRow(query, "Test post", "This is a test post", "1", "https://test_example.com", time.Now().Format(time.RFC3339), "true").Scan(&post_id)

	// Создание тестового комментария с длиной в 2000 символов
	comment := &model.Comment{
		Description: strings.Repeat("a", 2000),
		Author: &model.User{
			ID:    "2",
			Name:  "name_test",
			Email: "email_test",
		},
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err = repo.AddComment(comment, post_id)
	assert.Equal(t, fmt.Errorf("the maximum length of a comment description is more than or equal to 2000 characters"), err)

	// Создание тестового комментария с длиной больше 2000 символов
	comment = &model.Comment{
		Description: strings.Repeat("a", 2100),
		Author: &model.User{
			ID:    "2",
			Name:  "name_test",
			Email: "email_test",
		},
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err = repo.AddComment(comment, post_id)
	assert.Equal(t, fmt.Errorf("the maximum length of a comment description is more than or equal to 2000 characters"), err)
}

// Тестирование пустого комментария
func TestAddCommentEmpty(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewStoreComments(db)
	repo := store.CommentsRepository()

	// Создание тестового пользователя
	query := `INSERT INTO users (name, email) VALUES('name_test', 'email_test');`
	_, err := db.Exec(query)
	require.NoError(t, err)

	// Создание тестовой записи
	var post_id string
	query = `INSERT INTO posts (title, description, author_id, url, created_at, permission_to_comment) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;`
	_ = db.QueryRow(query, "Test post", "This is a test post", "1", "https://test_example.com", time.Now().Format(time.RFC3339), "true").Scan(&post_id)

	// Создание тестового комментария
	comment := &model.Comment{
		Description: "",
		Author: &model.User{
			ID:    "2",
			Name:  "name_test",
			Email: "email_test",
		},
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err = repo.AddComment(comment, post_id)
	assert.Equal(t, fmt.Errorf("the minimum length of a comment description should be more than or equal to 1 character"), err)
}
