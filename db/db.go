package db

import (
	"fmt"
	"test-ozon-2/interfaces"
	"test-ozon-2/models"
)

type Database struct{}

func New() interfaces.Database {
	return &Database{}
}

// просто написал заглушку для функции
func (db *Database) WriteToDatabase(doc models.Document) error {
	if doc.Header == "Error" {
		return fmt.Errorf("header has error")
	}
	return nil
}
