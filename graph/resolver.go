package graph

import (
	"fintech-app/graph/model"
	"math/rand"
	"strconv"
	"time"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	posts    []*model.Post
	comments []*model.Comment
}

func generateID() string {
	return strconv.Itoa(rand.Int())
}

func getCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}
