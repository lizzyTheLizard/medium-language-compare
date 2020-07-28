package rest

import (
	"errors"
	"lizzy/medium/compare/go-gin/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type issueDto struct {
	Id          uuid.UUID
	Name        string
	Description string
}

func parseBody(c *gin.Context) (issueDto, error) {
	var issueDto issueDto
	if err := c.ShouldBindJSON(&issueDto); err != nil {
		return issueDto, &gin.Error{
			Err:  errors.New("Cannot parse body: " + err.Error()),
			Type: gin.ErrorTypePublic}
	}
	return issueDto, nil
}

func issueToIssueDto(issue domain.Issue) issueDto {
	return issueDto{issue.GetId(), issue.GetName(), issue.GetDescription()}
}

func (i issueDto) toIssue() domain.Issue {
	return domain.NewIssue(i.Id, i.Name, i.Description)
}
