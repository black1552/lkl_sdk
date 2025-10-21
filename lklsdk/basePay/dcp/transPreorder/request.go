package transpreorder

type TransPreorderRequest struct {
	ReqTime string                    `json:"req_time"`
	Version string                    `json:"version"`
	ReqData *TransPreorderRequestData `json:"req_data"`
}

type TransPreorderRequestData struct {
	OutTradeNo  string `json:"out_trade_no"`
	TotalAmount string `json:"total_amount"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	NotifyUrl   string `json:"notify_url"`
	ExpireTime  string `json:"expire_time"`
}
