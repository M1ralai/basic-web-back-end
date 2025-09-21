package db

import (
	"database/sql"
	"log"

	"github.com/M1iralai/deneme/cmd/utils"
)

type Database struct {
	db     *sql.DB
	logger *log.Logger
}

func NewDB() *Database {
	return &Database{
		logger: utils.NewLogger("database"),
		db:     nil,
	}
}

func (db *Database) Initdb() {
	db.logger = utils.NewLogger("database")
	db.logger.Println("database successfully initialized")
}

//TODO add a function that drops and recreate users table from useres.db, unique id and username, encrypted password

//TODO add a function that  calls from a usrerHandler's post side, that will create a user with given specs

//TODO add a function that calls from a userHandler's delete side, that will drop a user with tat id

// gets userID oldPassword newPassword and securityAnswer for changing user's password, if oldPassword is empty the nit will check securityAnswer otherwise securithAnswer won't be used
func (db *Database) PatchUser(userID string, newPassword string, oldPassword string, securityAnswer string) error {
	return nil
	//TODO that will change the password of user, all control will be here and if there is error that function will return error
}

// gets username password then encrypt password from utilities package then check if there is a user like that, if not return 0 and error. username and passwords emptines and some other schnenigans will be handled in js of that form.
func (db *Database) LoginUser(username string, password string) (int, error) {
	return 0, nil
	//TODO fill this function with fully functional one
}

// gets username and password then check if username is taken if not encrypts password and register user to a db
func (db *Database) RegisterUser(username string, password string) error {
	return nil
}

func (db *Database) GetUserByID(userID string) *utils.User {
	return nil
}

func (db *Database) DeleteUser(userID string) {

}

func (db *Database) GetPosts(from int, destination int) ([]utils.User, error) {
	return nil, nil
}

func (db *Database) PostPost(title string, excerpt string, article string, userId string) error {
	// Create a post connectged to a userid and add creation time as dd/mm/yy do not add time for creation
	return nil
}

func (db *Database) PutPost(userId string, title string, newArticle string, newExcerpt string) error {
	//TODO check is userid is equal to a posts authors user id
	return nil
}

func (db *Database) DeletePost(userId string, title string) error {
	//TODO check if the request owner is a author of the post
	return nil
}
