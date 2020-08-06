//+build wireinject

package main

import (
	"lizzy/medium/compare/persistence"
	"lizzy/medium/compare/rest"

	"github.com/google/wire"
)

func InitializeEngine() (rest.Engine, func()) {
	wire.Build(persistence.Connect,
		persistence.NewIssueRepository,
		rest.NewIssueController,
		rest.NewEngine)
	return rest.Engine{}, func() {}
}
