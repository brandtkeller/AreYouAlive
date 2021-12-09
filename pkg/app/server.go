package app

import (
	"github.com/gorilla/mux"

	"AreYouAlive/pkg/api"
)

type Server struct {
	router        *mux.Router
	targetService api.TargetService
}

func NewServer(router *mux.Router, targetService api.TargetService) *Server {
	return &Server{
		router:        router,
		targetService: targetService,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()
	
	// run the server through the router
	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}