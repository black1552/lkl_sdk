package applymanual

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

// ApplyManual 电子合同人工复核申请服务
type ApplyManual struct {
	Client *common.Client[ECApplyManualResponse]
}

// NewApplyManual 创建新的电子合同人工复核申请服务
func NewApplyManual(client *common.Client[ECApplyManualResponse]) *ApplyManual {
	return &ApplyManual{
		Client: client,
	}
}

// ECApplyManual 电子合同人工复核申请（生产环境）
// https://o.lakala.com/#/home/document/detail?id=981
func (a *ApplyManual) ECApplyManual(req *ECApplyManualRequestData) (*ECApplyManualResponse, error) {
	// 创建请求对象
	request := ECApplyManualRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		ReqData: req,
		Version: "1.0",
	}
	return a.Client.DoRequest(consts.BASE_URL+consts.LKL_EC_APPLY_MANUAL, request)
}

// ECApplyManualTest 电子合同人工复核申请（测试环境）
// https://o.lakala.com/#/home/document/detail?id=981
func (a *ApplyManual) ECApplyManualTest(req *ECApplyManualRequestData) (*ECApplyManualResponse, error) {
	// 创建请求对象
	request := ECApplyManualRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		ReqData: req,
		Version: "1.0",
	}
	return a.Client.DoRequest(consts.BASE_TEST_URL+consts.LKL_EC_APPLY_MANUAL, request)
}
