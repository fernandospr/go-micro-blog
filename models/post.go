package models

type Post struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	User string `json:"user"`
	Text string `json:"text"`
}
