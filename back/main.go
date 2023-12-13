package main

import (
	"back/app"
	"back/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}