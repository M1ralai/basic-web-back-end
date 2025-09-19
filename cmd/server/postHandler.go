package server

import (
	"encoding/json"
	"net/http"

	"github.com/M1iralai/deneme/cmd/db"
)

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getPosts(w, r)
	case http.MethodPost:
		//TODO there will be function that just creates post, its will add atimestamp
	case http.MethodPut:
		//TODO there will be function that just get the changes and csrf_token control other controls will be in de db, control is just is the request owner really a owner of post
	case http.MethodDelete:
		//TODO there will be function that just get the changes and csrf_token control other controls will be in de db, control is just is the request owner really a owner of post
	default:
		http.Error(w, "unauthorized access try", http.StatusUnauthorized)
	}
}

func (s *Server) getPosts(w http.ResponseWriter, r *http.Request) {
	var data map[string]float64

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid json request", http.StatusBadRequest)
		return
	}
	fFrom, ok := data["from"]
	if !ok {
		http.Error(w, "from must be integer", http.StatusBadRequest)
		return
	}

	fDest, ok := data["destination"]
	if !ok {
		http.Error(w, " destination must be integer", http.StatusBadRequest)
		return
	}

	posts, err := db.GetPosts(int(fFrom), int(fDest))
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

func (s *Server) postPost(w http.ResponseWriter, r *http.Request) {
	//TODO CSRF_TOKEN control will be donbe here

	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(data)
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
