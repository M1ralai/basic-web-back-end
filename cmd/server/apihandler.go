package server

import (
	"fmt"
	"net/http"
)

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println("get method came")
	case http.MethodPost:
		fmt.Println("post method came")
	case http.MethodDelete:
		fmt.Println("delete method came")
	case http.MethodPatch:
		fmt.Println("patch method came")
	default:
		http.Error(w, "503 - unauthorized request", 503)
		return
	}
}
