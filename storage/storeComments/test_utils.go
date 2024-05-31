package storeComments

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Создание тестовых таблиц
	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(50) NOT NULL,
			email VARCHAR(100) NOT NULL
		);
		
		CREATE TABLE posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(100) NOT NULL,
			description TEXT NOT NULL,
			author_id INTEGER REFERENCES users(id),
			url VARCHAR(200),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			permission_to_comment BOOLEAN DEFAULT TRUE
		);

		CREATE TABLE comments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER REFERENCES posts(id),
			user_id INTEGER REFERENCES users(id),
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
    `)
	if err != nil {
		t.Fatalf("Failed to create test tables: %v", err)
	}

	return db
}
