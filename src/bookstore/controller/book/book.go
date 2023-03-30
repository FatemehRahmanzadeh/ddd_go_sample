package book_controller

import (
	"bookstore/domain/entities"
	"bookstore/domain/repository"
)

type Book_controller interface {
	Create(book entities.Book)
	Update(book entities.Book)
	GetById(id int64) entities.Book
	Delete(id int64)
	Get_all() entities.Books
}

type book_controller struct {
	repository repository.Book_repository
}

func New_book_controller(r repository.Book_repository) Book_controller {
	return &book_controller{
		repository: r,
	}
}

func (b *book_controller) Create(book entities.Book) {
}
func (b *book_controller) Update(book entities.Book) {

}
func (b *book_controller) GetById(id int64) entities.Book {
	return b.repository.GetById(id)
}
func (b *book_controller) Delete(id int64) {
	b.repository.Delete(id)
}
func (b *book_controller) Get_all() entities.Books {
	return b.repository.Get_all()
}
