package api_test

import (
	"fmt"
	"testing"

	goapi "github.com/pjsoftware/go-api"
)

func TestNewEndpoint(t *testing.T) {
	root := "http://test.io"
	epURL := "/endpoint"

	api := goapi.New(root)
	ep := api.NewEndpoint(epURL)

	t.Run("EndpointURL", func(t *testing.T) {
		exp := fmt.Sprintf("%s%s", root, epURL)
		got := ep.URL()
		if got != exp {
			t.Errorf("endpoint url: got '%s' but expected '%s'", got, exp)
		}
	})

}
