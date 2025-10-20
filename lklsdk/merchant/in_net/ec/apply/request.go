package apply

import (
	"github.com/black1552/lkl_sdk/consts"
)

// ECApplyRequestData 电子合同申请请求数据
// 提供与拉卡拉进行电子签约的第四方进行电子合同申请
// 电子合同签约成功后不需要将其下载出来作为附件上传，只需将电子合同编号（ecNo）
// 在“新增商户入网”接口中在（contractNo）字段中传入即可
// 合同类别适用业务场景说明：
// EC001: 特约商户支付服务合作协议V3.1(商户入网) - 历史存量，推荐使用EC008
// EC002: 特约商户支付服务合作协议V3.2+分账结算授权委托书 - 历史存量，推荐使用EC007
// EC003: 分账结算授权委托书 - 历史存量，推荐使用EC009
// EC004: 特约商户支付服务合作协议V3.3(商户入网) - 历史存量，推荐使用EC008
// EC005: 特约商户支付服务合作协议V3.3+分账结算授权委托书 - 历史存量，推荐使用EC007
// EC007: 特约商户支付服务合作协议V4.1+分账结算授权委托书(商户入网+分账业务) - 当前最新版本
// EC008: 特约商户支付服务合作协议V4.1(商户入网) - 当前最新版本
// EC009: 清分结算授权委托书(分账业务) - 当前最新版本
type ECApplyRequestData struct {
	OrderNo             string              `json:"order_no" dc:"四方机构自定义订单编号 必选，建议：平台编号+14位年月日时分秒+8位的随机数"`                                                          // 必选，建议：平台编号+14位年月日时分秒+8位的随机数
	OrgID               int                 `json:"org_id" dc:"机构号 必选，签约方所属拉卡拉机构"`                                                                                  // 必选，签约方所属拉卡拉机构
	EcTypeCode          consts.ECType       `json:"ec_type_code" dc:"合同类别 必选，合同类别编码"`                                                                               // 必选，合同类别编码
	CertType            consts.CertType     `json:"cert_type" dc:"法人/经营者证件类型 必选，RESIDENT_ID（身份证）；PASSPORT（护照）；HK_MACAO_PASS（港澳居民往来内地通行证）；TAIWAN_PASS（台湾居民来往大陆通行证）"` // 必选，RESIDENT_ID（身份证）；PASSPORT（护照）；HK_MACAO_PASS（港澳居民往来内地通行证）；TAIWAN_PASS（台湾居民来往大陆通行证）
	CertName            string              `json:"cert_name" dc:"法人/经营者姓名 必选"`                                                                                     // 必选
	CertNo              string              `json:"cert_no" dc:"法人/经营者证件号码 必选"`                                                                                     // 必选
	Mobile              string              `json:"mobile" dc:"签约手机号 必选，小微个人商户必须填写经营者本人手机号；个体工商户或企业商户必须填写法人手机号或者经办人手机号"`                                            // 必选，小微个人商户必须填写经营者本人手机号；个体工商户或企业商户必须填写法人手机号或者经办人手机号
	BusinessLicenseNo   string              `json:"business_license_no" dc:"营业执照号 可选，个体工商户或企业商户必传"`                                                                 // 可选，个体工商户或企业商户必传
	BusinessLicenseName string              `json:"business_license_name" dc:"营业执照名称 可选，个体工商户或企业商户必传"`                                                              // 可选，个体工商户或企业商户必传
	OpenningBankCode    string              `json:"openning_bank_code" dc:"企业/经营者结算开户行号 必选"`                                                                        // 必选
	OpenningBankName    string              `json:"openning_bank_name" dc:"企业/经营者结算开户行名称 必选"`                                                                       // 必选
	AcctTypeCode        consts.AcctTypeCode `json:"acct_type_code" dc:"企业/经营者结算卡性质 必选，57 对公、 58 对私"`                                                                // 必选，57 对公、 58 对私
	AcctNo              string              `json:"acct_no" dc:"企业/经营者结算卡号 必选"`                                                                                     // 必选
	AcctName            string              `json:"acct_name" dc:"企业/经营者结算卡名称 必选"`                                                                                  // 必选
	EcContentParameters string              `json:"ec_content_parameters" dc:"电子合同内容参数集合 必选，JSON字符串"`                                                               // 必选，JSON字符串
	AgentTag            consts.AgentTag     `json:"agent_tag" dc:"是否经办签约 可选，0 不启用 1启用，缺省 0"`                                                                        // 可选，0 不启用 1启用，缺省 0
	AgentName           string              `json:"agent_name" dc:"经办人名称 可选，agentTag为1时必传"`                                                                         // 可选，agentTag为1时必传
	AgentCertType       consts.CertType     `json:"agent_cert_type" dc:"经办人证件类型 可选，agentTag为1时必传"`                                                                  // 可选，agentTag为1时必传
	AgentCertNo         string              `json:"agent_cert_no" dc:"经办人证件号 可选，agentTag为1时必传"`                                                                     // 可选，agentTag为1时必传
	AgentFileName       string              `json:"agent_file_name" dc:"经办签约授权委托书文件名 可选，agentTag为1时必传"`                                                             // 可选，agentTag为1时必传
	AgentFilePath       string              `json:"agent_file_path" dc:"经办授权委托书文件路径 可选，agentTag为1时必传"`                                                              // 可选，agentTag为1时必传
	Remark              string              `json:"remark" dc:"备注说明 可选"`                                                                                            // 可选
	RetUrl              string              `json:"ret_url" dc:"电子合同签约结果回调通知 可选，成功签约才通知"`                                                                           // 可选，成功签约才通知
}

// ECApplyRequest 电子合同申请请求
// 请求URL：
// 测试环境：https://test.wsmsd.cn/sit/api/v3/mms/open_api/ec/apply
// 生产环境：https://s2.lakala.com/api/v3/mms/open_api/ec/apply
type ECApplyRequest struct {
	ReqTime string              `json:"req_time" dc:"请求时间"` // 必选，格式：yyyyMMddHHmmss
	ReqData *ECApplyRequestData `json:"req_data" dc:"请求数据"` // 请求数据
	Version string              `json:"version" dc:"版本号"`   // 必选，固定值：1.0
}
