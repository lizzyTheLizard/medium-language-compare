package persistence

import (
	"lizzy/medium/compare/go-pure/domain"

	"github.com/google/uuid"
)

type issueModel struct {
	id          uuid.UUID
	name        string
	description string
}

func (i issueModel) toIssue() domain.Issue {
	return domain.NewIssue(i.id, i.name, i.description)
}
