package file

import (
	"bytes"
	"mime/multipart"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/apiform"
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
