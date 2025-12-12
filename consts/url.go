package consts

const (
	BASE_URL      = "https://s2.lakala.com/api"
	BASE_TEST_URL = "https://test.wsmsd.cn/sit/api"
)

const (
	// LKL_UPLOAD_FILE_URL 拉卡拉附件上传地址
	LKL_UPLOAD_FILE_URL = "/v2/mms/openApi/uploadFile"
	// LKL_SPLIT_LEDGER_URL 拉卡拉商户分账业务开通申请
	LKL_SPLIT_LEDGER_URL = "/v2/mms/openApi/ledger/applyLedgerMer"
	// LKL_SPLIT_LEDGER_QUERY_URL 拉卡拉商户分账信息查询
	LKL_SPLIT_LEDGER_QUERY_URL = "/v2/mms/openApi/ledger/queryLedgerMer"
	// LKL_SPLIT_LEDGER_RECEIVE_URL 拉卡拉分账接收方创建申请
	LKL_SPLIT_LEDGER_RECEIVE_URL = "/v2/mms/openApi/ledger/applyLedgerReceiver"
	// LKL_SPLIT_LEDGER_RECEIVE_BIND_URL 拉卡拉分账关系绑定申请
	LKL_SPLIT_LEDGER_RECEIVE_BIND_URL = "/v2/mms/openApi/ledger/applyBind"
	// LKL_SPLIT_LEDGER_BALANCE_URL 拉卡拉可分账金额查询
	LKL_SPLIT_LEDGER_BALANCE_URL = "/v3/sacs/queryAmt"
	// LKL_ORDER_SPLIT_LEDGER_URL 拉卡拉订单分账
	LKL_ORDER_SPLIT_LEDGER_URL = "/v3/sacs/separate"
	// LKL_ACCOUNT_BALANCE_QUERY_URL 拉卡拉账户余额查询
	LKL_ACCOUNT_BALANCE_QUERY_URL = "/v2/laep/industry/ewalletBalanceQuery"
	// LKL_ACCOUNT_WITHDRAW_URL 拉卡拉账户提现
	LKL_ACCOUNT_WITHDRAW_URL          = "/v2/laep/industry/ewalletWithdrawD1"
	LKL_SPLIT_LEDGER_RESULT_QUERY_URL = "/v3/sacs/query_plus"
	LKL_SPLIT_LEDGER_FALLBACK_URL     = "/v3/sacs/fallback"
)

const (
	// LKL_TRADE_QUERY_URL 拉卡拉交易查询
	LKL_TRADE_QUERY_URL = "/v3/labs/query/tradequery"
	// LKL_PRE_ORDER_URL 拉卡拉聚合预下单
	LKL_PRE_ORDER_URL = "/v3/labs/trans/preorder"
	// LKL_MERGE_ORDER_URL 拉卡拉主扫合单交易
	LKL_MERGE_ORDER_URL = "/v3/labs/trans/merge/preorder"
	// LKL_REFOUND_URL 拉卡拉退款
	LKL_REFOUND_URL = "/v3/rfd/refund_front/refund"
)

const (
	// LKL_ADD_MER 拉卡拉商户进件
	LKL_ADD_MER = "/v2/mms/openApi/addMer"
	// LKL_QUERY_MER 拉卡拉商户查询
	LKL_QUERY_MER = "/v2/mms/openApi/queryContract"
	// LKL_MER_VALIDATE 拉卡拉商户进件效验
	LKL_MER_VALIDATE = "/v2/mms/openApi/verifyContractInfo"
	// LKL_RECONF_SUBMIT 拉卡拉商户进件复议提交
	LKL_RECONF_SUBMIT = "/v2/mms/openApi/reconsiderSubmit"
	// LKL_EC_APPLY 拉卡拉商户电子合同
	LKL_EC_APPLY = "/v3/mms/open_api/ec/apply"
	// LKL_EC_QUERY_STATUS 拉卡拉商户电子合同查询状态
	LKL_EC_QUERY_STATUS = "/v3/mms/open_api/ec/q_status"
	// LKL_EC_DOWNLOAD 拉卡拉商户电子合同下载
	LKL_EC_DOWNLOAD = "/v3/mms/open_api/ec/download"
	// LKL_EC_APPLY_MANUAL 拉卡拉商户电子合同人工复核申请
	LKL_EC_APPLY_MANUAL = "/v3/mms/open_api/ec/apply_manual"
	// LKL_EC_QMA_STATUS 拉卡拉商户电子合同人工复核结果查询
	LKL_EC_QMA_STATUS = "/v3/mms/open_api/ec/qma_status"
	// LKL_WECHAT_REAL_NAME_QUERY 拉卡拉微信实名认证结果查询
	LKL_WECHAT_REAL_NAME_QUERY = "/v2/mms/openApi/wechatRealNameQuery"
)

// unifiedReturn  统一退货API地址
const (
	// LKL_UNIFIED_RETURN_MERGE_REFUND_URL 拉卡拉合单退货
	LKL_UNIFIED_RETURN_MERGE_REFUND_URL = "/v3/rfd/refund_front/merge_refund"
	// LKL_UNIFIED_RETURN_REFUND_URL 拉卡拉退货
	LKL_UNIFIED_RETURN_REFUND_URL = "/v3/rfd/refund_front/refund"
	// LKL_UNIFIED_RETURN_REFUND_QUERY_URL 拉卡拉退货查询
	LKL_UNIFIED_RETURN_REFUND_QUERY_URL = "/v3/rfd/refund_front/refund_query"
	// LKL_UNIFIED_RETURN_REFUND_FEE_URL 拉卡拉退货手续费查询
	LKL_UNIFIED_RETURN_REFUND_FEE_URL = "/v3/rfd/refund_front/refund_fee"
)
