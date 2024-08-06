package utils

import (
	"context"
	"math"
	"math/rand"
	"time"
)

type RetryPolicy struct {
	MaxAttempts    int
	InitialBackoff time.Duration
	MaxBackoff     time.Duration
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
		if numRetriesSincePushback == rp.MaxAttempts {
			break
		}
		nbRetries := numRetriesSincePushback + 1
		sleepSeconds := math.Min(rp.InitialBackoff.Seconds()*math.Pow(2.0, float64(nbRetries)), rp.MaxBackoff.Seconds())
		jitter := 1.0 - 0.25*rand.Float64()
		dur := time.Duration(sleepSeconds*jitter) * time.Second

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
