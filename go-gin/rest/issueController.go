package rest

import (
	"errors"
	"lizzy/medium/compare/go-gin/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type IssueController struct {
	repository domain.IssueRepository
}

func (i IssueController) readSingle(c *gin.Context) {
	issue, err := i.getIssue(c)
	if err != nil {
		c.Error(err)
		return
	}
	log.Debug("Issue %v returned", issue)
	i.writeAnswer(c, issue)
}

func (i IssueController) readAll(c *gin.Context) {
	issues, err := i.repository.FindAll()
	if err != nil {
		c.Error(err)
		return
	}
	log.Debugf("Issues %v returned", issues)
	i.writeAnswers(c, issues)
}

func (i IssueController) create(c *gin.Context) {
	postBody, err := i.parseBody(c)
	if err != nil {
		c.Error(err)
		return
	}
	newIssue := postBody.toIssue()
	err = i.repository.Insert(newIssue)
	if err != nil {
		c.Error(err)
		return
	}
	log.Infof("Create issue %v", newIssue)
	i.writeAnswer(c, newIssue)
}

func (i IssueController) update(c *gin.Context) {
	postBody, err := i.parseBody(c)
	if err != nil {
		c.Error(err)
		return
	}
	newIssue := postBody.toIssue()
	err = i.repository.Update(newIssue)
	if err != nil {
		c.Error(err)
		return
	}
	log.Infof("Update issue %v", newIssue)
	i.writeAnswer(c, newIssue)
}

func (i IssueController) partialUpdate(c *gin.Context) {
	postBody, err := i.parseBody(c)
	if err != nil {
		c.Error(err)
		return
	}
	oldIssue, err := i.getIssue(c)
	if err != nil {
		c.Error(err)
		return
	}
	newIssue := oldIssue.Update(postBody.Name, postBody.Description)
	err = i.repository.Update(newIssue)
	if err != nil {
		c.Error(err)
		return
	}
	log.Infof("Update issue %v", newIssue)
	i.writeAnswer(c, newIssue)
}

func (i IssueController) delete(c *gin.Context) {
	id, err := i.parseId(c)
	if err != nil {
		c.Error(err)
		return
	}
	err = i.repository.Delete(id)
	if err != nil {
		c.Error(err)
		return
	}
	log.Infof("Delete issue %v", id)
	c.Status(http.StatusOK)
}

func (i IssueController) getIssue(c *gin.Context) (domain.Issue, error) {
	id, err := i.parseId(c)
	if err != nil {
		return domain.Issue{}, err
	}
	return i.repository.Find(id)
}

func (i IssueController) parseId(c *gin.Context) (uuid.UUID, error) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return id, &gin.Error{
			Err:  errors.New("Cannot parse id: " + err.Error()),
			Type: gin.ErrorTypePublic}
	}
	return id, nil
}

func (i IssueController) writeAnswer(c *gin.Context, issue domain.Issue) {
	issueDto := issueToIssueDto(issue)
	c.JSON(http.StatusOK, issueDto)
}

func (i IssueController) writeAnswers(c *gin.Context, issues []domain.Issue) {
	var issueDtos []issueDto
	for _, issue := range issues {
		issueDtos = append(issueDtos, issueToIssueDto(issue))
	}
	c.JSON(http.StatusOK, issueDtos)
}

func (i IssueController) parseBody(c *gin.Context) (issueDto, error) {
	var issueDto issueDto
	if err := c.ShouldBindJSON(&issueDto); err != nil {
		return issueDto, &gin.Error{
			Err:  errors.New("Cannot parse body: " + err.Error()),
			Type: gin.ErrorTypePublic}
	}
	return issueDto, nil
}

func NewIssueController(repository domain.IssueRepository) IssueController {
	return IssueController{repository}
}
