package transMicropayEncry

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

// TransMicropayEncryServer 聚合扫码被扫接口（全报文加密）服务结构体
type TransMicropayEncryServer[T any] struct {
	Client *common.Client[T]
}

// NewTransMicropayEncryServer 创建聚合扫码被扫接口（全报文加密）服务实例
func NewTransMicropayEncryServer(client *common.Client[TransMicropayEncryResponse]) *TransMicropayEncryServer[TransMicropayEncryResponse] {
	return &TransMicropayEncryServer[TransMicropayEncryResponse]{
		Client: client,
	}
}

// TransMicropayEncry 聚合扫码被扫接口（全报文加密）接口
func (a *TransMicropayEncryServer[TransMicropayEncryResponse]) TransMicropayEncry(reqData *TransMicropayEncryRequestData) (*TransMicropayEncryResponse, error) {
	// 构建请求体
	request := &TransMicropayEncryRequest{
		ReqTime: gtime.Now().Format("20060102150405"),
		Version: "3.0",
		ReqData: reqData,
	}

	// 调用Client发送请求
	return a.Client.DoRequest(consts.BASE_URL+consts.LKL_BASE_URL_TRANS_MICROPAY_ENCRY_URL, request)
}
