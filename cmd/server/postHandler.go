package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/M1iralai/deneme/cmd/db"
)

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getPostsHandler(w, r)
	case http.MethodPost:
		s.postPostHandler(w, r)
	case http.MethodPut:
		s.postPutHandler(w, r)
	case http.MethodDelete:
		s.postDeleteHandler(w, r)
	default:
		http.Error(w, "unauthorized access try", http.StatusUnauthorized)
	}
}

func (s *Server) getPostsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	sFrom := query.Get("from")
	if sFrom == "" {
		http.Error(w, "there must be a 'from' as a variable", http.StatusBadRequest)
		return
	}

	sDest := query.Get("destination")
	if sFrom == "" {
		http.Error(w, "there must be a 'destination' as a variable", http.StatusBadRequest)
		return
	}

	from, err := strconv.Atoi(sFrom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	destination, err := strconv.Atoi(sDest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	posts, err := db.GetPosts(from, destination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, "invalid json request", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("posts successfully loaded"))
	w.WriteHeader(http.StatusOK)
}

func (s *Server) postPostHandler(w http.ResponseWriter, r *http.Request) {
	//TODO CSRF_TOKEN control will be donbe here

	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	title, ok := data["title"]
	if !ok || title == "" {
		http.Error(w, "title must be string and cannot be empty", http.StatusBadRequest)
		return
	}

	excerpt, ok := data["excerpt"]
	if !ok || excerpt == "" {
		http.Error(w, "excerpt must be string and cannot be empty", http.StatusBadRequest)
		return
	}

	article, ok := data["article"]
	if !ok || article == "" {
		http.Error(w, "article must be string and cannot be empty", http.StatusBadRequest)
		return
	}

	val, err := s.getSessionId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := sessionPack.Load(val)

	err = db.PostPost(title, excerpt, article, session.(Session).userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("post successfully created"))
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) postPutHandler(w http.ResponseWriter, r *http.Request) {

	//TODO there will be a CSRF_TOKEN checker

	var data map[string]string

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid json request", http.StatusBadRequest)
		return
	}

	title, ok := data["title"]
	if !ok {
		http.Error(w, "there must be a title variable", http.StatusBadRequest)
		return
	}

	article, ok := data["article"]
	if !ok {
		http.Error(w, "there must be a aricle variable", http.StatusBadRequest)
		return
	}

	excerpt, ok := data["excerpt"]
	if !ok {
		http.Error(w, "there must be a excerpt variable", http.StatusBadRequest)
		return
	}

	sessionId, err := s.getSessionId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	uid, ok := sessionPack.Load(sessionId)
	if !ok {
		http.Error(w, "there is no session_id as that", http.StatusUnauthorized)
		return
	}

	err = db.PutPost(uid.(Session).userId, title, article, excerpt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("post editing is successfully done"))
	w.WriteHeader(http.StatusOK)

}

func (s *Server) postDeleteHandler(w http.ResponseWriter, r *http.Request) {

	//TOOD CSRF_TOKEN checker

	var data map[string]string

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid json request", http.StatusBadRequest)
		return
	}

	title, _ := data["title"]

	sessionId, err := s.getSessionId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, ok := sessionPack.Load(sessionId)

	if !ok {
		http.Error(w, "there is no session id as you have", http.StatusInternalServerError)
		return
	}

	err = db.DeletePost(userId.(Session).userId, title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(" post successfully deleted "))
	w.WriteHeader(http.StatusOK)

}
