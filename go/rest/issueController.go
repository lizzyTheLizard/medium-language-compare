package rest

import (
	"lizzy/medium/compare/domain"
	"net/http"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type IssueController struct {
	repository domain.IssueRepository
}

func NewIssueController(repository domain.IssueRepository) IssueController {
	return IssueController{repository}
}

func (i IssueController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer errorHandler(w, r)

	id := parseId(r)
	//First check the cases where no post body is needed
	switch {
	case r.Method == http.MethodGet && id != nil:
		i.readSingle(*id, w)
		return
	case r.Method == http.MethodGet && id == nil:
		i.readAll(w)
		return
	case r.Method == http.MethodDelete && id != nil:
		i.delete(*id, w)
		return
	case r.Method == http.MethodDelete && id == nil:
		panicMethodNotAllowed(r)
	}

	postBody := parseBody(r)

	//Then the cases wherer a post body is needed
	switch {
	case r.Method == http.MethodPost && id == nil:
		i.create(postBody, w)
	case r.Method == http.MethodPut && id != nil:
		i.update(*id, postBody, w)
	case r.Method == http.MethodPatch && id != nil:
		i.partialUpdate(*id, postBody, w)
	default:
		panicMethodNotAllowed(r)
	}
}

func parseId(r *http.Request) *uuid.UUID {
	p := strings.Split(r.URL.Path, "/")
	if len(p) <= 2 || len(p[2]) == 0 {
		return nil
	}
	id, err := uuid.Parse(p[2])
	if err != nil {
		panicBadRequest("Cannot parse ID: %v", err)

	}
	return &id
}

func (i IssueController) readSingle(id uuid.UUID, w http.ResponseWriter) {
	issue, err := i.repository.Find(id)
	if err == domain.IssueNotFoundError {
		panicNotFound("Issue with ID %v cannot be found: %v", id, err)
	}
	if err != nil {
		panicInternalServerError("Cannot read issue %v: %v", id, err)
	}
	log.Debug("Issue %v returned", issue)
	issueDto := issueToIssueDto(issue)
	writeJson(w, issueDto)
}

func (i IssueController) readAll(w http.ResponseWriter) {
	issues, err := i.repository.FindAll()
	if err != nil {
		panicInternalServerError("Cannot read issues: %v", err)
	}
	log.Debugf("Issues %v returned", issues)
	var issueDtos []issueDto
	for _, issue := range issues {
		issueDtos = append(issueDtos, issueToIssueDto(issue))
	}
	writeJson(w, issueDtos)
}

func (i IssueController) create(postBody issueDto, w http.ResponseWriter) {
	newIssue := postBody.toIssue()
	err := i.repository.Insert(newIssue)
	if err != nil {
		panicInternalServerError("Cannot create issue: %v", err)
	}
	log.Infof("Create issue %v", newIssue)
	issueDto := issueToIssueDto(newIssue)
	writeJson(w, issueDto)
}

func (i IssueController) update(id uuid.UUID, postBody issueDto, w http.ResponseWriter) {
	newIssue := postBody.toIssue()
	err := i.repository.Update(newIssue)
	if err == domain.IssueNotFoundError {
		panicNotFound("Issue with ID %v cannot be found: %v", id, err)
	}
	if err != nil {
		panicInternalServerError("Cannot update issue %v: %v", id, err)
	}
	log.Infof("Update issue %v", newIssue)
	issueDto := issueToIssueDto(newIssue)
	writeJson(w, issueDto)
}

func (i IssueController) partialUpdate(id uuid.UUID, postBody issueDto, w http.ResponseWriter) {
	oldIssue, err := i.repository.Find(id)
	if err == domain.IssueNotFoundError {
		panicNotFound("Issue with ID %v cannot be found: %v", id, err)
	}
	if err != nil {
		panicInternalServerError("Cannot read issue %v for update: %v", id, err)
	}
	newIssue := oldIssue.Update(postBody.Name, postBody.Description)
	err = i.repository.Update(newIssue)
	if err != nil {
		panicInternalServerError("Cannot partially update issue %v: %v", id, err)
	}
	log.Infof("Update issue %v", newIssue)
	issueDto := issueToIssueDto(newIssue)
	writeJson(w, issueDto)
}

func (i IssueController) delete(id uuid.UUID, w http.ResponseWriter) {
	err := i.repository.Delete(id)
	if err == domain.IssueNotFoundError {
		panicNotFound("Issue with ID %v cannot be found: %v", id, err)
	}
	if err != nil {
		panicInternalServerError("Cannot delete issue %v: %v", id, err)
	}
	log.Infof("Delete issue %v", id)
	w.WriteHeader(http.StatusOK)
}
