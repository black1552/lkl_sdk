package qmastatus

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

// QmaStatus 电子合同人工复核结果查询服务
type QmaStatus struct {
	Client *common.Client[ECQmaStatusResponse]
}

// NewQmaStatus 创建新的电子合同人工复核结果查询服务
func NewQmaStatus(client *common.Client[ECQmaStatusResponse]) *QmaStatus {
	return &QmaStatus{Client: client}
}

// ECQmaStatus
// https://o.lakala.com/#/home/document/detail?id=982
func (q *QmaStatus) ECQmaStatus(req *ECQmaStatusRequestData) (*ECQmaStatusResponse, error) {
	// 创建请求对象
	request := ECQmaStatusRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		ReqData: req,
		Version: "1.0",
	}
	return q.Client.DoRequest(consts.BASE_URL+consts.LKL_EC_QMA_STATUS, request)
}

// ECQmaStatusTest 电子合同人工复核结果查询（测试环境）
// https://o.lakala.com/#/home/document/detail?id=982
func (q *QmaStatus) ECQmaStatusTest(req *ECQmaStatusRequestData) (*ECQmaStatusResponse, error) {
	// 创建请求对象
	request := ECQmaStatusRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		ReqData: req,
		Version: "1.0",
	}
	return q.Client.DoRequest(consts.BASE_TEST_URL+consts.LKL_EC_QMA_STATUS, request)
}
