package utils

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type StreamReader struct {
	IsFinished bool
	Reader     *bufio.Reader
	Response   *http.Response
}

type ResponseData struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId"`
		Region    string `json:"Region"`
	} `json:"ResponseMetadata"`
	Result struct {
		Suggest string `json:"Suggest"`
	} `json:"Result"`
}

type StreamResponse struct {
	Event string        `json:"Event"`
	Data  *ResponseData `json:"Data"`
}

func (stream *StreamReader) Recv() (response *StreamResponse, err error) {
	response, err = stream.processLines()
	if stream.IsFinished {
		err = io.EOF
		return
	}
	return
}

func (stream *StreamReader) processLines() (*StreamResponse, error) {
	response := &StreamResponse{}
	for {
		var event string
		line, err := stream.Reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				stream.IsFinished = true
				return response, nil
			}
			return response, err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "event:") {
			event = strings.TrimPrefix(line, "event:")
			event = strings.TrimSpace(event)
			response.Event = event
		} else if strings.HasPrefix(line, "data:") {
			dataStr := strings.TrimPrefix(line, "data:")
			dataStr = strings.TrimSpace(dataStr)
			response.Data = &ResponseData{}
			err = json.Unmarshal([]byte(dataStr), response.Data)
			if err != nil {
				return response, err
			}
		}
		if response != nil && response.Event != "" && response.Data != nil {
			return response, nil
		}
	}
}

func (stream *StreamReader) Close() error {
	if stream.Response != nil && stream.Response.Body != nil {
		return stream.Response.Body.Close()
	}
	return nil
}
