package processor

import (
	"encoding/json"
	"errors"
	"test-ozon-2/interfaces"
	"test-ozon-2/models"
)

type Processor struct {
	log interfaces.Logger
	db  interfaces.Database
}

func New(log interfaces.Logger, db interfaces.Database) *Processor {
	return &Processor{
		log: log,
		db:  db,
	}
}
func (p *Processor) ProcessDocument(jsonData []byte) error {
	var doc models.Document

	err := json.Unmarshal(jsonData, &doc)
	if err != nil {
		p.log.LogError("Invalid JSON data")
		return errors.New(err.Error()) // я немного переделал тут вызов ошибки, чтобы в тесте проверить
	}

	if doc.Header == "" || len(doc.LineItems) == 0 {
		p.log.LogError("Missing header or line items")
		return errors.New("validation error")
	}

	err = p.db.WriteToDatabase(doc)
	if err != nil {
		p.log.LogError("Database write error")
		return err
	}

	p.log.LogInfo("Document processed successfully")
	return nil
}
