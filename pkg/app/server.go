package app

import (
	"net/http"
	"log"
	"fmt"
	"time"

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

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GET /")
    http.ServeFile(w, r, "web/build/index.html")
}

func (s *Server) Run() error {
	// run function that initializes the router
	router := s.router

	router.HandleFunc("/", index).Methods("GET")

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

	buildHandler := http.FileServer(http.Dir("web/build"))
    router.PathPrefix("/").Handler(buildHandler)

    srv := &http.Server{
        Handler:      router,
        Addr:         "0.0.0.0:3000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    fmt.Println("Server started on PORT 3000")
    log.Fatal(srv.ListenAndServe())
	return nil
}