package main

import (
	"fmt"
	"test-ozon-2/db"
	"test-ozon-2/logger"
	"test-ozon-2/processor"
)

func main() {
	log := logger.New()
	db := db.New()
	processor := processor.New(log, db)
	// Пример JSON-документа для тестирования
	jsonData := []byte(`{
		"header": "Sample Header",
		"line_items": ["item1", "item2", "item3"]
	}`)

	// Обрабатываем документ
	err := processor.ProcessDocument(jsonData)
	if err != nil {
		fmt.Println("Error processing document:", err)
	} else {
		fmt.Println("Document processed successfully!")
	}
}
