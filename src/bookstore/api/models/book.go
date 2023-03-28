package api_models

import "time"

type Book struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Writer      string    `json:"writer"`
	Pubication  string    `json:"publication"`
	Number      int       `json:"number"`
	IsActivated bool      `json:"isActivated"`
	IsDeleted   bool      `json:"isDeleted"`
	PublishedAt time.Time `json:"publishedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Books []Book