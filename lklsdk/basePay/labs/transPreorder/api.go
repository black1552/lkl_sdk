package transPreorder

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

// API 聚合主扫API结构体
type TransPreorderServer struct {
	client *common.Client[TransPreorderResponse]
}

// NewAPI 创建聚合主扫API实例
func NewTransPreorderServer(client *common.Client[TransPreorderResponse]) *TransPreorderServer {
	return &TransPreorderServer{
		client: client,
	}
}

// TransPreorder 执行聚合主扫交易
func (a *TransPreorderServer) TransPreorder(req *TransPreorderRequestData) (*TransPreorderResponse, error) {
	// 构建BaseModel请求
	baseReq := TransPreorderRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: req,
	}

	// 发送请求
	return a.client.DoRequest(consts.BASE_URL+consts.LKL_BASE_URL_TRANS_PREORDER_URL, baseReq)
}
