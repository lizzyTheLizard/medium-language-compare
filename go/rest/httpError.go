package rest

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HttpError struct {
	msg        string
	statusCode int
	statusText string
	level      log.Level
}

func (e HttpError) Error() string {
	return e.msg
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	err := recover()
	if err == nil {
		return
	}
	httpError, ok := err.(HttpError)
	if !ok {
		log.Error("Could not handle request: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	log.StandardLogger().Logf(httpError.level, "Could not handle request: %v", httpError.Error())
	http.Error(w, httpError.statusText, httpError.statusCode)
}

func panicMethodNotAllowed(r *http.Request) {
	log.Panic()
	msg := fmt.Sprintf("Method %v not allowed on %v", r.Method, r.URL)
	panic(HttpError{msg, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), log.InfoLevel})
}

func panicBadRequest(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	panic(HttpError{msg, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), log.InfoLevel})
}

func panicNotFound(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	panic(HttpError{msg, http.StatusNotFound, http.StatusText(http.StatusNotFound), log.InfoLevel})
}

func panicInternalServerError(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	panic(HttpError{msg, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), log.WarnLevel})
}
