package utils

import (
	"context"
	"math"
	"math/rand"
	"time"
)

type RetryPolicy struct {
	// MaxAttempts is the maximum number of attempts, including the original RPC.
	//
	// This field is required and must be two or greater.
	MaxAttempts int

	// Exponential backoff parameters. The initial retry attempt will occur at
	// random(0, initialBackoff). In general, the nth attempt will occur at
	// random(0,
	//   min(initialBackoff*backoffMultiplier**(n-1), maxBackoff)).
	//
	// These fields are required and must be greater than zero.
	InitialBackoff    time.Duration
	MaxBackoff        time.Duration
	BackoffMultiplier float64
}

func Retry(ctx context.Context,
	rp RetryPolicy,
	isNeedRetry func() bool,
	doFunc func() error, overRetryLimitError error,
	isNeedRetryError func(error) bool,
) error {
	var err error
	for numRetriesSincePushback := 0; numRetriesSincePushback <= rp.MaxAttempts; numRetriesSincePushback++ {
		err = doFunc()

		// no error: just return on this try
		if err == nil {
			return nil
		}

		// not need retry: first time do return
		if isNeedRetry != nil && !isNeedRetry() {
			return err
		}

		// no need to retry error
		if !isNeedRetryError(err) {
			return err
		}

		// need retry
		fact := math.Pow(rp.BackoffMultiplier, float64(numRetriesSincePushback))
		cur := float64(rp.InitialBackoff) * fact
		if max := float64(rp.MaxBackoff); cur > max {
			cur = max
		}
		dur := time.Duration(rand.Int63n(int64(cur)))

		t := time.NewTimer(dur)
		select {
		case <-t.C: // continue next retry
		case <-ctx.Done(): // whole context finish
			t.Stop()
			return ctx.Err()
		}
	}

	// Note: if not set over retry limit error, return last meet error
	if overRetryLimitError == nil {
		return err
	}
	return overRetryLimitError
}
