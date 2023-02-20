package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

// Test runs go test.
var Test = goyek.Define(goyek.Task{
	Name:  "test",
	Usage: "go test",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go test -race -covermode=atomic -coverprofile=coverage.out -coverpkg=./... ./...")
		cmd.Exec(a, "go tool cover -html=coverage.out -o coverage.html")
	},
})
