package api

import (
	"net/http"
	"fmt"
    "strconv"
	"time"

    "github.com/gorilla/mux"

	"github.com/brandtkeller/AreYouAlive/pkg/common"
)

type TargetCheck struct {
	CheckTime	time.Time	`json: checkTime`
	Check		bool		`json: check`
}

type Target struct {
    Id 		 int 	`json:"id"`
    Url 	 string `json:"url"`
	Interval int 	`json:"interval"`
	Checks	 []TargetCheck	`json:checks`
}

func GetAllTargets( w http.ResponseWriter, req *http.Request, targets *[]Target ) {
	fmt.Println("Endpoint Hit: returnAllTargets")

	common.RespondWithJSON(w, http.StatusOK, &targets)
}

func TargetByIdHandler(w http.ResponseWriter, req *http.Request, targets *[]Target) {

	vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        common.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
        return
    }
	
	switch req.Method {
	case http.MethodGet:
        GetTargetById(w, req, id, targets)
    // case http.MethodPut:
    //     // PUT can be only by /target/{id} - single function required
	// case http.MethodDelete:
	// 	// DELETE can be only by /target/{id} - single function required
	default:
        common.RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
    }
}

func GetTargetById(w http.ResponseWriter, req *http.Request, key int, targets *[]Target ) {
	fmt.Println("Endpoint Hit: returnTargetById")
	var requestedTarget Target
	for _, target := range *targets {
        if target.Id == key {
            requestedTarget = target
        }
    }

	if requestedTarget.Id == 0 {
		common.RespondWithError(w, http.StatusNotFound, "Invalid Target ID")
	} else {
		common.RespondWithJSON(w, http.StatusOK, requestedTarget)
	}

}