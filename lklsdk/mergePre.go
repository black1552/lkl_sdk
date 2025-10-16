package lklsdk

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/black1552/lkl_sdk/model"
	"github.com/gogf/gf/v2/os/gtime"
)

type MergePreService[T any] struct {
	client *common.Client[T]
}

// NewMergePreService 创建拉卡拉主扫合单交易
func NewMergePreService[T any](client *common.Client[T]) *MergePreService[T] {
	return &MergePreService[T]{
		client: client,
	}
}

func (s *MergePreService[T]) PreOrder(req *model.MergePreorderReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_MERGE_ORDER_URL

	// 构建BaseModel请求
	baseReq := model.MergePreorder{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "2.0",
		ReqData: req,
	}

	// 发送请求
	respBody, err := s.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
