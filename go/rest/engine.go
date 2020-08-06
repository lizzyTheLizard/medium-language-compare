package rest

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Engine struct{}

func (e Engine) Start() {
	http.ListenAndServe(":8080", nil)
}

func NewEngine(issueController IssueController) Engine {
	http.Handle("/issue/", issueController)
	return Engine{}
}

func writeJson(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		log.Panic("Cannot marshall data: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
