package app

import "github.com/gorilla/mux"

func (s *Server) Routes() *mux.Router {
	router := mux.NewRouter()

	return router
}