package graph

import (
	"database/sql"
	"fintech-app/graph/model"
	"fintech-app/storage/storeComments"
	"fintech-app/storage/storePosts"
	"math/rand"
	"strconv"
	"time"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	posts        []*model.Post
	comments     []*model.Comment
	PostStore    *storePosts.StorePosts
	CommentStore *storeComments.StoreComments
}

func NewResolver(db *sql.DB) *Resolver {
	return &Resolver{
		PostStore:    storePosts.NewStorePosts(db),
		CommentStore: storeComments.NewStoreComments(db),
	}
}

func generateID() string {
	return strconv.Itoa(rand.Int())
}

func getCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}
