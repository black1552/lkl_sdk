package mergerefund

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

type MergeRefund[T any] struct {
	client *common.Client[T]
}

func NewMergeRefund[T any](client *common.Client[T]) *MergeRefund[T] {
	return &MergeRefund[T]{
		client: client,
	}
}
func (t *MergeRefund[T]) MergeRefund(req *RequestDataMergeRefund) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_UNIFIED_RETURN_MERGE_REFUND_URL
	// 构建BaseModel请求
	baseReq := RequestMergeRefund{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: req,
	}

	// 发送请求
	respBody, err := t.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
