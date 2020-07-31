//+build wireinject

package main

import (
	"lizzy/medium/compare/go-pure/persistence"
	"lizzy/medium/compare/go-pure/rest"

	"github.com/google/wire"
)

func InitializeEngine() (rest.Engine, func()) {
	wire.Build(persistence.NewDB,
		persistence.NewIssueRepository,
		rest.NewIssueController,
		rest.NewEngine)
	return rest.Engine{}, func() {}
}
