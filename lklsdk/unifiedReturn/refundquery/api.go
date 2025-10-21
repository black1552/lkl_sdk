package refundquery

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

type RefundQuery struct {
	client *common.Client[ResponseRefundQuery]
}

// NewRefundQuery 创建统一退货查询API实例
func NewRefundQuery(client *common.Client[ResponseRefundQuery]) *RefundQuery {
	return &RefundQuery{
		client: client,
	}
}

// RefundQuery 发起统一退货查询请求
func (api *RefundQuery) RefundQuery(req *RequestDataRefundQuery) (*ResponseRefundQuery, error) {
	// 构建BaseModel请求
	baseReq := RequestRefundQuery{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: req,
	}

	return api.client.DoRequest(consts.BASE_URL+consts.LKL_UNIFIED_RETURN_REFUND_QUERY_URL, baseReq)
}
