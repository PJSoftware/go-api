package api_test

import (
	"fmt"
	"testing"

	goapi "github.com/pjsoftware/go-api"
)

func TestNewAPIIdent(t *testing.T) {
	url := "http://test.io"

	api := goapi.New(url)
	ver := api.Version()

	t.Run("Ident", func(t *testing.T) {
		exp := fmt.Sprintf("go-api %s: %s", ver, url)
		got := api.Ident()
		if got != exp {
			t.Errorf("ident 1: got '%s' but expected '%s'", got, exp)
		}
	})

}
