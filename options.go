package api

import (
	"fmt"
	"strings"
)

type Options struct {
	timeout   uint // milliseconds
	retries   uint // number of retries to attempt
}
type OptFunc func(*Options)

func (o *Options) string() string {
	rv := []string{}
	if o.timeout > 0   { rv = append(rv, fmt.Sprintf("timeout: %d", o.timeout)) }
	if o.retries > 0   { rv = append(rv, fmt.Sprintf("retries: %d", o.retries)) }
	return strings.Join(rv, ",")
}

// Option: Set Timeout in milliseconds. Any API calls which take longer than the
// specified timeout will be cancelled (and return a timeout error)
func Timeout(maxMillisecondsElapsed uint) OptFunc {
	return func(o *Options) {
		o.timeout = maxMillisecondsElapsed
	}
}

// Option: Set maximum allowed number of Retries. If a GET should happen to fail
// with a transient error, it will be retried the number of times specified
//
// The maximum number of retries permitted is 5
func RetriesPermitted(numRetries uint) OptFunc {
	return func(o *Options) {
		if numRetries > 5 {
			numRetries = 5
		}
		o.retries = numRetries
	}
}

// Set API options. Can be set on the *APIData object (via api.Options.Set()),
// in which case they provide the default value for each subsequent call -- or
// they can be set on individual *Requests (via req.Options.Set())
//
// Available options: RateLimit(n), Timeout(n)
func (o *Options) Set(opts ...OptFunc) {
	for _, fn := range opts {
		fn(o)
	}
}