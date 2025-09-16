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

// gets userID oldPassword newPassword and securityAnswer for changing user's password, if oldPassword is empty the nit will check securityAnswer otherwise securithAnswer won't be used
func PatchUser(userID int, newPassword string, oldPassword string, securityAnswer string) error {
	return nil
	//TODO that will change the password of user, all control will be donw on handler side
}
