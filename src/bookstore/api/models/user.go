package api_models

type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"lastName"`
	PhoneNumber int    `json:"phoneNumber"`
}

type Users []User