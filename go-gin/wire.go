//+build wireinject

package main

import (
	"lizzy/medium/compare/go-gin/persistence"
	"lizzy/medium/compare/go-gin/rest"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeEngine() (*gin.Engine, func()) {
	wire.Build(persistence.NewDB,
		persistence.NewIssueRepository,
		rest.NewIssueController,
		rest.NewEngine)
	return nil, func() {}
}
