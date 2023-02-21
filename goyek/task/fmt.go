package task

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

// Fmt runs go fmt.
var Fmt = goyek.Define(goyek.Task{
	Name:  "fmt",
	Usage: "go fmt",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go fmt ./...")
	},
})
