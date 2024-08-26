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

	t.Run("EndpointURL Drop Root Prefix", func(t *testing.T) {
		api := goapi.New(root)
		ep := api.NewEndpoint(root + "/" + epURL)
		got := ep.URL()
		if got != exp {
			t.Errorf("endpoint url: got '%s' but expected '%s'", got, exp)
		}
	})
}

func TestEndpointIsSingleton(t *testing.T) {
	root := "http://test.io"
	epURL1 := "endpoint1"
	epURL2 := "endpoint2"

	t.Run("NewEndpoint returns singleton", func(t *testing.T) {
		api := goapi.New(root)
		ep1 := api.NewEndpoint(epURL1)
		ep2 := api.NewEndpoint(epURL1)

		if ep1 != ep2 {
			t.Errorf("two endpoints should match, but do not: %v / %v", ep1, ep2)
		}
	})

	t.Run("Endpoint caching pre-processes URL", func(t *testing.T) {
		api := goapi.New(root)
		ep1 := api.NewEndpoint(epURL1)
		ep2 := api.NewEndpoint(root + "/" + epURL1)

		if ep1 != ep2 {
			t.Errorf("endpoint caching fails to detect equivalence: %v / %v", ep1, ep2)
		}
	})

	t.Run("Different endpoint per url", func(t *testing.T) {
		api := goapi.New(root)
		ep1 := api.NewEndpoint(epURL1)
		ep2 := api.NewEndpoint(epURL2)

		if ep1 == ep2 {
			t.Errorf("does not return unique EP: %v / %v", ep1, ep2)
		}
	})}