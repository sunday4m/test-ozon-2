package interfaces

import (
	"test-ozon-2/models"
)

type Logger interface {
	LogError(message string)
	LogInfo(message string)
}

type Database interface {
	WriteToDatabase(doc models.Document) error
}
