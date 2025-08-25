package arkruntime

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/utils"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

// CreateResponses Creates a model response.
func (c *Client) CreateResponses(ctx context.Context, body *responses.ResponsesRequest, opts ...requestOption) (res *responses.ResponseObject, err error) {
	path := "/responses"
	opts = append(opts, withBody(body))
	res = &responses.ResponseObject{}
	err = c.Do(ctx, http.MethodPost, c.fullURL(path), resourceTypeEndpoint, body.Model, res, opts...)
	return
}

// CreateResponsesStream Creates a model response.
func (c *Client) CreateResponsesStream(ctx context.Context, body *responses.ResponsesRequest, opts ...requestOption) (stream *utils.ResponsesStreamReader, err error) {
	body.Stream = volcengine.Bool(true)
	opts = append(opts, withBody(body))
	path := "/responses"
	return c.ResponsesRequestStreamDo(ctx, http.MethodPost, c.fullURL(path), resourceTypeEndpoint, body.Model, opts...)
}

// GetResponses Retrieves a model response with the given ID.
func (c *Client) GetResponses(ctx context.Context, responseID string, query *responses.GetResponseRequest, opts ...requestOption) (res *responses.ResponseObject, err error) {
	opts = append(opts, withBody(query))
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("/responses/%s", responseID)
	res = &responses.ResponseObject{}
	// NOTE: do not support ak/sk auth for now
	err = c.Do(ctx, http.MethodGet, c.fullURL(path), "", "", res, opts...)
	return
}

// DeleteResponse Deletes a model response with the given ID.
func (c *Client) DeleteResponse(ctx context.Context, responseID string, opts ...requestOption) (err error) {
	opts = append(opts, WithCustomHeader("Accept", ""))
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("/responses/%s", responseID)
	// NOTE: do not support ak/sk auth for now
	err = c.Do(ctx, http.MethodDelete, c.fullURL(path), "", "", nil, opts...)
	return
}

// ListResponseInputItems Returns a list of input items for a given response.
func (c *Client) ListResponseInputItems(ctx context.Context, responseID string, query *responses.ListInputItemsRequest, opts ...requestOption) (res *responses.ListInputItemsResponse, err error) {
	opts = append(opts, withBody(query))
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("/responses/%s/input_items", responseID)
	res = &responses.ListInputItemsResponse{}
	// NOTE: do not support ak/sk auth for now
	err = c.Do(ctx, http.MethodGet, c.fullURL(path), "", "", res, opts...)
	return
}
