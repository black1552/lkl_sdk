package transPreorederEncry

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

// TransPreorderEncryServer 聚合扫码主扫交易（全报文加密）服务结构体
type TransPreorderEncryServer[T any] struct {
	Client *common.Client[T]
}

// NewTransPreorderEncryServer 创建聚合扫码主扫交易（全报文加密）服务实例
func NewTransPreorderEncryServer(client *common.Client[TransPreorederEncryResponse]) *TransPreorderEncryServer[TransPreorederEncryResponse] {
	return &TransPreorderEncryServer[TransPreorederEncryResponse]{
		Client: client,
	}
}

// TransPreorderEncry 聚合扫码主扫交易（全报文加密）接口
func (a *TransPreorderEncryServer[TransPreorederEncryResponse]) TransPreorderEncry(reqData *TransPreorederEncryRequestData) (*TransPreorederEncryResponse, error) {
	// 构建请求体
	request := &TransPreorederEncryRequest{
		ReqTime: gtime.Now().Format("20060102150405"),
		Version: "3.0",
		ReqData: reqData,
	}

	// 调用Client发送请求
	return a.Client.DoRequest(consts.BASE_URL+consts.LKL_BASE_URL_TRANS_PREORDER_ENCRY_URL, request)
}
