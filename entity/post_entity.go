package entity

import "time"

type Post struct {
	ID        int
	UserID    int
	Title     string
	Excerpt   string
	ImageURL  *string
	Body      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
