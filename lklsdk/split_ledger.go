package lklsdk

import (
	"fmt"
	"time"

	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/black1552/lkl_sdk/model"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// SplitLedgerService 分账服务
type SplitLedgerService[T any] struct {
	client *common.Client[T]
}

// NewSplitLedgerService 创建分账服务实例
func NewSplitLedgerService[T any](client *common.Client[T]) *SplitLedgerService[T] {
	return &SplitLedgerService[T]{
		client: client,
	}
}

// ApplyLedgerMer 商户分账业务开通申请
func (s *SplitLedgerService[T]) ApplyLedgerMer(req *model.ApplyLedgerMerReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_SPLIT_LEDGER_URL

	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.ApplyLedgerMerRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "2.0",
		ReqData: req,
		ReqId:   md5,
	}

	// 发送请求
	respBody, err := s.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

// ApplyLedgerReceiver 分账接收方创建申请
func (s *SplitLedgerService[T]) ApplyLedgerReceiver(req *model.ApplyLedgerReceiverReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_SPLIT_LEDGER_RECEIVE_URL

	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.ApplyLedgerReceiverRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "2.0",
		ReqId:   md5,
		ReqData: req,
	}

	// 发送请求
	respBody, err := s.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

// QueryLedgerMer 商户分账信息查询
func (s *SplitLedgerService[T]) QueryLedgerMer(req *model.QueryLedgerMerReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_SPLIT_LEDGER_QUERY_URL

	// 构建BaseModel请求
	baseReq := model.QueryLedgerMerRequest{
		ReqTime: time.Now().Format("20060102150405"),
		Version: "3.0",
		ReqData: req,
	}

	// 发送请求
	respBody, err := s.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
