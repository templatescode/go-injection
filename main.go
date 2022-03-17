package main

import (
	"fmt"
	"go-injection/database"
)

func main() {

	// instancia objeto
	db := database.NewDatabase("dynamoDB")
	fmt.Println(db)

	err := db.Create("new database")
	fmt.Println(err)
}
