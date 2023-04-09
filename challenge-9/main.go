package main

import (
	"challenge-9/database"
	"challenge-9/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
