package consts

type AccountType string

const (
	ACCOUNT_TYPE_WECHAT     AccountType = "WECHAT"     // 微信
	ACCOUNT_TYPE_ALIPAY     AccountType = "ALIPAY"     // 支付宝
	ACCOUNT_TYPE_UQRCODEPAY AccountType = "UQRCODEPAY" // 银联
	ACCOUNT_TYPE_BASTPAY    AccountType = "BASTPAY"    // 翼支付
	ACCOUNT_TYPE_SUNING     AccountType = "SUNING"     // 苏宁
	ACCOUNT_TYPE_LKLACC     AccountType = "LKLACC"     // 拉卡拉支付账户
	ACCOUNT_TYPE_NUCSPAY    AccountType = "NUCSPAY"    // 网联小钱包
	ACCOUNT_TYPE_JD         AccountType = "JD"
)

func (at AccountType) Ptr() string {
	return string(at)
}

type TransType string

const (
	TRANS_TYPE_NATIVE    TransType = "41" // NATIVE（ALIPAY, 云闪付支持, 京东白条分期）
	TRANS_TYPE_JSAPI     TransType = "51" // JSAPI（微信公众号支付, 支付宝服务窗、JS支付, 翼支付JS支付, 拉卡拉钱包支付, 京东白条分期）
	TRANS_TYPE_WECHATPAY TransType = "71" // 微信小程序支付
	TRANS_TYPE_APP       TransType = "61" // APP支付（微信APP支付），String(2)
)

func (tt TransType) Ptr() string {
	return string(tt)
}

type SettleType string

const (
	SETTLE_TYPE_NORMAL  SettleType = "0" // 常规结算
	SETTLE_TYPE_SPECIAL SettleType = "1" // 拉卡拉分账通结算
)

func (st SettleType) Ptr() string {
	return string(st)
}

type RefundAccMode string

// 00退货账户余额 05商户余额 06终端余额 30终点账户
const (
	REFUND_ACC_MODE_USER_BALANCE     RefundAccMode = "00" // 退货账户余额
	REFUND_ACC_MODE_MERCHANT_BALANCE RefundAccMode = "05" // 商户余额
	REFUND_ACC_MODE_TERMINAL_BALANCE RefundAccMode = "06" // 终端余额
	REFUND_ACC_MODE_ACCOUNT          RefundAccMode = "30" // 终点账户
)

func (rat RefundAccMode) Ptr() string {
	return string(rat)
}

type RefundAmtSts string

const (
	REFUND_AMT_STS_NORMAL RefundAmtSts = "00" // 分账前
	REFUND_AMT_STS_SPLIT  RefundAmtSts = "01" // 分账后
)

func (ras RefundAmtSts) Ptr() string {
	return string(ras)
}
