package persistence

import (
	"lizzy/medium/compare/go-gin/domain"

	"github.com/google/uuid"
)

type issueModel struct {
	Id          uuid.UUID `gorm:"primary_key"`
	Name        string
	Description string
}

func (issueModel) TableName() string {
	return "issue"
}

func (i issueModel) toIssue() domain.Issue {
	return domain.NewIssue(i.Id, i.Name, i.Description)
}

func issueToIssueModel(issue domain.Issue) *issueModel {
	issueModel := issueModel{issue.GetId(), issue.GetName(), issue.GetDescription()}
	return &issueModel
}
