package server

import "net/http"

func (s *Server) requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Printf(" %s ip user send %s method request to a %s URL", r.RemoteAddr, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

//TODO add a function that creates session_id as a cookie

//TODO add a funtcion that checks session_id from a cookie

//TODO add a function that create CSRF_TOKEN for a login page

//TODO add a function that checks CSRF_TOKEN from a login page
