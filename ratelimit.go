package api

import (
	"time"
)

// To rate-limit a particular endpoint, we need to know how many calls are
// allowed per unit of time: N calls per T time; eg, 100 calls per minute.
//
// If we pre-fill a queue (a bucket) with N tokens, we can then do this:
//
// - If we need to make a call, check the bucket
// - If there is no token, add the requested call to a queue.
// - If there is a token available:
//   - Remove the token.
//   - Allow the call.
//   - Start a timer; after T time, add a new token to the bucket.
//
// Note that many implementations talk about adding N tokens per T to the bucket
// at a fixed rate -- but if we fire off N tokens in the first second (in a burst)
// and then add a new token to the bucket after N/T time, this would allow a new
// call to be sent (N+1 calls) within the T window of time.
//
// This is not what we want.
//
// If we define our bucket as a (buffered) channel of size n, then we can write
// a value to the bucket every time we request a token. If our bucket is full,
// this will automatically block until a value is removed from the channel. The
// goroutine which we fire off then only needs to read from the channel after T
// time has passed; this should automatically remove our block!

// rateLimiter allows N calls per T time
type rateLimiter struct {
	n int
	t time.Duration
	bucket chan bool
}

func newRateLimiter(n int, t time.Duration) *rateLimiter {
	rl := rateLimiter{
		n: n,
		t: t,
		bucket: make(chan bool, n),
	}
	return &rl
}

func (rl *rateLimiter) requestToken() {
	rl.bucket <- true // block if bucket is full
	go func() {
		time.Sleep(rl.t)
		<- rl.bucket
	}()
}
