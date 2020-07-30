package rest

import (
	"lizzy/medium/compare/go-gin/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func NewEngine(issueController IssueController) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: log.New().Writer()}), gin.Recovery(), errorHandler)
	engine.GET("/issue/", issueController.readAll)
	engine.GET("/issue/:id/", issueController.readSingle)
	engine.POST("/issue/", issueController.create)
	engine.PUT("/issue/:id/", issueController.update)
	engine.PATCH("/issue/:id/", issueController.partialUpdate)
	engine.DELETE("/issue/:id/", issueController.delete)
	return engine
}

func errorHandler(c *gin.Context) {
	c.Next()
	error := c.Errors.Last()
	switch {
	case error.Err == domain.IssueNotFoundError:
		log.Warnf("Error while handling request: %v", error)
		c.JSON(http.StatusNotFound, error)
	case error.Type == gin.ErrorTypePublic:
		log.Warnf("Error while handling request: %v", error)
		c.JSON(http.StatusBadRequest, error)
	case error == nil:
	default:
		log.Warnf("Error while handling request: %v", error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not process request"})
	}
}
