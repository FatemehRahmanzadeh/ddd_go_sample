package entities

import "time"

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

type Books []Book
