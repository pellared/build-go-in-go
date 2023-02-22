//go:build mage
// +build mage

package main

import (
	"errors"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	// mage:import
	"demo/common"
)

// All runs build pipeline.
func All() {
	mg.SerialDeps(common.Fmt, Test)
}

// Test runs go test with race detector and code covarage.
func Test() error {
	err := sh.Run("go", "test", "-race", "-covermode=atomic", "-coverprofile=coverage.out", "./...")
	return errors.Join(err, sh.Run("go", "tool", "cover", "-html=coverage.out", "-o", "coverage.html"))
}
