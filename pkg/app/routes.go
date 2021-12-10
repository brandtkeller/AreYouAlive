package app

import "github.com/gorilla/mux"

func (s *Server) RouterInit() *mux.Router {
	router := mux.NewRouter()

	// Room here to establish other common defaults to attach to the Router if applicable

	return router
}