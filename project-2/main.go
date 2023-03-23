package main

import (
	"project-2/database"
	"project-2/routers"
)

func main() {
	database.StartDB()
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
