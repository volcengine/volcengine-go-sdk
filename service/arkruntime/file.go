package arkruntime

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/file"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/utils"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

const (
	filePrefix   = "/files"
	fileScheme   = "file"
	pollInterval = 3 * time.Second
	maxWaitTime  = 10 * time.Minute
)

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

func (c *Client) preprocessResponseInput(ctx context.Context, input *responses.ResponsesInput) error {
	return c.preprocessResponseMultiModal(ctx, input)
}

func (c *Client) preprocessResponseMultiModal(ctx context.Context, input *responses.ResponsesInput) error {
	if len(input.GetListValue().GetListValue()) == 0 { // empty multi modal input items
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, maxWaitTime)
	defer cancel()

	eg, ctx := utils.WithContext(ctx)
	inputList := input.GetListValue().GetListValue()
	for _, item := range inputList {
		for _, contentItem := range item.GetInputMessage().GetContent() {
			multiModalFile, err := getMultiModalFile(contentItem)
			if err != nil {
				return err
			}

			contentItem := contentItem
			eg.Go(func() error {
				defer func() {
					if multiModalFile != nil {
						_ = multiModalFile.Close()
					}
				}()
				return c.preprocessResponseFile(ctx, contentItem, multiModalFile)
			})
		}
	}
	// wait for all files to be processed
	return eg.Wait()
}

func (c *Client) preprocessResponseFile(ctx context.Context, contentItem *responses.ContentItem, fileReader io.Reader) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("preprocessResponseContentItem recovered from panic: %v", r)
		}
	}()
	// upload local file to file service
	fileMeta, err := c.UploadFile(ctx, &file.UploadFileRequest{
		File:    fileReader,
		Purpose: file.PurposeUserData,
	})
	if err != nil {
		return err
	}
	fileMeta, err = c.RetrieveFile(ctx, fileMeta.ID)
	if err != nil {
		return err
	}
	// wait for file status to be ready
	for fileMeta.Status == file.StatusProcessing {
		time.Sleep(pollInterval)
		fileMeta, err = c.RetrieveFile(ctx, fileMeta.ID)
		if err != nil {
			return err
		}
	}

	if contentItem.GetVideo() != nil && len(contentItem.GetVideo().GetVideoUrl()) > 0 {
		contentItem.GetVideo().VideoUrl = ""
		contentItem.GetVideo().FileId = volcengine.String(fileMeta.ID)
	} else if contentItem.GetImage() != nil && len(contentItem.GetImage().GetImageUrl()) > 0 {
		contentItem.GetImage().ImageUrl = nil
		contentItem.GetImage().FileId = volcengine.String(fileMeta.ID)
	} else if contentItem.GetFile() != nil && len(contentItem.GetFile().GetFileUrl()) > 0 {
		contentItem.GetFile().FileUrl = nil
		contentItem.GetFile().FileId = volcengine.String(fileMeta.ID)
	}
	return nil
}

func getMultiModalFile(contentItem *responses.ContentItem) (io.ReadCloser, error) {
	if len(contentItem.GetVideo().GetVideoUrl()) > 0 {
		return parseFile(contentItem.GetVideo().GetVideoUrl())
	} else if len(contentItem.GetImage().GetImageUrl()) > 0 {
		return parseFile(contentItem.GetImage().GetImageUrl())
	} else if len(contentItem.GetFile().GetFileUrl()) > 0 {
		return parseFile(contentItem.GetFile().GetFileUrl())
	}
	return nil, nil
}

func parseFile(rawURL string) (io.ReadCloser, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if u.Scheme != fileScheme { // only parse file scheme locally
		return nil, nil
	}
	// On Unix-like systems, u.Host is typically empty.
	// On Windows, u.Host may contain a drive letter or UNC share.
	var path string
	if u.Host != "" && u.Host != "localhost" {
		// Example: file://server/share/file.txt → u.Host="server", u.Path="/share/file.txt"
		// Build UNC path: \\server\share\file.txt on Windows
		// Note: filepath.FromSlash will convert "/" to "\" on Windows
		path = fmt.Sprintf(`%s%s`, u.Host, u.Path)
	} else {
		path = u.Path
	}

	// On Windows, url.Parse yields a leading “/C:/…” form for "file:///C:/…"
	if runtime.GOOS == "windows" {
		// Remove leading slash if path starts with slash + drive letter
		if len(path) >= 3 && strings.HasPrefix(path, "/") && path[2] == ':' {
			path = path[1:]
		}
	}
	return os.Open(path)
}
