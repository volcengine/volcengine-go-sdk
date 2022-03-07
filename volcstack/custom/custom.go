package custom

import (
	"context"
	"net/http"
)

type ExtendHttpRequest func(ctx context.Context, request *http.Request)

type ExtraHttpParameters func(ctx context.Context) map[string]string
