package arkruntime

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

const contentGenerationTaskPath = "/contents/generations/tasks"

func (c *Client) CreateContentGenerationTask(
	ctx context.Context,
	request model.CreateContentGenerationTaskRequest,
	setters ...requestOption,
) (response model.CreateContentGenerationTaskResponse, err error) {
	if !c.isAPIKeyAuthentication() {
		return response, model.ErrAKSKNotSupported
	}

	requestOptions := append(setters, withBody(request))
	err = c.Do(ctx, http.MethodPost, c.fullURL(contentGenerationTaskPath), resourceTypeEndpoint, request.Model, &response, requestOptions...)
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

	url := fmt.Sprintf("%s/%s", c.fullURL(contentGenerationTaskPath), request.ID)

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

	url := fmt.Sprintf("%s/%s", c.fullURL(contentGenerationTaskPath), request.ID)

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
	if pageNum := request.PageNum; pageNum != nil && *pageNum > 0 {
		values.Add("page_num", strconv.Itoa(*pageNum))
	}
	if pageSize := request.PageSize; pageSize != nil && *pageSize > 0 {
		values.Add("page_size", strconv.Itoa(*pageSize))
	}

	if filter := request.Filter; filter != nil {
		if status := filter.Status; status != nil && *status != "" {
			values.Add("filter.status", *status)
		}
		if model := filter.Model; model != nil && *model != "" {
			values.Add("filter.model", *model)
		}
		for _, taskID := range filter.TaskIDs {
			values.Add("filter.task_ids", *taskID)
		}
	}

	endpoint := fmt.Sprintf("%s?%s", c.fullURL(contentGenerationTaskPath), values.Encode())

	err = c.Do(ctx, http.MethodGet, endpoint, resourceTypeEndpoint, "", &response, setters...)
	if err != nil {
		return response, err
	}

	return response, nil
}
