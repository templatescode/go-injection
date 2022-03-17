package database

import "fmt"

func (d databaseImp) Create(db string) error {
	fmt.Println("db impl:", d)
	fmt.Println("db create:", db)
	return nil
}
