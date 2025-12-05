package lklsdk

import (
	"fmt"
	"time"

	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/model"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// ApplyBind 分账关系绑定申请
func (s *SplitLedgerService[T]) ApplyBind(req *model.ApplyBindReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_SPLIT_LEDGER_RECEIVE_BIND_URL

	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.ApplyBindRequest{
		ReqTime: time.Now().Format("20060102150405"),
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

func (s *SplitLedgerService[T]) ApplyBindTest(req *model.ApplyBindReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_TEST_URL + consts.LKL_SPLIT_LEDGER_RECEIVE_BIND_URL

	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.ApplyBindRequest{
		ReqTime: time.Now().Format("20060102150405"),
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

// QuerySplitBalance 可分账金额查询
func (s *SplitLedgerService[T]) QuerySplitBalance(req *model.SplitBalanceReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_SPLIT_LEDGER_BALANCE_URL

	// 构建BaseModel请求
	baseReq := model.SplitBalanceRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
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

func (s *SplitLedgerService[T]) QuerySplitBalanceTest(req *model.SplitBalanceReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_TEST_URL + consts.LKL_SPLIT_LEDGER_BALANCE_URL

	// 构建BaseModel请求
	baseReq := model.SplitBalanceRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
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

// OrderSplitLedger 订单分账
func (s *SplitLedgerService[T]) OrderSplitLedger(req *model.OrderSplitLedgerReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_ORDER_SPLIT_LEDGER_URL

	// 构建BaseModel请求
	baseReq := model.OrderSplitLedgerRequest{
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

func (s *SplitLedgerService[T]) OrderSplitLedgerTest(req *model.OrderSplitLedgerReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_TEST_URL + consts.LKL_ORDER_SPLIT_LEDGER_URL

	// 构建BaseModel请求
	baseReq := model.OrderSplitLedgerRequest{
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

func (s *SplitLedgerService[T]) OrderSplitLedgerResultQuery(req *model.OrderSplitLedgerResultQueryReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_SPLIT_LEDGER_RESULT_QUERY_URL

	// 构建BaseModel请求
	baseReq := model.OrderSplitLedgerResultQueryRequest{
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
