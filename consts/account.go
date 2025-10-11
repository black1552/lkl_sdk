package consts

type AttType string

const (
	ATT_TYPE_FR_ID_CARD_FRONT                 AttType = "FR_ID_CARD_FRONT"                 // 法人身份证正面
	ATT_TYPE_FR_ID_CARD_BEHIND                AttType = "FR_ID_CARD_BEHIND"                // 法人身份证反面
	ATT_TYPE_ID_CARD_FRONT                    AttType = "ID_CARD_FRONT"                    // 结算人身份证正面
	ATT_TYPE_ID_CARD_BEHIND                   AttType = "ID_CARD_BEHIND"                   // 结算人身份证反面
	ATT_TYPE_BANK_CARD                        AttType = "BANK_CARD"                        // 银行卡
	ATT_TYPE_BUSINESS_LICENCE                 AttType = "BUSINESS_LICENCE"                 // 营业执照
	ATT_TYPE_MERCHANT_PHOTO                   AttType = "MERCHANT_PHOTO"                   // 商户门头照片
	ATT_TYPE_SHOPINNER                        AttType = "SHOPINNER"                        // 商铺内部照片
	ATT_TYPE_XY                               AttType = "XY"                               // 线下纸质协议
	ATT_TYPE_NETWORK_XY                       AttType = "NETWORK_XY"                       // 电子协议
	ATT_TYPE_HT                               AttType = "HT"                               // 租赁合同
	ATT_TYPE_COOPERATION_QUALIFICATION_PROOF  AttType = "COOPERATION_QUALIFICATION_PROOF"  // 合作资质证明
	ATT_TYPE_FOOD_QUALIFICATION_PROOF         AttType = "FOOD_QUALIFICATION_PROOF"         // 食品经营相关资质
	ATT_TYPE_NO_LEGAL_PERSON_SETT_AUTH_LETTER AttType = "NO_LEGAL_PERSON_SETT_AUTH_LETTER" // 非法人结算授权书
	ATT_TYPE_SPLIT_ENTRUST_FILE               AttType = "SPLIT_ENTRUST_FILE"               // 结算授权委托书
	ATT_TYPE_RENTAL_AGREEMENT                 AttType = "RENTAL_AGREEMENT"                 // 集市方与场地方间的租赁协议
	ATT_TYPE_SPLIT_COOPERATION_FILE           AttType = "SPLIT_COOPERATION_FILE"           // 集市方与摊主间的合作协议
	ATT_TYPE_OTHERS                           AttType = "OTHERS"                           //  其他
)

func (at AttType) Ptr() string {
	return string(at)
}

type MgtFlag string

const (
	MGT_FLAG_AVERAGE   MgtFlag = "01" // 一般账户
	MGT_FLAG_HOUSEHOLD MgtFlag = "03" // 虚户
)

type SplitRange string

const (
	SPLIT_RANGE_ALL  SplitRange = "ALL"  // 全部交易分账
	SPLIT_RANGE_MARK SplitRange = "MARK" // 标记交易分账
)

func (sr SplitRange) Ptr() string {
	return string(sr)
}

type SepFundSource string

const (
	SEP_FUND_SOURCE_TR SepFundSource = "TR" // 交易分账
	SEP_FUND_SOURCE_BA SepFundSource = "BR" // 余额分账
)

func (sf SepFundSource) Ptr() string {
	return string(sf)
}

type SplitLaunchMode string

const (
	SPLIT_LAUNCH_MODE_AUTO      SplitLaunchMode = "AUTO"      // 自动触发分账
	SPLIT_LAUNCH_MODE_POINTRULE SplitLaunchMode = "POINTRULE" // 指定规则分账
	SPLIT_LAUNCH_MODE_MANUAL    SplitLaunchMode = "MANUAL"    // 手动分账
)

func (slm SplitLaunchMode) Ptr() string {
	return string(slm)
}

type SplitSettleType string

const (
	SPLIT_SETTLE_TYPE_01 SplitSettleType = "01" // 主扫现结
	SPLIT_SETTLE_TYPE_03 SplitSettleType = "03" // 交易自动结算
)

func (sst SplitSettleType) Ptr() string {
	return string(sst)
}

type SplitRuleSource string

const (
	SPLIT_RULE_SOURCE_MER      SplitRuleSource = "MER"      // 商户分账规则
	SPLIT_RULE_SOURCE_PLATFORM SplitRuleSource = "PLATFORM" // 平台分账规则
)

func (srs SplitRuleSource) Ptr() string {
	return string(srs)
}

type AcctTypeCode string

const (
	ACCT_TYPE_CODE_57 AcctTypeCode = "57" // 对公
	ACCT_TYPE_CODE_58 AcctTypeCode = "58" // 对私
)

func (act AcctTypeCode) Ptr() string {
	return string(act)
}

type AcctCertificateType string

const (
	ACCT_CERTIFICATE_TYPE_ID_CARD        AcctCertificateType = "17" // 身份证
	ACCT_CERTIFICATE_TYPE_PASSPORT       AcctCertificateType = "18" // 护照
	ACCT_CERTIFICATE_TYPE_HONGKONG_MACAO AcctCertificateType = "19" // 港澳居民来往内地通行证
	ACCT_CERTIFICATE_TYPE_TRAVEL         AcctCertificateType = "20" // 台湾居民来往内地通行证
)

func (act AcctCertificateType) Ptr() string {
	return string(act)
}

type CalType string

const (
	CAL_TYPE_0 CalType = "0" // 按固定金额分账
	CAL_TYPE_1 CalType = "1" // 按比例分账
)

func (ct CalType) Ptr() string {
	return string(ct)
}

type PayType string

const (
	PAY_TYPE_01 PayType = "01" // 收款账户
	PAY_TYPE_02 PayType = "02" // 付款账户
	PAY_TYPE_03 PayType = "03" // 分账商户账户
	PAY_TYPE_04 PayType = "04" // 分账接收方账户
	PAY_TYPE_05 PayType = "05" // 充值代付账户
	PAY_TYPE_06 PayType = "06" // 结算代付账户
)

func (pt PayType) Ptr() string {
	return string(pt)
}
