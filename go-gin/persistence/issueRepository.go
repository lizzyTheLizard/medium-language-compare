package persistence

import (
	"lizzy/medium/compare/go-gin/domain"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type issueRepository struct {
	db *gorm.DB
}

func NewIssueRepository() domain.IssueRepository {
	return issueRepository{db}
}

func (i issueRepository) Find(id uuid.UUID) (domain.Issue, error) {
	var issueModel issueModel
	result := i.db.Where("id=?", id).FirstOrInit(&issueModel)
	if (issueModel.Id == uuid.UUID{}) {
		return domain.Issue{}, domain.IssueNotFoundError
	}
	return issueModel.toIssue(), result.Error
}

func (i issueRepository) FindAll() ([]domain.Issue, error) {
	var issueModels []issueModel
	var issues []domain.Issue
	result := i.db.Find(&issueModels)
	for _, issueModel := range issueModels {
		issues = append(issues, issueModel.toIssue())
	}
	return issues, result.Error
}

func (i issueRepository) Update(issue domain.Issue) error {
	result := i.db.Save(issueToIssueModel(issue))
	return result.Error
}

func (i issueRepository) Insert(issue domain.Issue) error {
	result := i.db.Create(issueToIssueModel(issue))
	return result.Error
}

func (i issueRepository) Delete(id uuid.UUID) error {
	result := i.db.Where("id=?", id).Delete(&issueModel{})
	return result.Error
}
