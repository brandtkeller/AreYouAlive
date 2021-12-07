package main

import (
    "net/http"
	"log"
	"io"
	"fmt"
	"encoding/json"

    "github.com/gorilla/mux"
)

type target struct {
	Id string `json:"id"`
    Url string `json:"url"`
	Interval int `json:"interval"`
}

var Targets []target

// Database in-memory load
func LoadData() {

	Targets = append(Targets, target{Id: "1", Url: "test-one.com", Interval: 30})
	Targets = append(Targets, target{Id: "2", Url: "test-two.com", Interval: 60})

}


// --------- Main response logic ----------
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

// ---------- Health Check Functions ----------
func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
    // A very simple health check.
	fmt.Printf(req.Method)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // In the future we could report back on the status of our DB
    io.WriteString(w, `{"alive": true}`)
}

func GetAllTargetHandler( w http.ResponseWriter, req *http.Request ) {
	fmt.Println("Endpoint Hit: returnAllTargets")
    json.NewEncoder(w).Encode(Targets)
}

func getTargetById(w http.ResponseWriter, req *http.Request, key string ) {
	fmt.Println("Endpoint Hit: returnAllTargets")

	for _, target := range Targets {
        if target.Id == key {
            json.NewEncoder(w).Encode(target)
        }
    }


}

// ---------- Target Functions ----------
func TargetHandler(w http.ResponseWriter, req *http.Request) {
	
	vars := mux.Vars(req)
    key := vars["id"]
	
	switch req.Method {
    // case http.MethodPost:
    //     // POST can only be on creation - single function required
	case http.MethodGet:
        getTargetById(w, req, key)
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
	router.HandleFunc("/target/{id}", TargetHandler).Methods("POST", "GET", "PUT", "DELETE")
	

	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")
	fmt.Println("Running the server on port 8080")
	LoadData()

	for i := 0; i < len(Targets); i++ {
		fmt.Println(Targets[i].Url)
	}

    log.Fatal(http.ListenAndServe("localhost:8080", router))
}