package api

import (
	"testing"
)

func TestAPIOptions(t *testing.T) {
	var (
		url = "http://test.io"
		epURL = "test"
		rateLim = 100
		timeout = 500 
	)

	t.Run("DefaultOptionValues", func(t *testing.T) {
		testAPI := New(url)

		// Default option zero-values
		if testAPI.Options.rateLimit != 0 {
			t.Errorf("api rateLimit default: got '%d' but expected '0'", testAPI.Options.rateLimit)
		}
		if testAPI.Options.timeout != 0 {
			t.Errorf("api timeout default: got '%d' but expected '0'", testAPI.Options.timeout)
		}

		// Confirm that setting API options works
		testAPI.Options.Set(RateLimit(rateLim), Timeout(timeout))
		if testAPI.Options.rateLimit != rateLim {
			t.Errorf("api rateLimit new value: got '%d' but expected '%d'", testAPI.Options.rateLimit, rateLim)
		}
		if testAPI.Options.timeout != timeout {
			t.Errorf("api timeout new value: got '%d' but expected '%d'", testAPI.Options.timeout, timeout)
		}

		// Confirm that api options carry over to request
		ep := testAPI.NewEndpoint(epURL)
		req := ep.NewRequest()
		if req.Options.rateLimit != rateLim {
			t.Errorf("req rateLimit default: got '%d' but expected '%d'", req.Options.rateLimit, rateLim)
		}
		if req.Options.timeout != timeout {
			t.Errorf("req timeout default: got '%d' but expected '%d'", req.Options.timeout, timeout)
		}

		// confirm that changing request options works
		nv := 1000
		req.Options.Set(RateLimit(nv), Timeout(nv))
		if req.Options.rateLimit != nv {
			t.Errorf("req rateLimit new value: got '%d' but expected '%d'", req.Options.rateLimit, nv)
		}
		if req.Options.timeout != nv {
			t.Errorf("req timeout new value: got '%d' but expected '%d'", req.Options.timeout, nv)
		}

		// confirm that changing request options did not change api defaults
		if testAPI.Options.rateLimit != rateLim {
			t.Errorf("req rateLimit new value: got '%d' but expected '%d'", testAPI.Options.rateLimit, rateLim)
		}
		if testAPI.Options.timeout != timeout {
			t.Errorf("req timeout new value: got '%d' but expected '%d'", testAPI.Options.timeout, timeout)
		}
	})

}
