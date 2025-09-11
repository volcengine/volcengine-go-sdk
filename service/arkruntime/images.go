package arkruntime

import (
	"context"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/utils"
)

const generateImagesPath = "/images/generations"

func (c *Client) GenerateImages(
	ctx context.Context,
	request model.GenerateImagesRequest,
	setters ...requestOption,
) (response model.ImagesResponse, err error) {
	if !c.isAPIKeyAuthentication() {
		return response, model.ErrAKSKNotSupported
	}

	if err = request.NormalizeImages(); err != nil {
		return
	}

	requestOptions := append(setters, withBody(request))
	err = c.Do(ctx, http.MethodPost, c.fullURL(generateImagesPath), resourceTypeEndpoint, request.Model, &response, requestOptions...)
	return
}

func (c *Client) GenerateImagesStreaming(ctx context.Context, request model.GenerateImagesRequest, setters ...requestOption) (stream *utils.ImageGenerationStreamReader, err error) {
	if !c.isAPIKeyAuthentication() {
		return stream, model.ErrAKSKNotSupported
	}

	if err = request.NormalizeImages(); err != nil {
		return
	}

	bodyMap, err := structToMap(request)
	if err != nil {
		return
	}
	bodyMap["stream"] = true
	requestOptions := append(setters, withBody(bodyMap))

	resp, err := c.ImageGenerationStreamDo(ctx, http.MethodPost, c.fullURL(generateImagesPath), request.Model, requestOptions...)
	if err != nil {
		return
	}
	stream = resp
	return
}
