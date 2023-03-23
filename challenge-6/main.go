package main

import (
	"challenge-6/routers"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func main() {
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
