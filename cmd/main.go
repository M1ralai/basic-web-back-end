package main

import (
	"github.com/M1iralai/deneme/cmd/server"
)

func main() {
	s := server.NewServer(":8080")
	s.RunServer()
}
