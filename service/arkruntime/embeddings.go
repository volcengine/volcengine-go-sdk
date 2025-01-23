package arkruntime

import (
	"context"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

const (
	embeddingsSuffix           = "/embeddings"
	multimodalEmbeddingsSuffix = "/embeddings/multimodal"
)

// CreateEmbeddings returns an EmbeddingResponse which will contain an Embedding for every item in |body.Input|.
//
// Body should be of type EmbeddingRequestStrings for embedding strings or EmbeddingRequestTokens
// for embedding groups of text already converted to tokens.
func (c *Client) CreateEmbeddings(
	ctx context.Context,
	conv model.EmbeddingRequestConverter,
	setters ...requestOption,
) (res model.EmbeddingResponse, err error) {
	baseReq := conv.Convert()

	requestOptions := append(setters, withBody(baseReq))
	if baseReq.EncodingFormat != model.EmbeddingEncodingFormatBase64 {
		err = c.Do(ctx, http.MethodPost, c.fullURL(embeddingsSuffix), resourceTypeEndpoint, baseReq.Model, &res, requestOptions...)
		return
	}

	base64Response := &model.EmbeddingResponseBase64{}
	err = c.Do(ctx, http.MethodPost, c.fullURL(embeddingsSuffix), resourceTypeEndpoint, baseReq.Model, base64Response, requestOptions...)
	if err != nil {
		return
	}

	res, err = base64Response.ToEmbeddingResponse()
	return
}

// CreateMultiModalEmbeddings returns an MultimodalEmbeddingResponse which will contain a MultimodalEmbedding for |body.Input|.
// Body should be MultiModalEmbeddingRequest.
func (c *Client) CreateMultiModalEmbeddings(
	ctx context.Context,
	request model.MultiModalEmbeddingRequest,
	setters ...requestOption,
) (res model.MultimodalEmbeddingResponse, err error) {

	requestOptions := append(setters, withBody(request))
	if request.EncodingFormat != nil && *request.EncodingFormat == model.EmbeddingEncodingFormatBase64 {
		base64Response := &model.MultiModalEmbeddingResponseBase64{}
		err = c.Do(ctx, http.MethodPost, c.fullURL(multimodalEmbeddingsSuffix), resourceTypeEndpoint, request.Model, base64Response, requestOptions...)
		if err != nil {
			return
		}
		res, err = base64Response.ToMultiModalEmbeddingResponse()
		return
	}

	err = c.Do(ctx, http.MethodPost, c.fullURL(multimodalEmbeddingsSuffix), resourceTypeEndpoint, request.Model, &res, requestOptions...)
	return
}
