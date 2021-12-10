package app

import (
	"net/http"
	"log"

	"github.com/gorilla/mux"

	"github.com/brandtkeller/AreYouAlive/pkg/api"
)

type Server struct {
	router	*mux.Router
	targets	[]api.Target
}

func NewServer(targets []api.Target) *Server {
	r := mux.NewRouter()
	
	return &Server{
		router:  r,
		targets: targets,
	}
}

func (s *Server) Run() error {
	// run function that initializes the router
	router := s.router

	// router.HandleFunc("/", rootHandler).Methods("GET")

	// // ----- GET all targets -----
	router.HandleFunc("/target", func(w http.ResponseWriter, r *http.Request) {
		api.GetAllTargets(w, r, &s.targets)
	}).Methods("GET")

	// // Do something with a specific Target
	router.HandleFunc("/target/{id}", func(w http.ResponseWriter, r *http.Request) {
		api.TargetByIdHandler(w, r, &s.targets)
	}).Methods("GET")

	// Health check GET endpoint
	router.HandleFunc("/health", api.HealthCheck).Methods("GET")
	
    log.Fatal(http.ListenAndServe("localhost:8080", router))

	return nil
}