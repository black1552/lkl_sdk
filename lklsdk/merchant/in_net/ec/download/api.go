package download

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

type Download struct {
	Client *common.Client[ECDownloadResponse]
}

// NewDownload 创建电子合同下载服务实例
func NewDownload(client *common.Client[ECDownloadResponse]) *Download {
	return &Download{
		Client: client,
	}
}

// ECDownload 电子合同下载
// 提供已完成的签约电子合同下载
func (d *Download) ECDownload(req *ECDownloadRequestData) (*ECDownloadResponse, error) {
	// 构建请求
	baseReq := ECDownloadRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		ReqData: req,
		Version: "1.0",
	}

	// 使用生产环境URL
	return d.Client.DoRequest(consts.BASE_URL+consts.LKL_EC_DOWNLOAD, baseReq)
}

// ECDownloadTest 电子合同下载（测试环境）
// 提供已完成的签约电子合同下载
func (d *Download) ECDownloadTest(req *ECDownloadRequestData) (*ECDownloadResponse, error) {
	// 构建请求
	baseReq := ECDownloadRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		ReqData: req,
		Version: "1.0",
	}

	// 使用测试环境URL
	return d.Client.DoRequest(consts.BASE_TEST_URL+consts.LKL_EC_DOWNLOAD, baseReq)
}
