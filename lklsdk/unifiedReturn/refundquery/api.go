package refundquery

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

type RefundQuery[T any] struct {
	client *common.Client[T]
}

// NewRefundQuery 创建统一退货查询API实例
func NewRefundQuery[T any](client *common.Client[T]) *RefundQuery[T] {
	return &RefundQuery[T]{
		client: client,
	}
}

// RefundQuery 发起统一退货查询请求
func (api *RefundQuery[T]) RefundQuery(req *RequestDataRefundQuery) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_UNIFIED_RETURN_REFUND_QUERY_URL
	// 构建BaseModel请求
	baseReq := RequestRefundQuery{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: req,
	}

	// 发送请求
	respBody, err := api.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
