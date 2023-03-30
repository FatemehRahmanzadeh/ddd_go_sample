package repository

import (
	"bookstore/domain/entities"
)

type Book_repository interface {
	Create(book entities.Book)
	Update(book entities.Book)
	GetById(id int64) entities.Book
	Delete(id int64)
	Get_all() entities.Books
}
