package arkruntime

import (
	"context"
	"fmt"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

const contentGenerationTasksSuffix = "/contents/generations/tasks"

func (c *Client) CreateContentGenerationTask(
	ctx context.Context,
	request model.CreateContentGenerationTaskRequest,
	setters ...requestOption,
) (response model.CreateContentGenerationTaskResponse, err error) {
	if !c.isApiKeyAuthentication() {
		err = model.ErrAKSKNotSupported
		return
	}

	requestOptions := append(setters, withBody(request))
	err = c.Do(ctx, http.MethodPost, c.fullURL(contentGenerationTasksSuffix), resourceTypeEndpoint, request.Model, &response, requestOptions...)
	if err != nil {
		return
	}
	return
}

func (c *Client) GetContentGenerationTask(
	ctx context.Context,
	request model.GetContentGenerationTaskRequest,
	setters ...requestOption,
) (response model.GetContentGenerationTaskResponse, err error) {
	if !c.isApiKeyAuthentication() {
		err = model.ErrAKSKNotSupported
		return
	}

	url := fmt.Sprintf("%s/%s", c.fullURL(contentGenerationTasksSuffix), request.ID)

	err = c.Do(ctx, http.MethodGet, url, resourceTypeEndpoint, "", &response, setters...)
	if err != nil {
		return
	}
	return
}

func (c *Client) DeleteContentGenerationTask(
	ctx context.Context,
	request model.DeleteContentGenerationTaskRequest,
	setters ...requestOption,
) (err error) {
	if !c.isApiKeyAuthentication() {
		err = model.ErrAKSKNotSupported
		return
	}

	url := fmt.Sprintf("%s/%s", c.fullURL(contentGenerationTasksSuffix), request.ID)

	err = c.Do(ctx, http.MethodDelete, url, resourceTypeEndpoint, "", nil, setters...)
	return err
}
