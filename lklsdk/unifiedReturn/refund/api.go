package refund

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

// Refund 统一退货API结构体
type Refund struct {
	client *common.Client[ResponseRefund]
}

// NewRefund 创建统一退货API实例
func NewRefund(client *common.Client[ResponseRefund]) *Refund {
	return &Refund{
		client: client,
	}
}

// Refund 发起统一退货请求
// request: 统一退货请求参数
// 返回统一退货响应结果和错误信息
func (api *Refund) Refund(req *RequestDataRefund) (*ResponseRefund, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_UNIFIED_RETURN_REFUND_URL
	// 构建BaseModel请求
	baseReq := RequestRefund{
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
