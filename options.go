package api

type Options struct {
	rateLimit int // rate: calls per minute
	timeout   int // milliseconds
}
type OptFunc func(*Options)

// Option: Set Rate Limit to specified number of calls per minute
//
// TODO: Rate limiting currently not implemented; added to support client calls
func RateLimit(callsPerMinute int) OptFunc {
	return func(o *Options) {
		o.rateLimit = callsPerMinute
	}
}

// Option: Set Timeout in milliseconds. Any API calls which take longer than the
// specified timeout will be cancelled (and return a timeout error)
//
// TODO: Timeout currently not implemented
func Timeout(maxMillisecondsElapsed int) OptFunc {
	return func(o *Options) {
		o.timeout = maxMillisecondsElapsed
	}
}

// Set API options. Can be set on the *APIData object (via api.Options.Set()),
// in which case they provide the default value for each subsequent call -- or
// they can be set on individual *Requests (via req.Options.Set())
func (o *Options) Set(opts ...OptFunc) {
	for _, fn := range opts {
		fn(o)
	}
}