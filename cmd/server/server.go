package server

import (
	"log"
	"net/http"

	"github.com/M1iralai/deneme/cmd/db"
	"github.com/M1iralai/deneme/cmd/utils"
)

type Server struct {
	serverAddr string
	logger     *log.Logger
	db         *db.Database
	mux        http.ServeMux
}

func NewServer(addr string) *Server {
	return &Server{
		serverAddr: addr,
		logger:     utils.NewLogger("server"),
		mux:        *http.NewServeMux(),
		db:         db.NewDB(),
	}
}

func (s *Server) RunServer() {
	go s.periodicSessionClear()
	//back-end is streaming here

	//BACK-END USER API START
	s.mux.Handle("/api/users/", s.requestLogger(s.sessionIdApi(http.HandlerFunc(s.getUserByID))))
	s.mux.Handle("/api/users", s.requestLogger(s.sessionIdApi(http.HandlerFunc(s.userHandler))))
	//BADCK-END USER API END

	//BACK-END POST API START
	s.mux.Handle("/api/posts", s.requestLogger(s.sessionIdApi(http.HandlerFunc(s.postHandler))))
	//BACK-END POST API END

	//front-end is streaming here

	s.mux.Handle("/", s.requestLogger(s.sessionIdFileServer(http.FileServer(http.Dir(".web/html")))))
	s.mux.Handle("/js/", s.requestLogger(s.sessionIdFileServer(http.StripPrefix("/js/", http.FileServer(http.Dir(".web/js"))))))

	// server is streaming right here

	s.logger.Printf("server is started at %s port", s.serverAddr)
	err := http.ListenAndServe(s.serverAddr, &s.mux)
	if err != nil {
		s.logger.Fatal(err)
		log.Fatal(err)
	}
}
