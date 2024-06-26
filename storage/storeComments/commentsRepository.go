package storeComments

import (
	"database/sql"
	"fintech-app/graph/model"
	"fmt"
	"log"
	"strconv"
)

type CommentsRepository struct {
	db *sql.DB
}

func NewCommentsRepository(db *sql.DB) *CommentsRepository {
	return &CommentsRepository{db: db}
}

// Добавление комментария в бд
func (r *CommentsRepository) AddComment(c *model.Comment, post_id string) error {
	// Проверки на длину комментария
	if len(c.Description) >= 2000 {
		return fmt.Errorf("the maximum length of a comment description is more than or equal to 2000 characters")
	}

	if len(c.Description) == 0 {
		return fmt.Errorf("the minimum length of a comment description should be more than or equal to 1 character")
	}

	postID, _ := strconv.Atoi(post_id)
	query := `INSERT INTO comments (post_id, user_id, content, created_at) VALUES($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(query, postID, c.Author.ID, c.Description, c.CreatedAt).Scan(&c.ID)
	if err != nil {
		log.Printf("Ошибка при добавлении комментария в базу данных: %v", err)
		return err
	}

	ID, _ := strconv.Atoi(c.ID)
	log.Printf("Комментарий успешно добавлен с ID: %d", ID)
	return nil
}
