package entities

import (
	"time"

	api_models "bookstore/api/models"
)

type Book struct {
	Id          int64
	Title       string
	Writer      string
	Pubication  string
	Number      int
	IsActivated bool
	IsDeleted   bool
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New_book_from_api_model(b *api_models.Book)*Book{
	//marshal and unmarshal model b to Book instance
	return &Book{}
}

func (b *Book) To_api_model() *api_models.Book {
	//marshal and unmarshal model b to api_models.Book instance
	return &api_models.Book{}
}


type Books []Book
