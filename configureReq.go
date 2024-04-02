package api

type ReqOptions struct {
	rateLimit int // rate: calls per minute
}
type ReqOptFunc func(*Request)

// Req Configuration: Set Rate Limit to specified number of calls per minute
//
// TODO: Rate limiting currently not implemented; added to support client calls
func RateLimit(callsPerMinute int) ReqOptFunc {
	return func(r *Request) {
		r.options.rateLimit = callsPerMinute
	}
}

func (r *Request) Configure(opts ...ReqOptFunc) {
	for _, fn := range opts {
		fn(r)
	}
}