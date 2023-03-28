package db_models

import (
	"bookstore/domain/entities"
	"time"
)

type Book struct {
	Id          int64     `db:"id"`
	Title       string    `db:"title"`
	Writer      string    `db:"writer"`
	Pubication  string    `db:"publication"`
	Number      int       `db:"number"`
	IsActivated bool      `db:"isActivated"`
	IsDeleted   bool      `db:"isDeleted"`
	PublishedAt time.Time `db:"publishedAt"`
	CreatedAt   time.Time `db:"createdAt"`
	UpdatedAt   time.Time `db:"updatedAt"`
}

func New_book_from_entity(b *entities.Book) *Book {
	//marshal and unmarshal model b to Book instance
	return &Book{}
}

func (b *Book) To_entity() *entities.Book {
	//marshal and unmarshal model b to entities.Book instance
	return &entities.Book{}
}

type Books []Book
