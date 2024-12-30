package arkruntime

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

const contentGenerationTasksSuffix = "/contents/generations/tasks"

func (c *Client) CreateContentGenerationTask(
	ctx context.Context,
	request model.CreateContentGenerationTaskRequest,
	setters ...requestOption,
) (response model.CreateContentGenerationTaskResponse, err error) {
	if !c.isAPIKeyAuthentication() {
		return response, model.ErrAKSKNotSupported
	}

	requestOptions := append(setters, withBody(request))
	err = c.Do(ctx, http.MethodPost, c.fullURL(contentGenerationTasksSuffix), resourceTypeEndpoint, request.Model, &response, requestOptions...)
	return
}

func (c *Client) GetContentGenerationTask(
	ctx context.Context,
	request model.GetContentGenerationTaskRequest,
	setters ...requestOption,
) (response model.GetContentGenerationTaskResponse, err error) {
	if !c.isAPIKeyAuthentication() {
		return response, model.ErrAKSKNotSupported
	}

	url := fmt.Sprintf("%s/%s", c.fullURL(contentGenerationTasksSuffix), request.ID)

	err = c.Do(ctx, http.MethodGet, url, resourceTypeEndpoint, "", &response, setters...)
	return
}

func (c *Client) DeleteContentGenerationTask(
	ctx context.Context,
	request model.DeleteContentGenerationTaskRequest,
	setters ...requestOption,
) (err error) {
	if !c.isAPIKeyAuthentication() {
		return model.ErrAKSKNotSupported
	}

	url := fmt.Sprintf("%s/%s", c.fullURL(contentGenerationTasksSuffix), request.ID)

	err = c.Do(ctx, http.MethodDelete, url, resourceTypeEndpoint, "", nil, setters...)
	return err
}

func (c *Client) ListContentGenerationTasks(
	ctx context.Context,
	request model.ListContentGenerationTasksRequest,
	setters ...requestOption,
) (response model.ListContentGenerationTasksResponse, err error) {
	if !c.isAPIKeyAuthentication() {
		return response, model.ErrAKSKNotSupported
	}

	values := url.Values{}
	if request.PageNum > 0 {
		values.Add("page_num", strconv.Itoa(request.PageNum))
	}
	if request.PageSize > 0 {
		values.Add("page_size", strconv.Itoa(request.PageSize))
	}
	if request.Filter.Status != "" {
		values.Add("filter.status", request.Filter.Status)
	}
	if request.Filter.Model != "" {
		values.Add("filter.model", request.Filter.Model)
	}
	for _, taskID := range request.Filter.TaskIDs {
		values.Add("filter.task_ids", taskID)
	}

	endpoint := fmt.Sprintf("%s?%s", c.fullURL(contentGenerationTasksSuffix), values.Encode())

	err = c.Do(ctx, http.MethodGet, endpoint, resourceTypeEndpoint, "", &response, setters...)
	if err != nil {
		return response, err
	}

	return response, nil
}
