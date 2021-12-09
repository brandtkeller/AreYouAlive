package api

import (
	"encoding/json"
)

type Targets struct {
    Targets []Target `json:"targets"`
}


type Target struct {
    Id 		 int 	`json:"id"`
    Url 	 string `json:"url"`
	Interval int 	`json:"interval"`
}