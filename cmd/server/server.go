package server

import (
	"log"
	"net/http"

	"github.com/M1iralai/deneme/cmd/utils"
)

type Server struct {
	serverAddr string
	logger     *log.Logger
	mux        http.ServeMux
}

func NewServer(addr string) *Server {
	return &Server{
		serverAddr: addr,
		logger:     utils.NewLogger("server"),
		mux:        *http.NewServeMux(),
	}
}

func (s *Server) RunServer() {
	//back-end is streaming here

	//TODO when request came from /api/users/:id that wil return this user's data
	s.mux.Handle("/api/users", s.requestLogger(http.HandlerFunc(s.userHandler)))

	//front-end is streaming here

	s.mux.Handle("/", s.requestLogger(http.FileServer(http.Dir(".web/html"))))

	// server is streaming right here

	err := http.ListenAndServe(s.serverAddr, &s.mux)
	s.logger.Printf("server is started at %s port", s.serverAddr)
	if err != nil {
		s.logger.Fatal(err)
		log.Fatal(err)
	}
}
