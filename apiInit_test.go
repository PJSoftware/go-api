package api_test

import (
	"fmt"
	"testing"

	goapi "github.com/pjsoftware/go-api"
)

func TestNewAPIIdent(t *testing.T) {
	url := "http://test.io"
	name := "TestAPI"

	api := goapi.New(url)
	ver := api.Version()

	t.Run("IdentNoName", func(t *testing.T) {
		exp := fmt.Sprintf("go-api v%s: %s (Auth: None)", ver, url)
		got := api.Ident()
		if got != exp {
			t.Errorf("ident 1: got '%s' but expected '%s'", got, exp)
		}
	})

	t.Run("IdentWithName", func(t *testing.T) {
		api.SetName(name)
		exp := fmt.Sprintf("%s via go-api v%s: %s (Auth: None)", name, ver, url)
		got := api.Ident()
		if got != exp {
			t.Errorf("ident 2: got '%s' but expected '%s'", got, exp)
		}
	})

}
