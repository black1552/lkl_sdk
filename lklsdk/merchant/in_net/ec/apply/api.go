package apply

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

type Apply struct {
	Client *common.Client[ECApplyResponse]
}

func NewEcApply(client *common.Client[ECApplyResponse]) *Apply {
	return &Apply{
		Client: client,
	}
}

// ECApply 电子合同申请
// 提供与拉卡拉进行电子签约的第四方进行电子合同申请
// 电子合同签约成功后不需要将其下载出来作为附件上传，只需将电子合同编号（ecNo）
// 在“新增商户入网”接口中在（contractNo）字段中传入即可
func (a *Apply) ECApply(req *ECApplyRequestData) (*ECApplyResponse, error) {
	// 构建请求
	baseReq := ECApplyRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		ReqData: req,
		Version: "1.0",
	}

	return a.Client.DoRequest(consts.BASE_URL+consts.LKL_EC_APPLY, baseReq)
}

// ECApplyTest 电子合同申请（测试环境）
// 提供与拉卡拉进行电子签约的第四方进行电子合同申请
// 电子合同签约成功后不需要将其下载出来作为附件上传，只需将电子合同编号（ecNo）
// 在“新增商户入网”接口中在（contractNo）字段中传入即可
func (a *Apply) ECApplyTest(req *ECApplyRequestData) (*ECApplyResponse, error) {
	// 构建请求
	baseReq := ECApplyRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		ReqData: req,
		Version: "1.0",
	}

	return a.Client.DoRequest(consts.BASE_TEST_URL+consts.LKL_EC_APPLY, baseReq)
}
