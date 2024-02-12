package main

import (
	"github.com/imperatrona/go-inertia-template/server"
	"github.com/petaki/inertia-go"
)

func main() {
	server := server.NewServer(inertia.New("http://localhost:1989", "./dist/index.html", ""))
	server.Serve()
}
