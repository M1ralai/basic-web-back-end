package server

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.userGetHandler(w, r)
	case http.MethodPost:
		s.userPutHandler(w, r)
	case http.MethodDelete:
		s.userDeleteHandler(w, r)
	case http.MethodPatch:
		s.userPatchHandler(w, r)
	default:
		http.Error(w, "503 - unauthorized request", 503)
		return
	}
}

// Thats basicly a login function that gets parameters from url query
func (s *Server) userGetHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	username := query.Get("username")

	if username == "" {
		http.Error(w, "username cannot be blank", http.StatusBadRequest)
		return
	}

	password := query.Get("password")

	if password == "" {
		http.Error(w, "password can bot be blank", http.StatusBadRequest)
		return
	}

	userID, err := s.db.LoginUser(username, password)

	sid, err := s.getSessionId(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sessionPack.Swap(sid, userID)

	w.Write([]byte("successfully logged in"))
	w.WriteHeader(http.StatusOK)
}

// This function basicly just changes users password
func (s *Server) userPatchHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]any

	//TODO CSRFT_TOKEN control will be added here

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userID, ok := data["userID"].(string)
	if !ok {
		http.Error(w, "userID must be a string and cannot be blank", http.StatusBadRequest)
		return
	}

	sid, err := s.getSessionId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uid, _ := sessionPack.Load(sid)
	if uid.(Session).userId != uid {
		http.Error(w, "unauthorized access denied", http.StatusUnauthorized)
		return
	}

	newPassowrd, ok := data["newPassword"].(string)
	if !ok || newPassowrd == "" {
		http.Error(w, "newPassword must be string", http.StatusBadRequest)
		return
	}

	oldPassword, _ := data["oldPassword"].(string)

	securityQuestion, _ := data["securityAnswer"].(string)

	err = s.db.PatchUser(userID, newPassowrd, oldPassword, securityQuestion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Write([]byte("user password successfully changed"))
		w.WriteHeader(http.StatusOK)
		return
	}
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

	err := s.db.RegisterUser(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("user successfully created please login"))
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]any

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "invalid json request", http.StatusBadRequest)
		return
	}

	user_id, ok := data["userID"].(string)
	if !ok {
		http.Error(w, "userID must be string", http.StatusBadRequest)
		return
	}

	sessionId, _ := s.getSessionId(r)

	val, _ := sessionPack.Load(sessionId)

	if val.(Session).userId != user_id {
		http.Error(w, "unauthorized raquest", http.StatusUnauthorized)
		return
	}

	s.db.DeleteUser(user_id)

	w.Write([]byte("user successfuly deleted"))
	w.WriteHeader(http.StatusOK)
}

func (s *Server) getUserByID(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "unauthorized access denied", http.StatusUnauthorized)
		return
	}

	var id string

	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/") // ["api","users","123"]

	if len(parts) == 3 && parts[0] == "api" && parts[1] == "users" {
		id = parts[2]
	} else {
		http.NotFound(w, r)
	}

	u := s.db.GetUserByID(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)

}
