// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package dnsiface provides an interface to enable mocking the DNS service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package dns

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// DNSAPI provides an interface to enable mocking the
// dns.DNS service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // DNS.
//    func myFunc(svc DNSAPI) bool {
//        // Make svc.BatchDeleteCustomLine request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := dns.New(sess)
//
//        myFunc(svc)
//    }
//
type DNSAPI interface {
	BatchDeleteCustomLineCommon(*map[string]interface{}) (*map[string]interface{}, error)
	BatchDeleteCustomLineCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	BatchDeleteCustomLineCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	BatchDeleteCustomLine(*BatchDeleteCustomLineInput) (*BatchDeleteCustomLineOutput, error)
	BatchDeleteCustomLineWithContext(volcengine.Context, *BatchDeleteCustomLineInput, ...request.Option) (*BatchDeleteCustomLineOutput, error)
	BatchDeleteCustomLineRequest(*BatchDeleteCustomLineInput) (*request.Request, *BatchDeleteCustomLineOutput)

	CheckRetrieveZoneCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CheckRetrieveZoneCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CheckRetrieveZoneCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CheckRetrieveZone(*CheckRetrieveZoneInput) (*CheckRetrieveZoneOutput, error)
	CheckRetrieveZoneWithContext(volcengine.Context, *CheckRetrieveZoneInput, ...request.Option) (*CheckRetrieveZoneOutput, error)
	CheckRetrieveZoneRequest(*CheckRetrieveZoneInput) (*request.Request, *CheckRetrieveZoneOutput)

	CheckZoneCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CheckZoneCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CheckZoneCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CheckZone(*CheckZoneInput) (*CheckZoneOutput, error)
	CheckZoneWithContext(volcengine.Context, *CheckZoneInput, ...request.Option) (*CheckZoneOutput, error)
	CheckZoneRequest(*CheckZoneInput) (*request.Request, *CheckZoneOutput)

	CreateCustomLineCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCustomLineCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCustomLineCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateCustomLine(*CreateCustomLineInput) (*CreateCustomLineOutput, error)
	CreateCustomLineWithContext(volcengine.Context, *CreateCustomLineInput, ...request.Option) (*CreateCustomLineOutput, error)
	CreateCustomLineRequest(*CreateCustomLineInput) (*request.Request, *CreateCustomLineOutput)

	CreateRecordCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateRecordCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateRecordCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateRecord(*CreateRecordInput) (*CreateRecordOutput, error)
	CreateRecordWithContext(volcengine.Context, *CreateRecordInput, ...request.Option) (*CreateRecordOutput, error)
	CreateRecordRequest(*CreateRecordInput) (*request.Request, *CreateRecordOutput)

	CreateUserZoneBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateUserZoneBackupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateUserZoneBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateUserZoneBackup(*CreateUserZoneBackupInput) (*CreateUserZoneBackupOutput, error)
	CreateUserZoneBackupWithContext(volcengine.Context, *CreateUserZoneBackupInput, ...request.Option) (*CreateUserZoneBackupOutput, error)
	CreateUserZoneBackupRequest(*CreateUserZoneBackupInput) (*request.Request, *CreateUserZoneBackupOutput)

	CreateZoneCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateZoneCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateZoneCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateZone(*CreateZoneInput) (*CreateZoneOutput, error)
	CreateZoneWithContext(volcengine.Context, *CreateZoneInput, ...request.Option) (*CreateZoneOutput, error)
	CreateZoneRequest(*CreateZoneInput) (*request.Request, *CreateZoneOutput)

	DeleteRecordCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRecordCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRecordCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRecord(*DeleteRecordInput) (*DeleteRecordOutput, error)
	DeleteRecordWithContext(volcengine.Context, *DeleteRecordInput, ...request.Option) (*DeleteRecordOutput, error)
	DeleteRecordRequest(*DeleteRecordInput) (*request.Request, *DeleteRecordOutput)

	DeleteUserZoneBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteUserZoneBackupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteUserZoneBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteUserZoneBackup(*DeleteUserZoneBackupInput) (*DeleteUserZoneBackupOutput, error)
	DeleteUserZoneBackupWithContext(volcengine.Context, *DeleteUserZoneBackupInput, ...request.Option) (*DeleteUserZoneBackupOutput, error)
	DeleteUserZoneBackupRequest(*DeleteUserZoneBackupInput) (*request.Request, *DeleteUserZoneBackupOutput)

	DeleteZoneCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteZoneCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteZoneCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteZone(*DeleteZoneInput) (*DeleteZoneOutput, error)
	DeleteZoneWithContext(volcengine.Context, *DeleteZoneInput, ...request.Option) (*DeleteZoneOutput, error)
	DeleteZoneRequest(*DeleteZoneInput) (*request.Request, *DeleteZoneOutput)

	ListCustomLinesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListCustomLinesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListCustomLinesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListCustomLines(*ListCustomLinesInput) (*ListCustomLinesOutput, error)
	ListCustomLinesWithContext(volcengine.Context, *ListCustomLinesInput, ...request.Option) (*ListCustomLinesOutput, error)
	ListCustomLinesRequest(*ListCustomLinesInput) (*request.Request, *ListCustomLinesOutput)

	ListDomainStatisticsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListDomainStatisticsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListDomainStatisticsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListDomainStatistics(*ListDomainStatisticsInput) (*ListDomainStatisticsOutput, error)
	ListDomainStatisticsWithContext(volcengine.Context, *ListDomainStatisticsInput, ...request.Option) (*ListDomainStatisticsOutput, error)
	ListDomainStatisticsRequest(*ListDomainStatisticsInput) (*request.Request, *ListDomainStatisticsOutput)

	ListLinesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListLinesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListLinesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListLines(*ListLinesInput) (*ListLinesOutput, error)
	ListLinesWithContext(volcengine.Context, *ListLinesInput, ...request.Option) (*ListLinesOutput, error)
	ListLinesRequest(*ListLinesInput) (*request.Request, *ListLinesOutput)

	ListRecordDigestByLineCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListRecordDigestByLineCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListRecordDigestByLineCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListRecordDigestByLine(*ListRecordDigestByLineInput) (*ListRecordDigestByLineOutput, error)
	ListRecordDigestByLineWithContext(volcengine.Context, *ListRecordDigestByLineInput, ...request.Option) (*ListRecordDigestByLineOutput, error)
	ListRecordDigestByLineRequest(*ListRecordDigestByLineInput) (*request.Request, *ListRecordDigestByLineOutput)

	ListRecordSetsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListRecordSetsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListRecordSetsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListRecordSets(*ListRecordSetsInput) (*ListRecordSetsOutput, error)
	ListRecordSetsWithContext(volcengine.Context, *ListRecordSetsInput, ...request.Option) (*ListRecordSetsOutput, error)
	ListRecordSetsRequest(*ListRecordSetsInput) (*request.Request, *ListRecordSetsOutput)

	ListRecordsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListRecordsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListRecordsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListRecords(*ListRecordsInput) (*ListRecordsOutput, error)
	ListRecordsWithContext(volcengine.Context, *ListRecordsInput, ...request.Option) (*ListRecordsOutput, error)
	ListRecordsRequest(*ListRecordsInput) (*request.Request, *ListRecordsOutput)

	ListUserZoneBackupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListUserZoneBackupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListUserZoneBackupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListUserZoneBackups(*ListUserZoneBackupsInput) (*ListUserZoneBackupsOutput, error)
	ListUserZoneBackupsWithContext(volcengine.Context, *ListUserZoneBackupsInput, ...request.Option) (*ListUserZoneBackupsOutput, error)
	ListUserZoneBackupsRequest(*ListUserZoneBackupsInput) (*request.Request, *ListUserZoneBackupsOutput)

	ListZoneStatisticsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListZoneStatisticsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListZoneStatisticsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListZoneStatistics(*ListZoneStatisticsInput) (*ListZoneStatisticsOutput, error)
	ListZoneStatisticsWithContext(volcengine.Context, *ListZoneStatisticsInput, ...request.Option) (*ListZoneStatisticsOutput, error)
	ListZoneStatisticsRequest(*ListZoneStatisticsInput) (*request.Request, *ListZoneStatisticsOutput)

	ListZonesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListZonesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListZonesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListZones(*ListZonesInput) (*ListZonesOutput, error)
	ListZonesWithContext(volcengine.Context, *ListZonesInput, ...request.Option) (*ListZonesOutput, error)
	ListZonesRequest(*ListZonesInput) (*request.Request, *ListZonesOutput)

	QueryBackupScheduleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	QueryBackupScheduleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	QueryBackupScheduleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	QueryBackupSchedule(*QueryBackupScheduleInput) (*QueryBackupScheduleOutput, error)
	QueryBackupScheduleWithContext(volcengine.Context, *QueryBackupScheduleInput, ...request.Option) (*QueryBackupScheduleOutput, error)
	QueryBackupScheduleRequest(*QueryBackupScheduleInput) (*request.Request, *QueryBackupScheduleOutput)

	QueryRecordCommon(*map[string]interface{}) (*map[string]interface{}, error)
	QueryRecordCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	QueryRecordCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	QueryRecord(*QueryRecordInput) (*QueryRecordOutput, error)
	QueryRecordWithContext(volcengine.Context, *QueryRecordInput, ...request.Option) (*QueryRecordOutput, error)
	QueryRecordRequest(*QueryRecordInput) (*request.Request, *QueryRecordOutput)

	QueryZoneCommon(*map[string]interface{}) (*map[string]interface{}, error)
	QueryZoneCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	QueryZoneCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	QueryZone(*QueryZoneInput) (*QueryZoneOutput, error)
	QueryZoneWithContext(volcengine.Context, *QueryZoneInput, ...request.Option) (*QueryZoneOutput, error)
	QueryZoneRequest(*QueryZoneInput) (*request.Request, *QueryZoneOutput)

	RestoreUserZoneBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestoreUserZoneBackupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestoreUserZoneBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestoreUserZoneBackup(*RestoreUserZoneBackupInput) (*RestoreUserZoneBackupOutput, error)
	RestoreUserZoneBackupWithContext(volcengine.Context, *RestoreUserZoneBackupInput, ...request.Option) (*RestoreUserZoneBackupOutput, error)
	RestoreUserZoneBackupRequest(*RestoreUserZoneBackupInput) (*request.Request, *RestoreUserZoneBackupOutput)

	RetrieveZoneCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RetrieveZoneCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RetrieveZoneCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RetrieveZone(*RetrieveZoneInput) (*RetrieveZoneOutput, error)
	RetrieveZoneWithContext(volcengine.Context, *RetrieveZoneInput, ...request.Option) (*RetrieveZoneOutput, error)
	RetrieveZoneRequest(*RetrieveZoneInput) (*request.Request, *RetrieveZoneOutput)

	UpdateBackupScheduleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateBackupScheduleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateBackupScheduleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateBackupSchedule(*UpdateBackupScheduleInput) (*UpdateBackupScheduleOutput, error)
	UpdateBackupScheduleWithContext(volcengine.Context, *UpdateBackupScheduleInput, ...request.Option) (*UpdateBackupScheduleOutput, error)
	UpdateBackupScheduleRequest(*UpdateBackupScheduleInput) (*request.Request, *UpdateBackupScheduleOutput)

	UpdateCustomLineCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateCustomLineCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateCustomLineCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateCustomLine(*UpdateCustomLineInput) (*UpdateCustomLineOutput, error)
	UpdateCustomLineWithContext(volcengine.Context, *UpdateCustomLineInput, ...request.Option) (*UpdateCustomLineOutput, error)
	UpdateCustomLineRequest(*UpdateCustomLineInput) (*request.Request, *UpdateCustomLineOutput)

	UpdateRecordCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateRecordCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateRecordCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateRecord(*UpdateRecordInput) (*UpdateRecordOutput, error)
	UpdateRecordWithContext(volcengine.Context, *UpdateRecordInput, ...request.Option) (*UpdateRecordOutput, error)
	UpdateRecordRequest(*UpdateRecordInput) (*request.Request, *UpdateRecordOutput)

	UpdateRecordSetCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateRecordSetCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateRecordSetCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateRecordSet(*UpdateRecordSetInput) (*UpdateRecordSetOutput, error)
	UpdateRecordSetWithContext(volcengine.Context, *UpdateRecordSetInput, ...request.Option) (*UpdateRecordSetOutput, error)
	UpdateRecordSetRequest(*UpdateRecordSetInput) (*request.Request, *UpdateRecordSetOutput)

	UpdateRecordStatusCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateRecordStatusCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateRecordStatusCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateRecordStatus(*UpdateRecordStatusInput) (*UpdateRecordStatusOutput, error)
	UpdateRecordStatusWithContext(volcengine.Context, *UpdateRecordStatusInput, ...request.Option) (*UpdateRecordStatusOutput, error)
	UpdateRecordStatusRequest(*UpdateRecordStatusInput) (*request.Request, *UpdateRecordStatusOutput)

	UpdateZoneCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateZoneCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateZoneCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateZone(*UpdateZoneInput) (*UpdateZoneOutput, error)
	UpdateZoneWithContext(volcengine.Context, *UpdateZoneInput, ...request.Option) (*UpdateZoneOutput, error)
	UpdateZoneRequest(*UpdateZoneInput) (*request.Request, *UpdateZoneOutput)
}

var _ DNSAPI = (*DNS)(nil)