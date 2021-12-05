package main

import (
    "net/http"
	"log"
	"io"
	"fmt"

    "github.com/gorilla/mux"
)

type target struct {
	ID int `json:"id"`
    Url string `json:"url"`
	Interval int `json:"interval"`
}

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
    // A very simple health check.
	fmt.Printf(req.Method)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // In the future we could report back on the status of our DB
    // by performing a simple PING, and include them in the response.
    io.WriteString(w, `{"alive": true}`)
}

func TargetHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
    case http.MethodPost:
        // POST can only be on creation - single function required
	case http.MethodGet:
        // Need to check for get all or get by ID
    case http.MethodPut:
        // PUT can be only by /target/{id} - single function required
	case http.MethodDelete:
		// DELETE can be only by /target/{id} - single functino required
	default:
		http.Error(w, "Bad request", http.StatusMethodNotAllowed)
		return
    }
}



func main() {
	router := mux.NewRouter()
	// router.HandleFunc("/", rootHandler).Methods("GET")
	//router.HandleFunc("/target", TargetHandler).Methods("POST", "GET", "PUT", "DELETE")

	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")
	fmt.Printf("Running the server on port 8080\n")

    log.Fatal(http.ListenAndServe("localhost:8080", router))
}