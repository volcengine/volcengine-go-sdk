// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package advdefence20230308iface provides an interface to enable mocking the ADVDEFENCE20230308 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package advdefence20230308

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// ADVDEFENCE20230308API provides an interface to enable mocking the
// advdefence20230308.ADVDEFENCE20230308 service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // ADVDEFENCE20230308.
//    func myFunc(svc ADVDEFENCE20230308API) bool {
//        // Make svc.AddWebDefCcRule request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := advdefence20230308.New(sess)
//
//        myFunc(svc)
//    }
//
type ADVDEFENCE20230308API interface {
	AddWebDefCcRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddWebDefCcRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddWebDefCcRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddWebDefCcRule(*AddWebDefCcRuleInput) (*AddWebDefCcRuleOutput, error)
	AddWebDefCcRuleWithContext(volcengine.Context, *AddWebDefCcRuleInput, ...request.Option) (*AddWebDefCcRuleOutput, error)
	AddWebDefCcRuleRequest(*AddWebDefCcRuleInput) (*request.Request, *AddWebDefCcRuleOutput)

	BatchAddFwdRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	BatchAddFwdRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	BatchAddFwdRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	BatchAddFwdRule(*BatchAddFwdRuleInput) (*BatchAddFwdRuleOutput, error)
	BatchAddFwdRuleWithContext(volcengine.Context, *BatchAddFwdRuleInput, ...request.Option) (*BatchAddFwdRuleOutput, error)
	BatchAddFwdRuleRequest(*BatchAddFwdRuleInput) (*request.Request, *BatchAddFwdRuleOutput)

	DelWebDefCcRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DelWebDefCcRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DelWebDefCcRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DelWebDefCcRule(*DelWebDefCcRuleInput) (*DelWebDefCcRuleOutput, error)
	DelWebDefCcRuleWithContext(volcengine.Context, *DelWebDefCcRuleInput, ...request.Option) (*DelWebDefCcRuleOutput, error)
	DelWebDefCcRuleRequest(*DelWebDefCcRuleInput) (*request.Request, *DelWebDefCcRuleOutput)

	DescribeAttackFlowCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAttackFlowCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAttackFlowCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAttackFlow(*DescribeAttackFlowInput) (*DescribeAttackFlowOutput, error)
	DescribeAttackFlowWithContext(volcengine.Context, *DescribeAttackFlowInput, ...request.Option) (*DescribeAttackFlowOutput, error)
	DescribeAttackFlowRequest(*DescribeAttackFlowInput) (*request.Request, *DescribeAttackFlowOutput)

	DescribeBizFlowAndConnCountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBizFlowAndConnCountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBizFlowAndConnCountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBizFlowAndConnCount(*DescribeBizFlowAndConnCountInput) (*DescribeBizFlowAndConnCountOutput, error)
	DescribeBizFlowAndConnCountWithContext(volcengine.Context, *DescribeBizFlowAndConnCountInput, ...request.Option) (*DescribeBizFlowAndConnCountOutput, error)
	DescribeBizFlowAndConnCountRequest(*DescribeBizFlowAndConnCountInput) (*request.Request, *DescribeBizFlowAndConnCountOutput)

	UpdWebDefCcRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdWebDefCcRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdWebDefCcRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdWebDefCcRule(*UpdWebDefCcRuleInput) (*UpdWebDefCcRuleOutput, error)
	UpdWebDefCcRuleWithContext(volcengine.Context, *UpdWebDefCcRuleInput, ...request.Option) (*UpdWebDefCcRuleOutput, error)
	UpdWebDefCcRuleRequest(*UpdWebDefCcRuleInput) (*request.Request, *UpdWebDefCcRuleOutput)

	UpdateFwdRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateFwdRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateFwdRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateFwdRule(*UpdateFwdRuleInput) (*UpdateFwdRuleOutput, error)
	UpdateFwdRuleWithContext(volcengine.Context, *UpdateFwdRuleInput, ...request.Option) (*UpdateFwdRuleOutput, error)
	UpdateFwdRuleRequest(*UpdateFwdRuleInput) (*request.Request, *UpdateFwdRuleOutput)
}

var _ ADVDEFENCE20230308API = (*ADVDEFENCE20230308)(nil)