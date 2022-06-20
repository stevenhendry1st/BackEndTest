package main

import (
	"backend-test/app"
	"backend-test/model"
	"log"
)

func main() {
	db, err := app.Connect()

	log.Println("ERROR CONNECT DB DUE TO: ", err)

	db.AutoMigrate(
		&model.Article{},
	)

	app.InitRouter()
}