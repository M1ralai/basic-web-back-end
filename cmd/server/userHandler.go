package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/M1iralai/deneme/cmd/db"
)

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.userGetHandler(w, r)
		fmt.Println("get method came")
	case http.MethodPost:
		s.userPutHandler(w, r)
		fmt.Println("post method came")
	case http.MethodDelete:
		//TODO tahts will check session_id and user's password then delete
		fmt.Println("delete method came")
	case http.MethodPatch:
		s.userPatchHandler(w, r)
	default:
		http.Error(w, "503 - unauthorized request", 503)
		return
	}
}

// This function basicly just changes users password
func (s *Server) userPatchHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]any

	//TODO CSRFT_TOKEN control will be added here

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userId, ok := data["userId"].(float64)
	if !ok {
		http.Error(w, "userId must be a integer", http.StatusBadRequest)
		return
	}

	sid, err := s.getSessionId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uid, _ := sessionPack.Load(sid)
	if uid != int(userId) {
		http.Error(w, "unauthorized access denied", http.StatusUnauthorized)
	}

	newPassowrd, ok := data["newPassword"].(string)
	if !ok || newPassowrd == "" {
		http.Error(w, "newPassword must be string", http.StatusBadRequest)
		return
	}

	oldPassword, _ := data["oldPassword"].(string)

	securityQuestion, _ := data["securityAnswer"].(string)

	err = db.PatchUser(int(userId), newPassowrd, oldPassword, securityQuestion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("user password successfully changed"))
		return
	}
}

// Thats basicly a login function
func (s *Server) userGetHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]any

	//TODO CSRFT_TOKEN control will be added here

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	username, ok := data["username"].(string)
	if !ok {
		http.Error(w, "username must be string", http.StatusBadRequest)
		return
	}

	password, ok := data["password"].(string)
	if !ok {
		http.Error(w, "password must be string", http.StatusBadRequest)
		return
	}

	userID, err := db.LoginUser(username, password)

	sid, err := s.getSessionId(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sessionPack.Swap(sid, userID)
}

// thats just basicly a register function, just cheks is there a username password and csrftoken, form controls will be in dthe js side and db checks wiull be db side
func (s *Server) userPutHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]any

	//TODO CSRFTOKEN checker will be added here

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	username, ok := data["username"].(string)
	if !ok {
		http.Error(w, "username field must be string and cannot be blank", http.StatusBadRequest)
		return
	}

	password, ok := data["password"].(string)
	if !ok {
		http.Error(w, "password field must be string and cannot be blank", http.StatusBadRequest)
		return
	}

	err := db.RegisterUser(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (s *Server) getUserByID(w http.ResponseWriter, r *http.Request) {
	var strid string

	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/") // ["api","users","123"]
	if len(parts) == 3 && parts[0] == "api" && parts[1] == "users" {
		strid = parts[2]
	} else {
		http.NotFound(w, r)
	}

	id, err := strconv.Atoi(strid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := db.GetUserByID(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

//TODO add a posts handler
