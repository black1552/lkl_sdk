package querystatus

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

type QStatus struct {
	Client *common.Client[ECQueryStatusResponse]
}

// NewQStatus 创建电子合同查询状态服务实例
func NewQStatus(client *common.Client[ECQueryStatusResponse]) *QStatus {
	return &QStatus{
		Client: client,
	}
}

// QueryStatus 电子合同查询状态
// 提供申请过与拉卡拉电子签约用户查询电子合同签署状态
func (q *QStatus) QueryStatus(req *ECQueryStatusRequestData) (*ECQueryStatusResponse, error) {
	// 构建请求
	baseReq := ECQueryStatusRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		ReqData: req,
		Version: "1.0",
	}

	// 使用生产环境URL
	return q.Client.DoRequest(consts.BASE_URL+consts.LKL_EC_QUERY_STATUS, baseReq)
}

// QueryStatusTest 电子合同查询状态（测试环境）
// 提供申请过与拉卡拉电子签约用户查询电子合同签署状态
func (q *QStatus) QueryStatusTest(req *ECQueryStatusRequestData) (*ECQueryStatusResponse, error) {
	// 构建请求
	baseReq := ECQueryStatusRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		ReqData: req,
		Version: "1.0",
	}

	// 使用测试环境URL
	return q.Client.DoRequest(consts.BASE_TEST_URL+consts.LKL_EC_QUERY_STATUS, baseReq)
}
