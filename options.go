package api

type Options struct {
	rateLimit uint // rate: calls per minute
	timeout   uint // milliseconds
	retries   uint // number of retries to attempt
}
type OptFunc func(*Options)

// Option: Set Rate Limit to specified number of calls per minute
//
// TODO: Rate limiting currently not implemented; added to support client calls
func RateLimit(callsPerMinute uint) OptFunc {
	return func(o *Options) {
		o.rateLimit = callsPerMinute
	}
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