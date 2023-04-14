package main

import (
	"challenge-10/database"
	"challenge-10/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
