package dto

import "mime/multipart"

type PostRequest struct {
	ID      int                   `json:"id"`
	Title   string                `json:"title"`
	Image   *multipart.FileHeader `json:"image"`
	Excerpt string                `json:"excerpt"`
	Body    string                `json:"body"`
}

type PostResponse struct {
	ID       int    `json:"-"`
	Title    string `json:"title"`
	ImageURL *string `json:"image_url"`
	Excerpt  string `json:"excerpt"`
	Body     string `json:"body"`
}
