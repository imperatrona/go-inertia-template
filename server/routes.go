package server

import (
	"net/http"

	"github.com/justinas/alice"
)

func (s *Server) routes() http.Handler {
	middleware := alice.New(
		s.inertiaManager.Middleware,
	)
	mux := http.NewServeMux()

	mux.Handle("/", middleware.ThenFunc(s.handler.Home))
	mux.Handle("/about", middleware.ThenFunc(s.handler.About))

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./dist/assets/"))))
	return middleware.Then(mux)
}
