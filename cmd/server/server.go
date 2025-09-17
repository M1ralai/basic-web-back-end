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
	go s.periodicSessionClear()
	//back-end is streaming here

	s.mux.Handle("/api/users/", s.requestLogger(s.sessionIdApi(http.HandlerFunc(s.getUserByID))))
	s.mux.Handle("/api/users", s.requestLogger(s.sessionIdApi(http.HandlerFunc(s.userHandler))))

	//front-end is streaming here

	s.mux.Handle("/", s.requestLogger(s.sessionIdFileServer(http.FileServer(http.Dir(".web/html")))))
	s.mux.Handle("/js/", s.requestLogger(s.sessionIdFileServer(http.StripPrefix("/js/", http.FileServer(http.Dir(".web/js"))))))

	// server is streaming right here

	err := http.ListenAndServe(s.serverAddr, &s.mux)
	s.logger.Printf("server is started at %s port", s.serverAddr)
	if err != nil {
		s.logger.Fatal(err)
		log.Fatal(err)
	}
}
