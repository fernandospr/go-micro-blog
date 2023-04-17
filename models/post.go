package models

type Post struct {
	ID   uint   `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
}
