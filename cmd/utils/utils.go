package utils

import (
	"log"
	"os"
)

func NewLogger(serviceName string) *log.Logger {
	file, err := os.OpenFile("logs/"+serviceName+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("log file cannot opened, check permissions or file is corrupted")
		return nil
	}
	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Ldate|log.Lshortfile)
	return logger
}
