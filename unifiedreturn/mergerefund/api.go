package mergerefund

import (
	"github.com/black1552/lkl_sdk/common"
	"github.com/black1552/lkl_sdk/consts"
	"github.com/gogf/gf/v2/os/gtime"
)

type MergeRefund struct {
	client *common.Client[ResponseMergeRefund]
}

func NewMergeRefund(client *common.Client[ResponseMergeRefund]) *MergeRefund {
	return &MergeRefund{
		client: client,
	}
}
func (t *MergeRefund) MergeRefund(req *RequestDataMergeRefund) (*ResponseMergeRefund, error) {
	// 构建BaseModel请求
	baseReq := RequestMergeRefund{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: req,
	}
	return t.client.DoRequest(consts.BASE_URL+consts.LKL_UNIFIED_RETURN_MERGE_REFUND_URL, baseReq)
}
func (t *MergeRefund) MergeRefundTest(req *RequestDataMergeRefund) (*ResponseMergeRefund, error) {
	// 构建BaseModel请求
	baseReq := RequestMergeRefund{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: req,
	}
	return t.client.DoRequest(consts.BASE_TEST_URL+consts.LKL_UNIFIED_RETURN_MERGE_REFUND_URL, baseReq)
}
