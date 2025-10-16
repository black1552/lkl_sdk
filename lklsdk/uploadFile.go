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

// UploadFileService 交易服务
type UploadFileService[T any] struct {
	client *common.Client[T]
}

// NewUploadFileService 创建交易服务实例
func NewUploadFileService[T any](client *common.Client[T]) *UploadFileService[T] {
	return &UploadFileService[T]{
		client: client,
	}
}

// UploadFileQuery 交易查询
func (t *UploadFileService[T]) UploadFileQuery(req *model.UploadFileReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_UPLOAD_FILE_URL
	md5, err := gmd5.Encrypt(gconv.String(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("创建ReqId失败")
	}
	// 构建BaseModel请求
	baseReq := model.UploadFileRequest{
		Timestamp: gconv.String(time.Now().Unix()),
		Ver:       "1.0",
		ReqId:     md5,
		ReqData:   req,
	}

	// 发送请求
	respBody, err := t.client.DoRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
