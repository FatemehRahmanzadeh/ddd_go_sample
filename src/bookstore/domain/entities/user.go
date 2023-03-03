package entities

type User struct {
	Id int64
	Name string
	LastName string
	PhoneNumber int
}

type Users []User