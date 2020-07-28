package main

import (
	"lizzy/medium/compare/go-pure/persistence"
	"lizzy/medium/compare/go-pure/rest"
)

func main() {
	//Setup persistance layer
	persistence.Connect()
	defer persistence.Close()

	//Setup rest layer
	rest.IssueRepository = persistence.NewIssueRepository()
	rest.Start()
}
