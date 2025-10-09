package consts

const (
	BASE_URL = "https://s2.lakala.com/api"
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
	LKL_ACCOUNT_WITHDRAW_URL = "/v2/laep/industry/ewalletWithdrawD1"
	// LKL_TRADE_QUERY_URL 拉卡拉交易查询
	LKL_TRADE_QUERY_URL = "/v3/labs/query/tradequery"
	// LKL_PRE_ORDER_URL 拉卡拉聚合预下单
	LKL_PRE_ORDER_URL = "/v3/labs/trans/preorder"
	// LKL_MERGE_ORDER_URL 拉卡拉主扫合单交易
	LKL_MERGE_ORDER_URL = "/v3/labs/trans/merge/preorder"
	// LKL_REFOUND_URL 拉卡拉退款
	LKL_REFOUND_URL = "/v3/rfd/refund_front/refund"
)
