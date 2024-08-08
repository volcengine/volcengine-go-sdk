package arkruntime

import (
	"context"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

const tokenizationSuffix = "/tokenization"

func (c *Client) CreateTokenization(
	ctx context.Context,
	conv model.TokenizationRequestConverter,
	setters ...requestOption,
) (res model.TokenizationResponse, err error) {
	baseReq := conv.Convert()

	requestOptions := append(setters, withBody(baseReq))
	err = c.Do(ctx, http.MethodPost, c.fullURL(tokenizationSuffix), resourceTypeEndpoint, baseReq.Model, &res, requestOptions...)
	if err != nil {
		return
	}
	return
}
