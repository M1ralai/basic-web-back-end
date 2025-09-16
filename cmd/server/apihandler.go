package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/M1iralai/deneme/cmd/db"
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
		userPatchHandler(w, r)
	default:
		http.Error(w, "503 - unauthorized request", 503)
		return
	}
}

func userPatchHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]any

	//TODO sessionID and CSRF_TOKEN control will be added here

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userId, ok := data["userId"].(float64)
	if !ok {
		http.Error(w, "userId must be a integer", http.StatusBadRequest)
		return
	}

	newPassowrd, ok := data["newPassword"].(string)
	if !ok || newPassowrd == "" {
		http.Error(w, "newPassword must be string", http.StatusBadRequest)
		return
	}

	oldPassword, _ := data["oldPassword"].(string)

	securityQuestion, _ := data["securityAnswer"].(string)

	err := db.PatchUser(int(userId), newPassowrd, oldPassword, securityQuestion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("user password successfully changed"))
		return
	}
}

//TODO add a posts handler
