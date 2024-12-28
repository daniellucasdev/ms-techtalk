package event

import "github.com/braiphub/go-scaffold/internal/domain/entity"

const (
	BookCreated    = "book:created"
	ChapterCreated = "chapter:created"
)

type BookCreatedEvent struct {
	ID   uint
	Name string
}

type ChapterCreatedEvent struct {
	Chapter entity.Chapter
}
