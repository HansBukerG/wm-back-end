package main

import (
	"log"

	app "github.com/HansBukerG/wm-back-end/src"
	"github.com/joho/godotenv"
)

/*
	Environment variables:
		HTTP_PORT
		MONGO_HOST
		MONGO_PORT
*/

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	
	app.App_init()
}
