package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Issue struct {
	id          uuid.UUID
	name        string
	description string
}

func (i Issue) GetId() uuid.UUID {
	return i.id
}

func (i Issue) GetName() string {
	return i.name
}

func (i Issue) GetDescription() string {
	return i.description
}

func (i Issue) Update(newName string, newDescription string) Issue {
	if newName != "" {
		i.name = newName
	}
	if newDescription != "" {
		i.name = newDescription
	}
	return i
}

func NewIssue(id uuid.UUID, name string, description string) Issue {
	return Issue{id, name, description}
}

type IssueRepository interface {
	Find(uuid.UUID) (Issue, error)
	FindAll() (issues []Issue, err error)
	Update(Issue) error
	Insert(Issue) error
	Delete(uuid.UUID) error
}

var IssueNotFoundError = errors.New("Issue not found")
