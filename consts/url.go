package consts

const (
	BASE_URL = "https://s2.lakala.com/api"
)

// basePay/labs 基础支付聚合扫码API地址
const (
	// LKL_BASE_URL_TRANS_PREORDER_URL 拉卡拉聚合主扫
	LKL_BASE_URL_TRANS_PREORDER_URL = "/v3/labs/trans/preorder"
	// LKL_DCP_TRANS_PREORDER_URL 拉卡拉直连主扫
	LKL_DCP_TRANS_PREORDER_URL = "/v3/dcp/trans/preorder"
	// LKL_BASE_URL_TRANS_PREORDER_ENCRY_URL 拉卡拉聚合主扫交易（全报文加密）
	LKL_BASE_URL_TRANS_PREORDER_ENCRY_URL = "/v3/labs/trans/preorder_encry"
	// LKL_BASE_URL_TRANS_MICROPAY_URL 拉卡拉聚合被扫
	LKL_BASE_URL_TRANS_MICROPAY_URL = "/v3/labs/trans/micropay"
	// LKL_BASE_URL_TRANS_MICROPAY_ENCRY_URL 拉卡拉聚合被扫接口（全报文加密）
	LKL_BASE_URL_TRANS_MICROPAY_ENCRY_URL = "/v3/labs/trans/micropay_encry"
	// LKL_BASE_URL_TRANS_SHARE_CODE_URL 拉卡拉聚合扫码申请支付宝吱口令
	LKL_BASE_URL_TRANS_SHARE_CODE_URL = "/v3/labs/trans/share_code"
	// LKL_BASE_URL_TRANS_QUERY_URL 拉卡拉聚合扫码交易查询
	LKL_BASE_URL_TRANS_QUERY_URL = "/v3/labs/query/tradequery"
	// LKL_BASE_URL_RELATION_CLOSE_URL 拉卡拉聚合扫码关单
	LKL_BASE_URL_RELATION_CLOSE_URL = "/v3/labs/relation/close"
	// LKL_BASE_URL_RELATION_REVOKED_URL 拉卡拉聚合扫码撤销
	LKL_BASE_URL_RELATION_REVOKED_URL = "/v3/labs/relation/revoked"
	// LKL_BASE_URL_RELATION_REFUND_URL 拉卡拉聚合扫码退款
	LKL_BASE_URL_RELATION_REFUND_URL = "/v3/labs/relation/refund"
)

// basePay/dcp 基础支付直接扫码API地址
const ()

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
