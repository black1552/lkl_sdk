package common

// Config 拉卡拉SDK配置
type Config struct {
	PublicKey   string `json:"public_key" dc:"公钥字符串"`
	PrivateKey  string `json:"private_key" dc:"私钥字符串"`
	AppId       string `json:"app_id" dc:"lakala应用ID"`
	SerialNo    string `json:"serial_no" dc:"序列号"`
	SubAppId    string `json:"sub_app_id" dc:"子应用ID 微信AppId"`
	Version     string `json:"version" dc:"lakala版本号"`
	AccountType string `json:"account_type" dc:"账户类型"`
	TransType   string `json:"trans_type" dc:"交易类型"`
	NotifyUrl   string `json:"notify_url" dc:"回调地址"`
}
