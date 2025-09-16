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

//TODO add a function that drops and recreate users table from useres.db, unique id and username, encrypted password

//TODO add a function that  calls from a usrerHandler's post side, that will create a user with given specs

//TODO add a function that calls from a userHandler's delete side, that will drop a user with tat id

//TODO add a fubnction that calls from a userHandler's patch side, that will change the password of user
