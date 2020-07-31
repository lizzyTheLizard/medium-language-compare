package domain

import (
	"errors"

	"github.com/google/uuid"
)

var IssueNotFoundError = errors.New("Issue not found")

type IssueRepository interface {
	Find(uuid.UUID) (Issue, error)
	FindAll() (issues []Issue, err error)
	Update(Issue) error
	Insert(Issue) error
	Delete(uuid.UUID) error
}
