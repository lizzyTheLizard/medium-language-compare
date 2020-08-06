package persistence

import (
	"database/sql"
	"lizzy/medium/compare/domain"

	"github.com/google/uuid"
)

type issueRepository struct {
	db *sql.DB
}

func NewIssueRepository(db *sql.DB) domain.IssueRepository {
	return issueRepository{db}
}

func (i issueRepository) Find(id uuid.UUID) (domain.Issue, error) {
	rows, err := i.db.Query("SELECT name, description FROM issue WHERE id = $1", id)
	if err != nil {
		return domain.Issue{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return domain.Issue{}, domain.IssueNotFoundError
	}
	var name, description string
	err = rows.Scan(&name, &description)
	if err != nil {
		return domain.Issue{}, err
	}
	return domain.NewIssue(id, name, description), rows.Err()
}

func (i issueRepository) FindAll() ([]domain.Issue, error) {
	var issues []domain.Issue
	rows, err := i.db.Query("SELECT id, name, description FROM issue")
	if err != nil {
		return issues, err
	}
	defer rows.Close()

	for rows.Next() {
		var name, description string
		var id uuid.UUID
		err = rows.Scan(&id, &name, &description)
		if err != nil {
			return issues, err
		}
		issues = append(issues, domain.NewIssue(id, name, description))
	}
	err = rows.Err()
	return issues, err
}

func (i issueRepository) Update(issue domain.Issue) error {
	_, err := i.db.Exec("UPDATE issue SET name = $2, description = $3 WHERE id =$1", issue.GetId(), issue.GetName(), issue.GetDescription())
	return err
}

func (i issueRepository) Insert(issue domain.Issue) error {
	_, err := i.db.Exec("INSERT INTO issue(id, name, description) VALUES($1,$2,$3)", issue.GetId(), issue.GetName(), issue.GetDescription())
	return err
}

func (i issueRepository) Delete(id uuid.UUID) error {
	_, err := i.db.Exec("DELETE FROM issue WHERE id = $1", id)
	return err
}
