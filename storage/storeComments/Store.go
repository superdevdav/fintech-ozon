package storeComments

import "database/sql"

type StoreComments struct {
	db                 *sql.DB
	commentsRepository *CommentsRepository
}

func NewStoreComments(db *sql.DB) *StoreComments {
	return &StoreComments{
		db:                 db,
		commentsRepository: NewCommentsRepository(db),
	}
}

func (s *StoreComments) CommentsRepository() *CommentsRepository {
	return s.commentsRepository
}
