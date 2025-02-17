package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/waf"
	"github.com/volcengine/volcengine-go-sdk/service/wafruntime"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"io"
)

func main() {
	ak := ""               // replace real: api ak
	sk := ""               // replace real: api sk
	region := "cn-beijing" // replace real: api region
	config := volcengine.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithRegion(region)

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	client := wafruntime.New(sess)

	// judge prompt
	var respCheckLLMPrompt *waf.CheckLLMPromptOutput
	respCheckLLMPrompt, err = client.CheckLLMPrompt(&waf.CheckLLMPromptInput{
		Content:     volcengine.String("xxx"),
		ContentType: volcengine.Int32(1),
		Host:        volcengine.String("xxx.xxx.com"), // replace real: waf host
		Region:      volcengine.String("cn-beijing"),  // replace real: waf region
	})

	if err != nil {
		panic(fmt.Sprintf("Call CheckLLMPrompt err: %v\n", err))
	}

	fmt.Println(respCheckLLMPrompt.String())

	// llm generate stream
	var respQueryLLMGenerateStream *wafruntime.QueryLLMGenerateStreamOutput
	respQueryLLMGenerateStream, err = client.QueryLLMGenerateStream(
		&wafruntime.QueryLLMGenerateStreamInput{
			MsgID:     volcengine.String(*respCheckLLMPrompt.MsgID),
			UseStream: volcengine.Bool(true),
		})

	if err != nil {
		panic(fmt.Sprintf("Call QueryLLMGenerate err: %v\n", err))
	}

	defer func(respQueryLLMGenerateStream *wafruntime.QueryLLMGenerateStreamOutput) {
		err := respQueryLLMGenerateStream.Close()
		if err != nil {
			panic(err)
		}
	}(respQueryLLMGenerateStream)

	for {
		data, err := respQueryLLMGenerateStream.Recv()
		if err != nil {
			if err == io.EOF {
				return
			}

			fmt.Printf("Stream chat error: %v\n", err)
			return
		}

		dataBytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n\n", string(dataBytes))
	}

}
