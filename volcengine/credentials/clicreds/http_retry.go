package clicreds

import (
	"context"
	"errors"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type retryOptions struct {
	maxAttempts int
	baseDelay   time.Duration
	maxDelay    time.Duration
}

var (
	retryRandMu sync.Mutex
	retryRand   = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func doWithRetry(ctx context.Context, opts retryOptions, fn func() error) error {
	if opts.maxAttempts <= 0 {
		opts.maxAttempts = 1
	}
	if opts.baseDelay <= 0 {
		opts.baseDelay = 200 * time.Millisecond
	}
	if opts.maxDelay <= 0 {
		opts.maxDelay = 2 * time.Second
	}

	var lastErr error
	for attempt := 1; attempt <= opts.maxAttempts; attempt++ {
		if ctx != nil && ctx.Err() != nil {
			return ctx.Err()
		}

		lastErr = fn()
		if lastErr == nil {
			return nil
		}

		if attempt == opts.maxAttempts || !shouldRetryError(lastErr) {
			return lastErr
		}

		delay := computeBackoff(opts, attempt)
		if err := sleepWithContext(ctx, delay); err != nil {
			return err
		}
	}
	return lastErr
}

func shouldRetryError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}

	var oauthErr *OAuthAPIError
	if errors.As(err, &oauthErr) {
		return isRetryableHTTPStatus(oauthErr.StatusCode)
	}

	var portalErr *PortalAPIError
	if errors.As(err, &portalErr) {
		return isRetryableHTTPStatus(portalErr.StatusCode)
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		return true
	}

	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		if errors.Is(urlErr, context.Canceled) || errors.Is(urlErr, context.DeadlineExceeded) {
			return false
		}
		return true
	}

	if errors.Is(err, io.EOF) {
		return true
	}

	return false
}

func isRetryableHTTPStatus(code int) bool {
	return code == http.StatusTooManyRequests || code == http.StatusRequestTimeout || code/100 == 5
}

func computeBackoff(opts retryOptions, attempt int) time.Duration {
	// attempt is 1-based; backoff after the first failure starts at baseDelay.
	exp := attempt - 1
	delay := opts.baseDelay
	for i := 0; i < exp; i++ {
		delay *= 2
		if delay >= opts.maxDelay {
			delay = opts.maxDelay
			break
		}
	}

	// add jitter up to 100ms to reduce thundering herd
	retryRandMu.Lock()
	jitter := time.Duration(retryRand.Int63n(int64(100 * time.Millisecond)))
	retryRandMu.Unlock()

	return delay + jitter
}

func sleepWithContext(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return nil
	}
	if ctx == nil {
		time.Sleep(d)
		return nil
	}
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
