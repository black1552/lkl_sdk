package transquery

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

// API 聚合扫码交易查询API结构体
type API struct {
	Client *common.Client[TransQueryResponse]
}

// NewAPI 创建聚合扫码交易查询API实例
func NewAPI(client *common.Client[TransQueryResponse]) *API {
	return &API{
		Client: client,
	}
}

// TransQuery 执行聚合扫码交易查询请求
// - merchantNo: 商户号
// - termNo: 终端号
// - outTradeNo: 商户交易流水号（可选，与tradeNo二选一）
// - tradeNo: 拉卡拉交易流水号（可选，与outTradeNo二选一）
func (a *API) TransQuery(reqData *TransQueryRequestData) (*TransQueryResponse, error) {
	// 构建请求结构体
	request := &TransQueryRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: reqData,
	}

	// 发送请求 - 使用consts中定义的聚合扫码交易查询URL
	return a.Client.DoRequest(consts.BASE_URL+consts.LKL_BASE_URL_TRANS_QUERY_URL, request)
}
