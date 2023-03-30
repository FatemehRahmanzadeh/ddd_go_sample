package persistance

import (
	"bookstore/domain/entities"

	abstraction_database "swap-gitlab.apk-group.net/apkswap/apk_go/database_layer/abstraction/database"
)

type Book_repository struct {
	db abstraction_database.Imysql_abstraction
}

func Create(book entities.Book) {

}
func Update(book entities.Book) {

}
func GetById(id int64) entities.Book {
return entities.Book{}
}
func Delete(id int64) {

}
func Get_all() entities.Books {
return entities.Books{}
}
