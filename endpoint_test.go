package api_test

import (
	"testing"

	goapi "github.com/pjsoftware/go-api"
)

func TestNewEndpoint(t *testing.T) {
	root := "http://test.io"
	epURL := "endpoint"
	exp := root + "/" + epURL

	t.Run("EndpointURL No Slash", func(t *testing.T) {
		api := goapi.New(root)
		ep := api.NewEndpoint(epURL)
		got := ep.URL()
		if got != exp {
			t.Errorf("endpoint url: got '%s' but expected '%s'", got, exp)
		}
	})

	t.Run("EndpointURL One Slash L", func(t *testing.T) {
		api := goapi.New(root + "/")
		ep := api.NewEndpoint(epURL)
		got := ep.URL()
		if got != exp {
			t.Errorf("endpoint url: got '%s' but expected '%s'", got, exp)
		}
	})

	t.Run("EndpointURL One Slash R", func(t *testing.T) {
		api := goapi.New(root)
		ep := api.NewEndpoint("/" + epURL)
		got := ep.URL()
		if got != exp {
			t.Errorf("endpoint url: got '%s' but expected '%s'", got, exp)
		}
	})

	t.Run("EndpointURL Two Slash", func(t *testing.T) {
		api := goapi.New(root + "/")
		ep := api.NewEndpoint("/" + epURL)
		got := ep.URL()
		if got != exp {
			t.Errorf("endpoint url: got '%s' but expected '%s'", got, exp)
		}
	})

}
