package main

import (
	"github.com/goyek/goyek/v2"

	"demo/task"
)

// All runs the build workflow.
var All = goyek.Define(goyek.Task{
	Name:  "all",
	Usage: "build workflow",
	Deps: goyek.Deps{
		task.Fmt,
		Test,
	},
})
