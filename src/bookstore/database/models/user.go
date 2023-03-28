package db_models

type User struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	LastName    string `db:"lastName"`
	PhoneNumber int    `db:"phoneNumber"`
}

type Users []User