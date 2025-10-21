package transShareCode

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

// TransShareCodeServer 聚合扫码申请支付宝吱口令服务
type TransShareCodeServer struct {
	client *common.Client[TransShareCodeResponse]
}

// NewTransShareCodeServer 创建聚合扫码申请支付宝吱口令服务实例
func NewTransShareCodeServer(client *common.Client[TransShareCodeResponse]) *TransShareCodeServer {
	return &TransShareCodeServer{
		client: client,
	}
}

// TransShareCode 聚合扫码申请支付宝吱口令
func (s *TransShareCodeServer) TransShareCode(reqData *TransShareCodeRequestData) (*TransShareCodeResponse, error) {
	// 构建请求
	request := &TransShareCodeRequest{
		ReqTime: gtime.Now().Format("20060102150405"),
		Version: "3.0",
		ReqData: reqData,
	}
	return s.client.DoRequest(consts.BASE_URL+consts.LKL_BASE_URL_TRANS_SHARE_CODE_URL, request)
}
