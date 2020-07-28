package rest

import (
	"errors"
	"lizzy/medium/compare/go-pure/domain"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var methodNotAllowed = errors.New("Method not allowed on path")

func Start() {
	http.HandleFunc("/issue/", issueHandler)
	http.ListenAndServe(":8080", nil)
}

func handleError(error error, w http.ResponseWriter) {
	switch error {
	case domain.IssueNotFoundError:
		log.Infof(error.Error())
		http.Error(w, error.Error(), http.StatusNotFound)
	case methodNotAllowed:
		log.Infof(error.Error())
		http.Error(w, error.Error(), http.StatusMethodNotAllowed)
	default:
		log.Warnf("Error during request %v", error.Error())
		http.Error(w, "Could not handle request", http.StatusInternalServerError)
	}
}
