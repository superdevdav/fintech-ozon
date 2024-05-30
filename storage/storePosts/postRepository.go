package storePosts

import (
	"database/sql"
	"fintech-app/graph/model"
	"log"
	"strconv"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

// Добавление поста в бд
func (r *PostRepository) AddPost(p *model.Post) error {
	query := `INSERT INTO posts (title, description, author_id, url, created_at, edit) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;`
	err := r.db.QueryRow(query, p.Title, p.Description, p.Author.ID, p.URL, p.CreatedAt, p.Edit).Scan(&p.ID)
	if err != nil {
		log.Printf("Ошибка при добавлении поста в базу данных: %v", err)
		return err
	}

	ID, _ := strconv.Atoi(p.ID)
	log.Printf("Пост успешно добавлен с ID: %d", ID)
	return nil
}

// Получение всех постов из бд
func (r *PostRepository) GetAllPosts() ([]*model.Post, error) {
	query := `SELECT id, title, description, author_id, url, created_at, edit FROM posts;`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Ошибка при получении всех постов: %v", err)
		return nil, err
	}
	defer rows.Close()

	var posts []*model.Post
	for rows.Next() {
		var post model.Post
		var authorID string

		err := rows.Scan(&post.ID, &post.Title, &post.Description, &authorID, &post.URL, &post.CreatedAt, &post.Edit)
		if err != nil {
			log.Printf("Ошибка при сканировании строки поста: %v", err)
			return nil, err
		}

		post.Author = &model.User{ID: authorID}

		comments, err := r.getCommentsForPost(post.ID)
		if err != nil {
			log.Printf("Ошибка при получении комментариев для поста с ID %s: %v", post.ID, err)
			return nil, err
		}
		post.Comments = comments

		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Ошибка после итерации по строкам: %v", err)
		return nil, err
	}

	return posts, nil
}

// Получение списка комментариев к посту
func (r *PostRepository) getCommentsForPost(postID string) ([]*model.Comment, error) {
	query := `SELECT user_id, content, created_at FROM comments WHERE post_id=$1;`
	rows, err := r.db.Query(query, postID)
	if err != nil {
		log.Printf("Ошибка при получении комментариев для поста: %v", err)
		return nil, err
	}
	defer rows.Close()

	var comments []*model.Comment
	for rows.Next() {
		var comment model.Comment
		var userID string

		err := rows.Scan(&userID, &comment.Description, &comment.CreatedAt)
		if err != nil {
			log.Printf("Ошибка при сканировании строки комментария: %v", err)
			return nil, err
		}

		// Получение автора комментария
		user, err := r.getUserByID(userID)
		if err != nil {
			log.Printf("Ошибка при получении данных автора: %v", err)
			return nil, err
		}

		comment.Author = user

		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Ошибка после итерации по строкам комментариев: %v", err)
		return nil, err
	}

	return comments, nil
}

// Метод для получения данных пользователя по ID
func (r *PostRepository) getUserByID(userID string) (*model.User, error) {
	query := `SELECT name, email FROM users WHERE id=$1`
	user := &model.User{ID: userID}

	err := r.db.QueryRow(query, userID).Scan(&user.Name, &user.Email)
	if err != nil {
		log.Printf("Ошибка при получении данных пользователя с ID %s: %v", userID, err)
		return nil, err
	}

	return user, nil
}
