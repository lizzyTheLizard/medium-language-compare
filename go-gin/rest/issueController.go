package rest

import (
	"errors"
	"lizzy/medium/compare/go-gin/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var IssueRepository domain.IssueRepository

func readSingle(c *gin.Context) {
	issue, err := getIssue(c)
	if err != nil {
		c.Error(err)
		return
	}
	log.Debug("Issue %v returned", issue)
	c.JSON(http.StatusOK, issueToIssueDto(issue))
}

func readAll(c *gin.Context) {
	issues, err := IssueRepository.FindAll()
	if err != nil {
		c.Error(err)
		return
	}
	log.Debugf("Issues %v returned", issues)
	var issueDtos []issueDto
	for _, issue := range issues {
		issueDtos = append(issueDtos, issueToIssueDto(issue))
	}
	c.JSON(http.StatusOK, issueDtos)
}

func create(c *gin.Context) {
	issue, err := parseBody(c)
	if err != nil {
		c.Error(err)
		return
	}

	err = IssueRepository.Insert(issue.toIssue())
	if err != nil {
		c.Error(err)
		return
	}
	log.Infof("Create issue %v", issue)
	c.JSON(http.StatusOK, issue)
}

func update(c *gin.Context) {
	issue, err := parseBody(c)
	if err != nil {
		c.Error(err)
		return
	}
	err = IssueRepository.Update(issue.toIssue())
	if err != nil {
		c.Error(err)
		return
	}
	log.Infof("Create issue %v", issue)
	c.JSON(http.StatusOK, issue)
}

func partialUpdate(c *gin.Context) {
	issue, err := parseBody(c)
	if err != nil {
		c.Error(err)
		return
	}
	oldIssue, err := getIssue(c)
	if err != nil {
		c.Error(err)
		return
	}
	oldIssue = oldIssue.Update(issue.Name, issue.Description)
	err = IssueRepository.Update(oldIssue)
	if err != nil {
		c.Error(err)
		return
	}
	log.Info("Update issue %v", issue)
	c.JSON(http.StatusOK, oldIssue)
}

func delete(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		c.Error(err)
		return
	}
	err = IssueRepository.Delete(id)
	if err != nil {
		c.Error(err)
		return
	}
	log.Info("Delete issue %v", id)
	c.Status(http.StatusOK)
}

func getIssue(c *gin.Context) (domain.Issue, error) {
	id, err := parseId(c)
	if err != nil {
		return domain.Issue{}, err
	}
	return IssueRepository.Find(id)
}

func parseId(c *gin.Context) (uuid.UUID, error) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return id, &gin.Error{
			Err:  errors.New("Cannot parse id: " + err.Error()),
			Type: gin.ErrorTypePublic}
	}
	return id, nil
}
