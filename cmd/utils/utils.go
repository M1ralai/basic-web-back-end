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

type User struct {
	ID       string `json:"userID" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Post struct {
	Title   string `json:"title" db:"title"`
	Article string `json:"article" db:"article"`
	Excerpt string `json:"excerpt" db:"excerpt"`
	Author  string `json:"author" db:"author"`
	Date    string `json:"date" db:"date"`
}

//TODO create a strcut as a posts for helping db and json return
