package arkruntime

import (
	"context"
	"errors"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/file"
)

const filePrefix = "/files"

func (c *Client) UploadFile(ctx context.Context, request *file.UploadFileRequest, setters ...requestOption) (response *file.FileMeta, err error) {
	opts := append(setters, withBody(request))
	response = &file.FileMeta{}
	// files do not support ak/sk authorization, do not have resource
	err = c.Do(ctx, http.MethodPost, c.fullURL(filePrefix), "", "", response, opts...)
	return
}

func (c *Client) RetrieveFile(ctx context.Context, fileID string, setters ...requestOption) (response *file.FileMeta, err error) {
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	opts := append(setters, withBody(nil))
	response = &file.FileMeta{}
	// files do not support ak/sk authorization, do not have resource
	err = c.Do(ctx, http.MethodGet, c.fullURL(filePrefix+"/"+fileID), "", "", response, opts...)
	return
}
func (c *Client) ListFiles(ctx context.Context, request *file.ListFilesRequest, setters ...requestOption) (response *file.ListFilesResponse, err error) {
	opts := append(setters, withBody(request))
	response = &file.ListFilesResponse{}
	// files do not support ak/sk authorization, do not have resource
	err = c.Do(ctx, http.MethodGet, c.fullURL(filePrefix), "", "", response, opts...)
	return
}
func (c *Client) DeleteFile(ctx context.Context, fileID string, setters ...requestOption) (response *file.DeleteFileResponse, err error) {
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	opts := append(setters, withBody(nil))
	response = &file.DeleteFileResponse{}
	// files do not support ak/sk authorization, do not have resource
	err = c.Do(ctx, http.MethodDelete, c.fullURL(filePrefix+"/"+fileID), "", "", response, opts...)
	return
}
