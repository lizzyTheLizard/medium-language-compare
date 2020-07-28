package rest

import (
	"encoding/json"
	"lizzy/medium/compare/go-pure/domain"
	"net/http"

	"github.com/google/uuid"
)

type issueDto struct {
	Id          uuid.UUID
	Name        string
	Description string
}

func parseBody(r *http.Request) (issueDto, error) {
	var body issueDto
	err := json.NewDecoder(r.Body).Decode(&body)
	return body, err
}

func issueToIssueDto(issue domain.Issue) issueDto {
	return issueDto{issue.GetId(), issue.GetName(), issue.GetDescription()}
}

func (i issueDto) toIssue() domain.Issue {
	return domain.NewIssue(i.Id, i.Name, i.Description)
}
