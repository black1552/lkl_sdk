package lklsdk

import (
	"fmt"
	"time"

	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/black1552/lkl_sdk/model"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

// MerService 商户服务
type MerService[T any] struct {
	client *common.Client[T]
}

// NewMerService 创建交易服务实例
func NewMerService[T any](client *common.Client[T]) *MerService[T] {
	return &MerService[T]{
		client: client,
	}
}

// AddMer 商户进件
func (t *MerService[T]) AddMer(req *model.MerchantApplyReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_ADD_MER
	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.MerchantApplyRequest{
		ReqData:   req,
		ReqId:     md5,
		Timestamp: time.Now().Unix(),
		Ver:       "1.0.0",
	}

	// 发送请求
	respBody, err := t.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// AddMerTest 商户进件(测试)
func (t *MerService[T]) AddMerTest(req *model.MerchantApplyReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_TEST_URL + consts.LKL_ADD_MER
	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.MerchantApplyRequest{
		ReqData:   req,
		ReqId:     md5,
		Timestamp: time.Now().Unix(),
		Ver:       "1.0.0",
	}

	// 发送请求
	respBody, err := t.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func (t *MerService[T]) QueryMer(req *model.QueryMerRequestData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_QUERY_MER
	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.QueryMerRequest{
		ReqData:   req,
		ReqId:     md5,
		Timestamp: time.Now().Unix(),
		Ver:       "1.0.0",
	}

	// 发送请求
	respBody, err := t.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
func (t *MerService[T]) QueryMerTest(req *model.QueryMerRequestData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_TEST_URL + consts.LKL_QUERY_MER
	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.QueryMerRequest{
		ReqData:   req,
		ReqId:     md5,
		Timestamp: time.Now().Unix(),
		Ver:       "1.0.0",
	}

	// 发送请求
	respBody, err := t.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
func (t *MerService[T]) MerValidate(req *model.MerValidateRequestData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_MER_VALIDATE
	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.MerValidateRequest{
		ReqData:   req,
		ReqId:     md5,
		Timestamp: time.Now().Unix(),
		Ver:       "1.0.0",
	}

	// 发送请求
	respBody, err := t.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
func (t *MerService[T]) MerValidateTest(req *model.MerValidateRequestData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_TEST_URL + consts.LKL_MER_VALIDATE
	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.MerValidateRequest{
		ReqData:   req,
		ReqId:     md5,
		Timestamp: time.Now().Unix(),
		Ver:       "1.0.0",
	}

	// 发送请求
	respBody, err := t.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func (t *MerService[T]) ReconsiderSubmitTest(req *model.ReConfSubmitRequestData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_TEST_URL + consts.LKL_RECONF_SUBMIT
	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.ReConfSubmitRequest{
		ReqData:   req,
		ReqId:     md5,
		Timestamp: time.Now().Unix(),
		Ver:       "1.0",
	}

	// 发送请求
	respBody, err := t.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
