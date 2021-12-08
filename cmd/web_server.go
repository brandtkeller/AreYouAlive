package main

import (
    "net/http"
	"log"
	"fmt"
	"encoding/json"
	"strconv"

    "github.com/gorilla/mux"
)

type target struct {
	Id int `json:"id"`
    Url string `json:"url"`
	Interval int `json:"interval"`
}

var Targets []target

// Load data from a configuration file
// This could later be refactored to populate a database instead of loading into memory
func LoadData() {

	Targets = append(Targets, target{Id: 1, Url: "test-one.com", Interval: 30})
	Targets = append(Targets, target{Id: 2, Url: "test-two.com", Interval: 60})

}


// --------- Main response logic ----------
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

// ---------- Health Check Functions ----------
// GET is the only acceptable Method, logging for posterity
func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Endpoint Hit: " + req.Method + " /health")
    // In the future we could report back on the status of our DB
	respondWithJSON(w, http.StatusOK, `{"alive": true}`)
}

func GetAllTargetHandler( w http.ResponseWriter, req *http.Request ) {
	fmt.Println("Endpoint Hit: returnAllTargets")

	respondWithJSON(w, http.StatusOK, Targets)
}

func getTargetById(w http.ResponseWriter, req *http.Request, key int ) {
	fmt.Println("Endpoint Hit: returnTargetById")
	var requestedTarget target
	for _, target := range Targets {
        if target.Id == key {
            requestedTarget = target
        }
    }

	if requestedTarget.Id == 0 {
		respondWithError(w, http.StatusNotFound, "Invalid Target ID")
	} else {
		respondWithJSON(w, http.StatusOK, requestedTarget)
	}

}

// ---------- Target Functions ----------
func TargetHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid product ID")
        return
    }
	
	switch req.Method {
	case http.MethodGet:
        getTargetById(w, req, id)
    // case http.MethodPut:
    //     // PUT can be only by /target/{id} - single function required
	// case http.MethodDelete:
	// 	// DELETE can be only by /target/{id} - single function required
	default:
		http.Error(w, "Bad request", http.StatusMethodNotAllowed)
		return
    }
}




func main() {
	router := mux.NewRouter()
	// router.HandleFunc("/", rootHandler).Methods("GET")

	// ----- GET all targets -----
	router.HandleFunc("/target", GetAllTargetHandler).Methods("GET")

	// Do something with a specific Target
	router.HandleFunc("/target/{id}", TargetHandler).Methods("GET")
	

	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")
	fmt.Println("Running the server on port 8080")
	LoadData()

    log.Fatal(http.ListenAndServe("localhost:8080", router))
}