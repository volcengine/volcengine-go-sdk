package file

import (
	"fmt"
	"io"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

const (
	PurposeUserData = "user_data"

	StatusActive = "active"

	StatusProcessing = "processing"

	StatusFailed = "failed"

	ObjectTypeList = "list"

	ObjectTypeFile = "file"

	OrderAsc = "asc"

	OrderDesc = "desc"
)

type Purpose string

type Status string

type ObjectType string

type Order string

type Error struct {
	Code    string `thrift:"Code,1,required" form:"code,required" json:"code,required"`
	Message string `thrift:"Message,2,required" form:"message,required" json:"message,required"`
	Param   string `thrift:"Param,3,required" form:"param,required" json:"param,required"`
	Type    string `thrift:"Type,4,required" form:"type,required" json:"type,required"`
}

func NewError() *Error {
	return &Error{}
}

func (p *Error) InitDefault() {
}

func (p *Error) GetCode() (v string) {
	return p.Code
}

func (p *Error) GetMessage() (v string) {
	return p.Message
}

func (p *Error) GetParam() (v string) {
	return p.Param
}

func (p *Error) GetType() (v string) {
	return p.Type
}

func (p *Error) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Error(%+v)", *p)
}

type FileMeta struct {
	ObjectType ObjectType `thrift:"ObjectType,1,required" form:"object,required" json:"object,required"`
	ID         string     `thrift:"ID,2,required" form:"id,required" json:"id,required"`
	Purpose    Purpose    `thrift:"Purpose,3,required" form:"purpose,required" json:"purpose,required"`
	FileName   string     `thrift:"FileName,4,required" form:"filename,required" json:"filename,required"`
	Bytes      *int64     `thrift:"Bytes,5,optional" form:"bytes" json:"bytes,omitempty"`
	MimeType   *string    `thrift:"MimeType,6,optional" form:"mime_type" json:"mime_type,omitempty"`
	/*
	   time stamp for file
	*/
	CreatedAt int64  `thrift:"CreatedAt,7,required" form:"created_at,required" json:"created_at,required"`
	ExpireAt  int64  `thrift:"ExpireAt,8,required" form:"expire_at,required" json:"expire_at,required"`
	Status    Status `thrift:"Status,9,required" form:"status,required" json:"status,required"`

	// only set when status = failed
	Error             *Error             `thrift:"Error,11,optional" form:"error" json:"error,omitempty"`
	PreprocessConfigs *PreprocessConfigs `thrift:"PreprocessConfigs,12,optional" form:"preprocess_configs" json:"preprocess_configs,omitempty"`
	model.HttpHeader
}

func NewFileMeta() *FileMeta {
	return &FileMeta{}
}

func (p *FileMeta) InitDefault() {
}

func (p *FileMeta) GetObjectType() (v ObjectType) {
	return p.ObjectType
}

func (p *FileMeta) GetID() (v string) {
	return p.ID
}

func (p *FileMeta) GetPurpose() (v Purpose) {
	return p.Purpose
}

func (p *FileMeta) GetFileName() (v string) {
	return p.FileName
}

var FileMeta_Bytes_DEFAULT int64

func (p *FileMeta) GetBytes() (v int64) {
	if !p.IsSetBytes() {
		return FileMeta_Bytes_DEFAULT
	}
	return *p.Bytes
}

var FileMeta_MimeType_DEFAULT string

func (p *FileMeta) GetMimeType() (v string) {
	if !p.IsSetMimeType() {
		return FileMeta_MimeType_DEFAULT
	}
	return *p.MimeType
}

func (p *FileMeta) GetCreatedAt() (v int64) {
	return p.CreatedAt
}

func (p *FileMeta) GetExpireAt() (v int64) {
	return p.ExpireAt
}

func (p *FileMeta) GetStatus() (v Status) {
	return p.Status
}

var FileMeta_Error_DEFAULT *Error

func (p *FileMeta) GetError() (v *Error) {
	if !p.IsSetError() {
		return FileMeta_Error_DEFAULT
	}
	return p.Error
}

var FileMeta_PreprocessConfigs_DEFAULT *PreprocessConfigs

func (p *FileMeta) GetPreprocessConfigs() (v *PreprocessConfigs) {
	if !p.IsSetPreprocessConfigs() {
		return FileMeta_PreprocessConfigs_DEFAULT
	}
	return p.PreprocessConfigs
}

func (p *FileMeta) IsSetBytes() bool {
	return p.Bytes != nil
}

func (p *FileMeta) IsSetMimeType() bool {
	return p.MimeType != nil
}

func (p *FileMeta) IsSetError() bool {
	return p.Error != nil
}

func (p *FileMeta) IsSetPreprocessConfigs() bool {
	return p.PreprocessConfigs != nil
}

func (p *FileMeta) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FileMeta(%+v)", *p)
}

type Video struct {
	Fps *float64 `thrift:"fps,1,optional" form:"fps" json:"fps,omitempty"`
}

func NewVideo() *Video {
	return &Video{}
}

func (p *Video) InitDefault() {
}

var Video_Fps_DEFAULT float64

func (p *Video) GetFps() (v float64) {
	if !p.IsSetFps() {
		return Video_Fps_DEFAULT
	}
	return *p.Fps
}

func (p *Video) IsSetFps() bool {
	return p.Fps != nil
}

func (p *Video) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Video(%+v)", *p)
}

type PreprocessConfigs struct {
	Video *Video `thrift:"video,1,optional" form:"video" json:"video,omitempty"`
}

func NewPreprocessConfigs() *PreprocessConfigs {
	return &PreprocessConfigs{}
}

func (p *PreprocessConfigs) InitDefault() {
}

var PreprocessConfigs_Video_DEFAULT *Video

func (p *PreprocessConfigs) GetVideo() (v *Video) {
	if !p.IsSetVideo() {
		return PreprocessConfigs_Video_DEFAULT
	}
	return p.Video
}

func (p *PreprocessConfigs) IsSetVideo() bool {
	return p.Video != nil
}

func (p *PreprocessConfigs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PreprocessConfigs(%+v)", *p)
}

// UploadFileRequest upload file request, it's a form request with file field
type UploadFileRequest struct {
	// The File object (not file name) to be uploaded.
	File io.Reader `json:"file,required" format:"binary"`
	// The intended purpose of the uploaded file.
	Purpose Purpose `form:"purpose,required"`
	// The preprocessing configurations for the uploaded file.
	PreprocessConfigs *PreprocessConfigs `form:"preprocess_configs,omitempty"`
	// The expiration timestamp for a file.
	ExpireAt *int64 `form:"expire_at,omitempty"`
}

func NewUploadFileRequest() *UploadFileRequest {
	return &UploadFileRequest{}
}

func (p *UploadFileRequest) InitDefault() {
}

func (p *UploadFileRequest) GetPurpose() (v Purpose) {
	return p.Purpose
}

var UploadFileRequest_PreprocessConfigs_DEFAULT *PreprocessConfigs

func (p *UploadFileRequest) GetPreprocessConfigs() (v *PreprocessConfigs) {
	if !p.IsSetPreprocessConfigs() {
		return UploadFileRequest_PreprocessConfigs_DEFAULT
	}
	return p.PreprocessConfigs
}

var UploadFileRequest_ExpireAt_DEFAULT int64

func (p *UploadFileRequest) GetExpireAt() (v int64) {
	if !p.IsSetExpireAt() {
		return UploadFileRequest_ExpireAt_DEFAULT
	}
	return *p.ExpireAt
}

func (p *UploadFileRequest) IsSetPreprocessConfigs() bool {
	return p.PreprocessConfigs != nil
}

func (p *UploadFileRequest) IsSetExpireAt() bool {
	return p.ExpireAt != nil
}

func (p *UploadFileRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UploadFileRequest(%+v)", *p)
}

type RetrieveFileRequest struct {
	ID string `thrift:"ID,1,required" json:"ID,required" path:"id,required"`
}

func NewRetrieveFileRequest() *RetrieveFileRequest {
	return &RetrieveFileRequest{}
}

func (p *RetrieveFileRequest) InitDefault() {
}

func (p *RetrieveFileRequest) GetID() (v string) {
	return p.ID
}

func (p *RetrieveFileRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RetrieveFileRequest(%+v)", *p)
}

type ListFilesRequest struct {
	Purpose *Purpose `thrift:"Purpose,1,optional" json:"Purpose,omitempty" query:"purpose"`
	After   *string  `thrift:"After,2,optional" json:"After,omitempty" query:"after"`
	Limit   *int64   `thrift:"Limit,3,optional" json:"Limit,omitempty" query:"limit"`
	Order   *Order   `thrift:"Order,4,optional" json:"Order,omitempty" query:"order"`
}

func NewListFilesRequest() *ListFilesRequest {
	return &ListFilesRequest{}
}

func (p *ListFilesRequest) InitDefault() {
}

var ListFilesRequest_Purpose_DEFAULT Purpose

func (p *ListFilesRequest) GetPurpose() (v Purpose) {
	if !p.IsSetPurpose() {
		return ListFilesRequest_Purpose_DEFAULT
	}
	return *p.Purpose
}

var ListFilesRequest_After_DEFAULT string

func (p *ListFilesRequest) GetAfter() (v string) {
	if !p.IsSetAfter() {
		return ListFilesRequest_After_DEFAULT
	}
	return *p.After
}

var ListFilesRequest_Limit_DEFAULT int64

func (p *ListFilesRequest) GetLimit() (v int64) {
	if !p.IsSetLimit() {
		return ListFilesRequest_Limit_DEFAULT
	}
	return *p.Limit
}

var ListFilesRequest_Order_DEFAULT Order

func (p *ListFilesRequest) GetOrder() (v Order) {
	if !p.IsSetOrder() {
		return ListFilesRequest_Order_DEFAULT
	}
	return *p.Order
}

func (p *ListFilesRequest) IsSetPurpose() bool {
	return p.Purpose != nil
}

func (p *ListFilesRequest) IsSetAfter() bool {
	return p.After != nil
}

func (p *ListFilesRequest) IsSetLimit() bool {
	return p.Limit != nil
}

func (p *ListFilesRequest) IsSetOrder() bool {
	return p.Order != nil
}

func (p *ListFilesRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ListFilesRequest(%+v)", *p)
}

type ListFilesResponse struct {
	ObjectType ObjectType  `thrift:"ObjectType,1,required" form:"object,required" json:"object,required"`
	Data       []*FileMeta `thrift:"data,2,required" form:"data,required" json:"data,required"`
	FirstID    string      `thrift:"FirstID,3,required" form:"first_id,required" json:"first_id,required"`
	LastID     string      `thrift:"LastID,4,required" form:"last_id,required" json:"last_id,required"`
	HasMore    bool        `thrift:"HasMore,5,required" form:"has_more,required" json:"has_more,required"`
	model.HttpHeader
}

func NewListFilesResponse() *ListFilesResponse {
	return &ListFilesResponse{}
}

func (p *ListFilesResponse) InitDefault() {
}

func (p *ListFilesResponse) GetObjectType() (v ObjectType) {
	return p.ObjectType
}

func (p *ListFilesResponse) GetData() (v []*FileMeta) {
	return p.Data
}

func (p *ListFilesResponse) GetFirstID() (v string) {
	return p.FirstID
}

func (p *ListFilesResponse) GetLastID() (v string) {
	return p.LastID
}

func (p *ListFilesResponse) GetHasMore() (v bool) {
	return p.HasMore
}

func (p *ListFilesResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ListFilesResponse(%+v)", *p)
}

type DeleteFileRequest struct {
	ID string `thrift:"ID,1,required" json:"ID,required" path:"id,required"`
}

func NewDeleteFileRequest() *DeleteFileRequest {
	return &DeleteFileRequest{}
}

func (p *DeleteFileRequest) InitDefault() {
}

func (p *DeleteFileRequest) GetID() (v string) {
	return p.ID
}

func (p *DeleteFileRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteFileRequest(%+v)", *p)
}

type DeleteFileResponse struct {
	ObjectType ObjectType `thrift:"ObjectType,1,required" form:"object,required" json:"object,required"`
	ID         string     `thrift:"ID,2,required" form:"id,required" json:"id,required"`
	Deleted    bool       `thrift:"Deleted,3,required" form:"deleted,required" json:"deleted,required"`
	model.HttpHeader
}

func NewDeleteFileResponse() *DeleteFileResponse {
	return &DeleteFileResponse{}
}

func (p *DeleteFileResponse) InitDefault() {
}

func (p *DeleteFileResponse) GetObjectType() (v ObjectType) {
	return p.ObjectType
}

func (p *DeleteFileResponse) GetID() (v string) {
	return p.ID
}

func (p *DeleteFileResponse) GetDeleted() (v bool) {
	return p.Deleted
}

func (p *DeleteFileResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteFileResponse(%+v)", *p)
}
