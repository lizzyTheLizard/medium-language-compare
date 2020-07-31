package rest

import (
	"errors"
	"lizzy/medium/compare/go-pure/domain"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Engine struct{}

func (e Engine) Start() {
	http.ListenAndServe(":8080", nil)
}

var methodNotAllowed = errors.New("Method not allowed on path")

func NewEngine(issueController IssueController) Engine {
	http.HandleFunc("/issue/", func(w http.ResponseWriter, r *http.Request) {
		p := strings.Split(r.URL.Path, "/")
		hasIdInPath := len(p) > 2 && len(p[2]) != 0
		var err error
		switch {
		case !hasIdInPath && r.Method == http.MethodGet:
			err = issueController.readAll(w, r)
		case !hasIdInPath && r.Method == http.MethodPost:
			err = issueController.create(w, r)
		case hasIdInPath && r.Method == http.MethodGet:
			err = issueController.readSingle(w, r)
		case hasIdInPath && r.Method == http.MethodPut:
			err = issueController.update(w, r)
		case hasIdInPath && r.Method == http.MethodPatch:
			err = issueController.partialUpdate(w, r)
		case hasIdInPath && r.Method == http.MethodDelete:
			err = issueController.delete(w, r)
		default:
			err = methodNotAllowed
		}
		if err != nil {
			errorHandler(err, w)
		}
	})
	return Engine{}
}

func errorHandler(error error, w http.ResponseWriter) {
	switch error {
	case domain.IssueNotFoundError:
		log.Infof(error.Error())
		http.Error(w, error.Error(), http.StatusNotFound)
	case methodNotAllowed:
		log.Infof(error.Error())
		http.Error(w, error.Error(), http.StatusMethodNotAllowed)
	case nil:
	default:
		log.Warnf("Error while handling request: %v", error)
		http.Error(w, "Could not process request", http.StatusInternalServerError)
	}
}
