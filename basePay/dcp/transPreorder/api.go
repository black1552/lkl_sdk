package transpreorder

import (
	"github.com/black1552/lkl_sdk/common"
	"github.com/black1552/lkl_sdk/consts"
	"github.com/gogf/gf/v2/os/gtime"
)

type TransPreorder struct {
	client *common.Client[TransPreorderResponse]
}

func NewTransPreorderApi(client *common.Client[TransPreorderResponse]) *TransPreorder {
	return &TransPreorder{client: client}
}

func (tp *TransPreorder) TransPreorder(reqData *TransPreorderRequestData) (*TransPreorderResponse, error) {
	// 构建请求结构体
	request := &TransPreorderRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: reqData,
	}
	return tp.client.DoRequest(consts.BASE_URL+consts.LKL_DCP_TRANS_PREORDER_URL, request)
}
func (tp *TransPreorder) TransPreorderTest(reqData *TransPreorderRequestData) (*TransPreorderResponse, error) {
	// 构建请求结构体
	request := &TransPreorderRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: reqData,
	}
	return tp.client.DoRequest(consts.BASE_TEST_URL+consts.LKL_DCP_TRANS_PREORDER_URL, request)
}
