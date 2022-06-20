package model

import "time"

type Article struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	CreatedDate time.Time	`json:"created_date"`
	UpdatedDate	time.Time	`json:"updated_date"`
	Status	string	`json:"status"`
}