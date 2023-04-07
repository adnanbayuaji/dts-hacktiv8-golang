package main

import (
	"challenge-8/database"
	"challenge-8/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
