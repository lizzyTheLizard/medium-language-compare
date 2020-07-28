package main

import (
	"lizzy/medium/compare/go-gin/persistence"
	"lizzy/medium/compare/go-gin/rest"
)

func main() {
	//Setup persistance layer
	persistence.Connect()
	defer persistence.Close()

	//Setup rest layer
	rest.IssueRepository = persistence.NewIssueRepository()
	rest.Start()
}
