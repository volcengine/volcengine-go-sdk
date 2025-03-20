// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package translate20250301iface provides an interface to enable mocking the TRANSLATE20250301 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package translate20250301

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// TRANSLATE20250301API provides an interface to enable mocking the
// translate20250301.TRANSLATE20250301 service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // TRANSLATE20250301.
//    func myFunc(svc TRANSLATE20250301API) bool {
//        // Make svc.LangDetect request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := translate20250301.New(sess)
//
//        myFunc(svc)
//    }
//
type TRANSLATE20250301API interface {
	LangDetectCommon(*map[string]interface{}) (*map[string]interface{}, error)
	LangDetectCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	LangDetectCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	LangDetect(*LangDetectInput) (*LangDetectOutput, error)
	LangDetectWithContext(volcengine.Context, *LangDetectInput, ...request.Option) (*LangDetectOutput, error)
	LangDetectRequest(*LangDetectInput) (*request.Request, *LangDetectOutput)

	QueryAudioCommon(*map[string]interface{}) (*map[string]interface{}, error)
	QueryAudioCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	QueryAudioCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	QueryAudio(*QueryAudioInput) (*QueryAudioOutput, error)
	QueryAudioWithContext(volcengine.Context, *QueryAudioInput, ...request.Option) (*QueryAudioOutput, error)
	QueryAudioRequest(*QueryAudioInput) (*request.Request, *QueryAudioOutput)

	SubmitAudioCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SubmitAudioCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SubmitAudioCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SubmitAudio(*SubmitAudioInput) (*SubmitAudioOutput, error)
	SubmitAudioWithContext(volcengine.Context, *SubmitAudioInput, ...request.Option) (*SubmitAudioOutput, error)
	SubmitAudioRequest(*SubmitAudioInput) (*request.Request, *SubmitAudioOutput)

	TranslateImageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	TranslateImageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	TranslateImageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	TranslateImage(*TranslateImageInput) (*TranslateImageOutput, error)
	TranslateImageWithContext(volcengine.Context, *TranslateImageInput, ...request.Option) (*TranslateImageOutput, error)
	TranslateImageRequest(*TranslateImageInput) (*request.Request, *TranslateImageOutput)

	TranslateTextCommon(*map[string]interface{}) (*map[string]interface{}, error)
	TranslateTextCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	TranslateTextCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	TranslateText(*TranslateTextInput) (*TranslateTextOutput, error)
	TranslateTextWithContext(volcengine.Context, *TranslateTextInput, ...request.Option) (*TranslateTextOutput, error)
	TranslateTextRequest(*TranslateTextInput) (*request.Request, *TranslateTextOutput)
}

var _ TRANSLATE20250301API = (*TRANSLATE20250301)(nil)
