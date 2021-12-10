package api

import (
	"net/http"
	"fmt"

	"github.com/brandtkeller/AreYouAlive/pkg/common"
)

// ---------- Health Check Functions ----------
// GET is the only acceptable Method, logging for posterity
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Endpoint Hit: " + req.Method + " /health")
    // In the future we could report back on the status of our DB
	common.RespondWithJSON(w, http.StatusOK, `{"alive": true}`)
}