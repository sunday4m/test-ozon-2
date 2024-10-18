package logger

import (
	"log"
	"test-ozon-2/interfaces"
)

type Log struct{}

func New() interfaces.Logger {
	return &Log{}
}

func (l *Log) LogError(message string) {
	log.Fatalf("%s", message)
}

func (l *Log) LogInfo(message string) {
	log.Println(message)
}
