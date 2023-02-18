package repository

import (
	"bookstore/domain/entities"
)

type Book_repository interface {
	Create(book entities.Book)
	GetById(id int64) entities.Books
}
