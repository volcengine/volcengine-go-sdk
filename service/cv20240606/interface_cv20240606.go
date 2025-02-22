// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cv20240606iface provides an interface to enable mocking the CV20240606 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cv20240606

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// CV20240606API provides an interface to enable mocking the
// cv20240606.CV20240606 service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // CV20240606.
//    func myFunc(svc CV20240606API) bool {
//        // Make svc.AIGCStylizeImage request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cv20240606.New(sess)
//
//        myFunc(svc)
//    }
//
type CV20240606API interface {
	AIGCStylizeImageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AIGCStylizeImageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AIGCStylizeImageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AIGCStylizeImage(*AIGCStylizeImageInput) (*AIGCStylizeImageOutput, error)
	AIGCStylizeImageWithContext(volcengine.Context, *AIGCStylizeImageInput, ...request.Option) (*AIGCStylizeImageOutput, error)
	AIGCStylizeImageRequest(*AIGCStylizeImageInput) (*request.Request, *AIGCStylizeImageOutput)

	AIGCStylizeImageUsageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AIGCStylizeImageUsageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AIGCStylizeImageUsageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AIGCStylizeImageUsage(*AIGCStylizeImageUsageInput) (*AIGCStylizeImageUsageOutput, error)
	AIGCStylizeImageUsageWithContext(volcengine.Context, *AIGCStylizeImageUsageInput, ...request.Option) (*AIGCStylizeImageUsageOutput, error)
	AIGCStylizeImageUsageRequest(*AIGCStylizeImageUsageInput) (*request.Request, *AIGCStylizeImageUsageOutput)

	EmotionPortraitCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EmotionPortraitCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EmotionPortraitCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EmotionPortrait(*EmotionPortraitInput) (*EmotionPortraitOutput, error)
	EmotionPortraitWithContext(volcengine.Context, *EmotionPortraitInput, ...request.Option) (*EmotionPortraitOutput, error)
	EmotionPortraitRequest(*EmotionPortraitInput) (*request.Request, *EmotionPortraitOutput)

	EntitySegmentCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EntitySegmentCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EntitySegmentCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EntitySegment(*EntitySegmentInput) (*EntitySegmentOutput, error)
	EntitySegmentWithContext(volcengine.Context, *EntitySegmentInput, ...request.Option) (*EntitySegmentOutput, error)
	EntitySegmentRequest(*EntitySegmentInput) (*request.Request, *EntitySegmentOutput)

	FaceFusionMovieGetResultCommon(*map[string]interface{}) (*map[string]interface{}, error)
	FaceFusionMovieGetResultCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	FaceFusionMovieGetResultCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	FaceFusionMovieGetResult(*FaceFusionMovieGetResultInput) (*FaceFusionMovieGetResultOutput, error)
	FaceFusionMovieGetResultWithContext(volcengine.Context, *FaceFusionMovieGetResultInput, ...request.Option) (*FaceFusionMovieGetResultOutput, error)
	FaceFusionMovieGetResultRequest(*FaceFusionMovieGetResultInput) (*request.Request, *FaceFusionMovieGetResultOutput)

	FaceFusionMovieSubmitTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	FaceFusionMovieSubmitTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	FaceFusionMovieSubmitTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	FaceFusionMovieSubmitTask(*FaceFusionMovieSubmitTaskInput) (*FaceFusionMovieSubmitTaskOutput, error)
	FaceFusionMovieSubmitTaskWithContext(volcengine.Context, *FaceFusionMovieSubmitTaskInput, ...request.Option) (*FaceFusionMovieSubmitTaskOutput, error)
	FaceFusionMovieSubmitTaskRequest(*FaceFusionMovieSubmitTaskInput) (*request.Request, *FaceFusionMovieSubmitTaskOutput)

	FaceSwapCommon(*map[string]interface{}) (*map[string]interface{}, error)
	FaceSwapCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	FaceSwapCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	FaceSwap(*FaceSwapInput) (*FaceSwapOutput, error)
	FaceSwapWithContext(volcengine.Context, *FaceSwapInput, ...request.Option) (*FaceSwapOutput, error)
	FaceSwapRequest(*FaceSwapInput) (*request.Request, *FaceSwapOutput)

	FaceSwapAICommon(*map[string]interface{}) (*map[string]interface{}, error)
	FaceSwapAICommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	FaceSwapAICommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	FaceSwapAI(*FaceSwapAIInput) (*FaceSwapAIOutput, error)
	FaceSwapAIWithContext(volcengine.Context, *FaceSwapAIInput, ...request.Option) (*FaceSwapAIOutput, error)
	FaceSwapAIRequest(*FaceSwapAIInput) (*request.Request, *FaceSwapAIOutput)

	HairStyleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	HairStyleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	HairStyleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	HairStyle(*HairStyleInput) (*HairStyleOutput, error)
	HairStyleWithContext(volcengine.Context, *HairStyleInput, ...request.Option) (*HairStyleOutput, error)
	HairStyleRequest(*HairStyleInput) (*request.Request, *HairStyleOutput)

	HighAesAnimeV13Common(*map[string]interface{}) (*map[string]interface{}, error)
	HighAesAnimeV13CommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	HighAesAnimeV13CommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	HighAesAnimeV13(*HighAesAnimeV13Input) (*HighAesAnimeV13Output, error)
	HighAesAnimeV13WithContext(volcengine.Context, *HighAesAnimeV13Input, ...request.Option) (*HighAesAnimeV13Output, error)
	HighAesAnimeV13Request(*HighAesAnimeV13Input) (*request.Request, *HighAesAnimeV13Output)

	HighAesGeneralV13Common(*map[string]interface{}) (*map[string]interface{}, error)
	HighAesGeneralV13CommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	HighAesGeneralV13CommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	HighAesGeneralV13(*HighAesGeneralV13Input) (*HighAesGeneralV13Output, error)
	HighAesGeneralV13WithContext(volcengine.Context, *HighAesGeneralV13Input, ...request.Option) (*HighAesGeneralV13Output, error)
	HighAesGeneralV13Request(*HighAesGeneralV13Input) (*request.Request, *HighAesGeneralV13Output)

	HighAesGeneralV14Common(*map[string]interface{}) (*map[string]interface{}, error)
	HighAesGeneralV14CommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	HighAesGeneralV14CommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	HighAesGeneralV14(*HighAesGeneralV14Input) (*HighAesGeneralV14Output, error)
	HighAesGeneralV14WithContext(volcengine.Context, *HighAesGeneralV14Input, ...request.Option) (*HighAesGeneralV14Output, error)
	HighAesGeneralV14Request(*HighAesGeneralV14Input) (*request.Request, *HighAesGeneralV14Output)

	HighAesGeneralV14IPKeepCommon(*map[string]interface{}) (*map[string]interface{}, error)
	HighAesGeneralV14IPKeepCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	HighAesGeneralV14IPKeepCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	HighAesGeneralV14IPKeep(*HighAesGeneralV14IPKeepInput) (*HighAesGeneralV14IPKeepOutput, error)
	HighAesGeneralV14IPKeepWithContext(volcengine.Context, *HighAesGeneralV14IPKeepInput, ...request.Option) (*HighAesGeneralV14IPKeepOutput, error)
	HighAesGeneralV14IPKeepRequest(*HighAesGeneralV14IPKeepInput) (*request.Request, *HighAesGeneralV14IPKeepOutput)

	HighAesGeneralV20Common(*map[string]interface{}) (*map[string]interface{}, error)
	HighAesGeneralV20CommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	HighAesGeneralV20CommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	HighAesGeneralV20(*HighAesGeneralV20Input) (*HighAesGeneralV20Output, error)
	HighAesGeneralV20WithContext(volcengine.Context, *HighAesGeneralV20Input, ...request.Option) (*HighAesGeneralV20Output, error)
	HighAesGeneralV20Request(*HighAesGeneralV20Input) (*request.Request, *HighAesGeneralV20Output)

	HighAesGeneralV20LCommon(*map[string]interface{}) (*map[string]interface{}, error)
	HighAesGeneralV20LCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	HighAesGeneralV20LCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	HighAesGeneralV20L(*HighAesGeneralV20LInput) (*HighAesGeneralV20LOutput, error)
	HighAesGeneralV20LWithContext(volcengine.Context, *HighAesGeneralV20LInput, ...request.Option) (*HighAesGeneralV20LOutput, error)
	HighAesGeneralV20LRequest(*HighAesGeneralV20LInput) (*request.Request, *HighAesGeneralV20LOutput)

	HighAesIPV20Common(*map[string]interface{}) (*map[string]interface{}, error)
	HighAesIPV20CommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	HighAesIPV20CommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	HighAesIPV20(*HighAesIPV20Input) (*HighAesIPV20Output, error)
	HighAesIPV20WithContext(volcengine.Context, *HighAesIPV20Input, ...request.Option) (*HighAesIPV20Output, error)
	HighAesIPV20Request(*HighAesIPV20Input) (*request.Request, *HighAesIPV20Output)

	HignAesGeneralV12Common(*map[string]interface{}) (*map[string]interface{}, error)
	HignAesGeneralV12CommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	HignAesGeneralV12CommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	HignAesGeneralV12(*HignAesGeneralV12Input) (*HignAesGeneralV12Output, error)
	HignAesGeneralV12WithContext(volcengine.Context, *HignAesGeneralV12Input, ...request.Option) (*HignAesGeneralV12Output, error)
	HignAesGeneralV12Request(*HignAesGeneralV12Input) (*request.Request, *HignAesGeneralV12Output)

	Img2ImgInpaintingCommon(*map[string]interface{}) (*map[string]interface{}, error)
	Img2ImgInpaintingCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	Img2ImgInpaintingCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	Img2ImgInpainting(*Img2ImgInpaintingInput) (*Img2ImgInpaintingOutput, error)
	Img2ImgInpaintingWithContext(volcengine.Context, *Img2ImgInpaintingInput, ...request.Option) (*Img2ImgInpaintingOutput, error)
	Img2ImgInpaintingRequest(*Img2ImgInpaintingInput) (*request.Request, *Img2ImgInpaintingOutput)

	Img2ImgInpaintingEditCommon(*map[string]interface{}) (*map[string]interface{}, error)
	Img2ImgInpaintingEditCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	Img2ImgInpaintingEditCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	Img2ImgInpaintingEdit(*Img2ImgInpaintingEditInput) (*Img2ImgInpaintingEditOutput, error)
	Img2ImgInpaintingEditWithContext(volcengine.Context, *Img2ImgInpaintingEditInput, ...request.Option) (*Img2ImgInpaintingEditOutput, error)
	Img2ImgInpaintingEditRequest(*Img2ImgInpaintingEditInput) (*request.Request, *Img2ImgInpaintingEditOutput)

	Img2ImgOutpaintingCommon(*map[string]interface{}) (*map[string]interface{}, error)
	Img2ImgOutpaintingCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	Img2ImgOutpaintingCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	Img2ImgOutpainting(*Img2ImgOutpaintingInput) (*Img2ImgOutpaintingOutput, error)
	Img2ImgOutpaintingWithContext(volcengine.Context, *Img2ImgOutpaintingInput, ...request.Option) (*Img2ImgOutpaintingOutput, error)
	Img2ImgOutpaintingRequest(*Img2ImgOutpaintingInput) (*request.Request, *Img2ImgOutpaintingOutput)

	Img2ImgXLSftCommon(*map[string]interface{}) (*map[string]interface{}, error)
	Img2ImgXLSftCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	Img2ImgXLSftCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	Img2ImgXLSft(*Img2ImgXLSftInput) (*Img2ImgXLSftOutput, error)
	Img2ImgXLSftWithContext(volcengine.Context, *Img2ImgXLSftInput, ...request.Option) (*Img2ImgXLSftOutput, error)
	Img2ImgXLSftRequest(*Img2ImgXLSftInput) (*request.Request, *Img2ImgXLSftOutput)

	LensLqirCommon(*map[string]interface{}) (*map[string]interface{}, error)
	LensLqirCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	LensLqirCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	LensLqir(*LensLqirInput) (*LensLqirOutput, error)
	LensLqirWithContext(volcengine.Context, *LensLqirInput, ...request.Option) (*LensLqirOutput, error)
	LensLqirRequest(*LensLqirInput) (*request.Request, *LensLqirOutput)

	LensNnsr2PicCommonCommon(*map[string]interface{}) (*map[string]interface{}, error)
	LensNnsr2PicCommonCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	LensNnsr2PicCommonCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	LensNnsr2PicCommon(*LensNnsr2PicCommonInput) (*LensNnsr2PicCommonOutput, error)
	LensNnsr2PicCommonWithContext(volcengine.Context, *LensNnsr2PicCommonInput, ...request.Option) (*LensNnsr2PicCommonOutput, error)
	LensNnsr2PicCommonRequest(*LensNnsr2PicCommonInput) (*request.Request, *LensNnsr2PicCommonOutput)

	LensOprCommon(*map[string]interface{}) (*map[string]interface{}, error)
	LensOprCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	LensOprCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	LensOpr(*LensOprInput) (*LensOprOutput, error)
	LensOprWithContext(volcengine.Context, *LensOprInput, ...request.Option) (*LensOprOutput, error)
	LensOprRequest(*LensOprInput) (*request.Request, *LensOprOutput)

	LensVidaVideoGetResultCommon(*map[string]interface{}) (*map[string]interface{}, error)
	LensVidaVideoGetResultCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	LensVidaVideoGetResultCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	LensVidaVideoGetResult(*LensVidaVideoGetResultInput) (*LensVidaVideoGetResultOutput, error)
	LensVidaVideoGetResultWithContext(volcengine.Context, *LensVidaVideoGetResultInput, ...request.Option) (*LensVidaVideoGetResultOutput, error)
	LensVidaVideoGetResultRequest(*LensVidaVideoGetResultInput) (*request.Request, *LensVidaVideoGetResultOutput)

	LensVidaVideoSubmitTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	LensVidaVideoSubmitTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	LensVidaVideoSubmitTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	LensVidaVideoSubmitTask(*LensVidaVideoSubmitTaskInput) (*LensVidaVideoSubmitTaskOutput, error)
	LensVidaVideoSubmitTaskWithContext(volcengine.Context, *LensVidaVideoSubmitTaskInput, ...request.Option) (*LensVidaVideoSubmitTaskOutput, error)
	LensVidaVideoSubmitTaskRequest(*LensVidaVideoSubmitTaskInput) (*request.Request, *LensVidaVideoSubmitTaskOutput)

	LensVideoNnsrGetResultCommon(*map[string]interface{}) (*map[string]interface{}, error)
	LensVideoNnsrGetResultCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	LensVideoNnsrGetResultCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	LensVideoNnsrGetResult(*LensVideoNnsrGetResultInput) (*LensVideoNnsrGetResultOutput, error)
	LensVideoNnsrGetResultWithContext(volcengine.Context, *LensVideoNnsrGetResultInput, ...request.Option) (*LensVideoNnsrGetResultOutput, error)
	LensVideoNnsrGetResultRequest(*LensVideoNnsrGetResultInput) (*request.Request, *LensVideoNnsrGetResultOutput)

	LensVideoNnsrSubmitTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	LensVideoNnsrSubmitTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	LensVideoNnsrSubmitTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	LensVideoNnsrSubmitTask(*LensVideoNnsrSubmitTaskInput) (*LensVideoNnsrSubmitTaskOutput, error)
	LensVideoNnsrSubmitTaskWithContext(volcengine.Context, *LensVideoNnsrSubmitTaskInput, ...request.Option) (*LensVideoNnsrSubmitTaskOutput, error)
	LensVideoNnsrSubmitTaskRequest(*LensVideoNnsrSubmitTaskInput) (*request.Request, *LensVideoNnsrSubmitTaskOutput)

	MaintainIDUsageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	MaintainIDUsageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	MaintainIDUsageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	MaintainIDUsage(*MaintainIDUsageInput) (*MaintainIDUsageOutput, error)
	MaintainIDUsageWithContext(volcengine.Context, *MaintainIDUsageInput, ...request.Option) (*MaintainIDUsageOutput, error)
	MaintainIDUsageRequest(*MaintainIDUsageInput) (*request.Request, *MaintainIDUsageOutput)

	PhotoVerseCommon(*map[string]interface{}) (*map[string]interface{}, error)
	PhotoVerseCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	PhotoVerseCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	PhotoVerse(*PhotoVerseInput) (*PhotoVerseOutput, error)
	PhotoVerseWithContext(volcengine.Context, *PhotoVerseInput, ...request.Option) (*PhotoVerseOutput, error)
	PhotoVerseRequest(*PhotoVerseInput) (*request.Request, *PhotoVerseOutput)

	Text2ImgXLSftCommon(*map[string]interface{}) (*map[string]interface{}, error)
	Text2ImgXLSftCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	Text2ImgXLSftCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	Text2ImgXLSft(*Text2ImgXLSftInput) (*Text2ImgXLSftOutput, error)
	Text2ImgXLSftWithContext(volcengine.Context, *Text2ImgXLSftInput, ...request.Option) (*Text2ImgXLSftOutput, error)
	Text2ImgXLSftRequest(*Text2ImgXLSftInput) (*request.Request, *Text2ImgXLSftOutput)
}

var _ CV20240606API = (*CV20240606)(nil)
