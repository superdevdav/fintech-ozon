package storePosts

import "database/sql"

type StorePosts struct {
	db              *sql.DB
	postsRepository *PostRepository
}

func NewStorePosts(db *sql.DB) *StorePosts {
	return &StorePosts{
		db:              db,
		postsRepository: NewPostRepository(db),
	}
}

func (s *StorePosts) PostRepository() *PostRepository {
	return s.postsRepository
}
