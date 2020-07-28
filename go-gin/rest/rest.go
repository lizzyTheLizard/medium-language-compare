package rest

import (
	"lizzy/medium/compare/go-gin/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Start() {
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: log.New().Writer()}), gin.Recovery(), errorHandler)
	r.GET("/issue/", readAll)
	r.GET("/issue/:id/", readSingle)
	r.POST("/issue/", create)
	r.PUT("/issue/:id/", update)
	r.PATCH("/issue/:id/", partialUpdate)
	r.DELETE("/issue/:id/", delete)
	r.Run()
}

func errorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) == 0 {
		return
	}
	error := c.Errors.Last()
	switch {
	case error.Err == domain.IssueNotFoundError:
		log.Warnf("Error while handling request: %v", error)
		c.JSON(http.StatusNotFound, error)
	case error.Type == gin.ErrorTypePublic:
		log.Warnf("Error while handling request: %v", error)
		c.JSON(http.StatusBadRequest, error)
	default:
		log.Warnf("Error while handling request: %v", error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not process request"})
	}
}
