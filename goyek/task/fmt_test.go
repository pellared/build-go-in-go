package task

import (
	"testing"

	"github.com/goyek/goyek/v2"
)

func TestFmt(t *testing.T) {
	r := goyek.NewRunner(func(a *goyek.A) {
		Fmt.Action()(a)
	})

	if got, want := r(goyek.Input{}), goyek.StatusPassed; got.Status != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
