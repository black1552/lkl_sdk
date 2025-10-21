package transmicropay

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

// API 聚合被扫API结构体
type TransMicropayServer struct {
	Client *common.Client[TransMicropayResponse]
}

// NewAPI 创建聚合被扫API实例
func NewTransMicropayServer(client *common.Client[TransMicropayResponse]) *TransMicropayServer {
	return &TransMicropayServer{
		Client: client,
	}
}

// TransMicropay 执行聚合被扫请求
// - merchantNo: 商户号
// - termNo: 终端号
// - outTradeNo: 商户交易流水号
// - authCode: 支付授权码
// - totalAmount: 金额
// - locationInfo: 地址位置信息
// - 其他可选参数通过TransMicropayRequestData的其他字段设置
func (a *TransMicropayServer) TransMicropay(reqData *TransMicropayRequestData) (*TransMicropayResponse, error) {
	// 构建请求结构体
	request := &TransMicropayRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: reqData,
	}

	// 发送请求 - 使用consts中定义的聚合被扫URL
	return a.Client.DoRequest(consts.BASE_URL+consts.LKL_BASE_URL_TRANS_MICROPAY_URL, request)
}
