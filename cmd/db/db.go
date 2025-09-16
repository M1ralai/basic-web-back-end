package db

import (
	"log"

	"github.com/M1iralai/deneme/cmd/utils"
)

var logger *log.Logger

func Initdb() {
	logger = utils.NewLogger("database")
	logger.Println("database successfully initialized")
}
