package consts

type ECType string

const (
	ECType001 ECType = "EC001" // EC001 : 特约商户支付服务合作协议V3.1(商户入网)
	ECType002 ECType = "EC002" // EC002 : 特约商户支付服务合作协议V3.2（商户入网+分账）
	ECType003 ECType = "EC003" // EC003 : 分账结算授权委托书
	ECType004 ECType = "EC004" // EC004 : 特约商户支付服务合作协议V3.3（商户入网）
	ECType005 ECType = "EC005" // EC005 : 特约商户支付服务合作协议V3.3（商户入网+分账）
	ECType007 ECType = "EC007" // EC007 : 特约商户支付服务合作协议V4.1 + 结算授权委托书 (商户入网 + 分账 )
	ECType008 ECType = "EC008" // EC008 : 特约商户支付服务合作协议V4.1 (商户入网)
	ECType009 ECType = "EC009" // EC009 : 结算授权委托书
)

func (e ECType) Ptr() string {
	return string(e)
}

type CertType string

const (
	CertType_RESIDENT_ID   CertType = "RESIDENT_ID"   // RESIDENT_ID（身份证)
	CertType_PASSPORT      CertType = "PASSPORT"      // PASSPORT（护照）
	CertType_HK_MACAO_PASS CertType = "HK_MACAO_PASS" // HK_MACAO_PASS（港澳居民往来内地通行证）
	CertType_TAIWAN_PASS   CertType = "TAIWAN_PASS"   // TAIWAN_PASS（台湾居民来往大陆通行证）
)

func (c CertType) Ptr() string {
	return string(c)
}

type AgentTag int

const (
	AgentTag0 AgentTag = iota // 0：不启用  缺省
	AgentTag1  // 1: 启用
)

func (a AgentTag) Ptr() int {
	return int(a)
}
