package aoreblg

import (
	"time"
)

type Posts struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	Title        string    `json:"title"`
	PreviewText  string    `json:"previewtext"`
	PreviewImage string    `json:"previewimage"`
	Text         string    `json:"text"`
	Date         time.Time `json:"date"`
}
