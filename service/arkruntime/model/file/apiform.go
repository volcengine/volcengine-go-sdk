package file

import (
	"bytes"
	"mime/multipart"
	"net/url"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/apiform"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/apiquery"
)

func (p *UploadFileRequest) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(p, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

func (p *ListFilesRequest) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(p, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
