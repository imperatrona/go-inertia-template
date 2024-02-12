package server

import (
	"net/http"

	"github.com/imperatrona/go-inertia-template/server/handles"
	"github.com/petaki/inertia-go"
)

type Server struct {
	inertiaManager *inertia.Inertia
	handler        *handles.Handler
}

func NewServer(inertiaManager *inertia.Inertia) *Server {
	return &Server{
		inertiaManager: inertiaManager,
		handler:        handles.NewHandler(inertiaManager),
	}
}

func (s *Server) Serve() {
	mux := s.routes()
	http.ListenAndServe("localhost:1989", mux)
}
