package rest

import (
	"lizzy/medium/compare/go-pure/domain"

	"github.com/google/uuid"
)

type issueDto struct {
	Id          uuid.UUID
	Name        string
	Description string
}

func issueToIssueDto(issue domain.Issue) issueDto {
	return issueDto{issue.GetId(), issue.GetName(), issue.GetDescription()}
}

func (i issueDto) toIssue() domain.Issue {
	return domain.NewIssue(i.Id, i.Name, i.Description)
}
