package models

type Books struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	AmountPage int    `json:"amount_page"`
}
