package main

import (
	"github.com/M1iralai/deneme/cmd/db"
	"github.com/M1iralai/deneme/cmd/server"
)

func main() {
	db.Initdb()
	s := server.NewServer(":8080")
	s.RunServer()
}
