package rest

import (
	"encoding/json"
	"errors"
	"lizzy/medium/compare/go-pure/domain"
	"net/http"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type IssueController struct {
	repository domain.IssueRepository
}

func (i IssueController) readSingle(w http.ResponseWriter, r *http.Request) error {
	issue, err := i.getIssue(r)
	if err != nil {
		return err
	}
	log.Debug("Issue %v returned", issue)
	return i.writeAnswer(w, issue)
}

func (i IssueController) readAll(w http.ResponseWriter, r *http.Request) error {
	issues, err := i.repository.FindAll()
	if err != nil {
		return err
	}
	log.Debugf("Issues %v returned", issues)
	return i.writeAnswers(w, issues)
}

func (i IssueController) create(w http.ResponseWriter, r *http.Request) error {
	postBody, err := i.parseBody(r)
	if err != nil {
		return err
	}
	newIssue := postBody.toIssue()
	err = i.repository.Insert(newIssue)
	if err != nil {
		return err
	}
	log.Infof("Create issue %v", newIssue)
	return i.writeAnswer(w, newIssue)
}

func (i IssueController) update(w http.ResponseWriter, r *http.Request) error {
	postBody, err := i.parseBody(r)
	if err != nil {
		return err
	}
	newIssue := postBody.toIssue()
	err = i.repository.Update(newIssue)
	if err != nil {
		return err
	}
	log.Infof("Update issue %v", newIssue)
	return i.writeAnswer(w, newIssue)
}

func (i IssueController) partialUpdate(w http.ResponseWriter, r *http.Request) error {
	postBody, err := i.parseBody(r)
	if err != nil {
		return err
	}

	oldIssue, err := i.getIssue(r)
	if err != nil {
		return err
	}
	newIssue := oldIssue.Update(postBody.Name, postBody.Description)
	err = i.repository.Update(newIssue)
	if err != nil {
		return err
	}
	log.Infof("Update issue %v", newIssue)
	return i.writeAnswer(w, newIssue)
}

func (i IssueController) delete(w http.ResponseWriter, r *http.Request) error {
	id, err := i.parseId(r)
	if err != nil {
		return err
	}
	err = i.repository.Delete(id)
	if err != nil {
		return err
	}
	log.Infof("Delete issue %v", id)
	w.WriteHeader(http.StatusOK)
	return nil
}

func (i IssueController) getIssue(r *http.Request) (domain.Issue, error) {
	id, err := i.parseId(r)
	if err != nil {
		return domain.Issue{}, err
	}
	return i.repository.Find(id)
}

func (i IssueController) parseId(r *http.Request) (uuid.UUID, error) {
	p := strings.Split(r.URL.Path, "/")
	id, err := uuid.Parse(p[2])
	if err != nil {
		return id, errors.New("Cannot parse id: " + err.Error())

	}
	return id, nil
}

func (i IssueController) writeAnswer(w http.ResponseWriter, issue domain.Issue) error {
	issueDto := issueToIssueDto(issue)
	js, err := json.Marshal(issueDto)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return nil
}

func (i IssueController) writeAnswers(w http.ResponseWriter, issues []domain.Issue) error {
	var issueDtos []issueDto
	for _, issue := range issues {
		issueDtos = append(issueDtos, issueToIssueDto(issue))
	}
	js, err := json.Marshal(issueDtos)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return nil
}

func (i IssueController) parseBody(r *http.Request) (issueDto, error) {
	var body issueDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return body, errors.New("Cannot parse body: " + err.Error())
	}
	return body, nil
}

func NewIssueController(repository domain.IssueRepository) IssueController {
	return IssueController{repository}
}
