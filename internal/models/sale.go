package models

type Sale struct {
	ID      string  `json:"id"`
	Date    string  `json:"date"`
	Product string  `json:"product"`
	Total   float32 `json:"total"`
}

type Sales []*Sale
