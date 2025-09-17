package server

import (
	"net/http"
)

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//TODO there will be function that gets posts from n to k, most recent post will be first then second most recent follows that...
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
