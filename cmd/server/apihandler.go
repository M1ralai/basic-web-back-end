package server

import (
	"fmt"
	"net/http"
)

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//TODO thats will be for login, will check usrename and password
		fmt.Println("get method came")
	case http.MethodPost:
		//TODO thats will be for reister, will check CSRF_TOKEN
		fmt.Println("post method came")
	case http.MethodDelete:
		//TODO tahts will check session_id and user's password then delete
		fmt.Println("delete method came")
	case http.MethodPatch:
		//TODO thats will change user password so i don't know how to check
		fmt.Println("patch method came")
	default:
		http.Error(w, "503 - unauthorized request", 503)
		return
	}
}

//TODO add a posts handler
