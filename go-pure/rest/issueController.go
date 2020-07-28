package rest

import (
	"encoding/json"
	"lizzy/medium/compare/go-pure/domain"
	"net/http"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var IssueRepository domain.IssueRepository

func issueHandler(w http.ResponseWriter, r *http.Request) {
	id, hasId, err := parseIdFromUrl(r)
	if err != nil {
		handleError(err, w)
		return
	}
	if hasId {
		switch r.Method {
		case http.MethodGet:
			err = getSingle(w, id)
		case http.MethodPut:
			err = update(w, r, id)
		case http.MethodPatch:
			err = partialUpdate(w, r, id)
		case http.MethodDelete:
			err = delete(w, id)
		default:
			err = methodNotAllowed
		}
	} else if r.Method == http.MethodGet {
		err = getAll(w)
	} else if r.Method == http.MethodPost {
		err = create(w, r)
	}
	if err != nil {
		handleError(err, w)
	}
}

func parseIdFromUrl(r *http.Request) (uuid.UUID, bool, error) {
	p := strings.Split(r.URL.Path, "/")

	if len(p) < 3 {
		return uuid.UUID{}, false, nil
	}

	if len(p[2]) == 0 {
		return uuid.UUID{}, false, nil
	}

	id, err := uuid.Parse(p[2])
	if err != nil {
		return uuid.UUID{}, false, err

	}
	return id, true, nil
}

func getSingle(w http.ResponseWriter, id uuid.UUID) error {
	issue, err := IssueRepository.Find(id)
	if err != nil {
		return err
	}
	js, err := json.Marshal(issueToIssueDto(issue))
	if err != nil {
		return err
	}
	log.Debug("Issue %v returned", issue)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return nil
}

func getAll(w http.ResponseWriter) error {
	issues, err := IssueRepository.FindAll()
	if err != nil {
		return err
	}
	var issueBodies []issueDto
	for _, issue := range issues {
		issueBodies = append(issueBodies, issueToIssueDto(issue))
	}

	js, err := json.Marshal(issueBodies)
	if err != nil {
		return err
	}
	log.Debugf("Issues %v returned", issues)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return nil
}

func create(w http.ResponseWriter, r *http.Request) error {
	postBody, err := parseBody(r)
	if err != nil {
		return err
	}
	newIssue := postBody.toIssue()
	err = IssueRepository.Insert(newIssue)
	if err != nil {
		return err
	}
	log.Infof("Create issue %v", newIssue)
	return nil
}

func update(w http.ResponseWriter, r *http.Request, id uuid.UUID) error {
	postBody, err := parseBody(r)
	if err != nil {
		return err
	}
	newIssue := postBody.toIssue()
	err = IssueRepository.Update(newIssue)
	if err != nil {
		return err
	}
	log.Info("Update issue %v", newIssue)
	return nil
}

func partialUpdate(w http.ResponseWriter, r *http.Request, id uuid.UUID) error {
	postBody, err := parseBody(r)
	if err != nil {
		return err
	}

	oldIssue, err := IssueRepository.Find(id)
	if err != nil {
		return err
	}
	oldIssue = oldIssue.Update(postBody.Name, postBody.Description)

	err = IssueRepository.Update(oldIssue)
	if err != nil {
		return err
	}
	log.Info("Update issue %v", oldIssue)
	return nil
}

func delete(w http.ResponseWriter, id uuid.UUID) error {
	err := IssueRepository.Delete(id)
	if err != nil {
		return err
	}
	log.Info("Delete issue %v", id)
	return nil
}
