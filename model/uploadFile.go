package model

import "github.com/black1552/lkl_sdk/consts"

type UploadFileRequest struct {
	ReqData   *UploadFileReqData `json:"reqData"`
	Ver       string             `json:"ver"`
	Timestamp string             `json:"timestamp"`
	ReqId     string             `json:"reqId"`
}

type UploadFileReqData struct {
	Version    string         `json:"version"`
	OrderNo    string         `json:"orderNo"`
	AttType    consts.AttType `json:"attType"`
	AttExtName string         `json:"attExtName"`
	AttContext string         `json:"attContext"`
	OrgCode    string         `json:"orgCode"`
}

type UploadFileResponse struct {
	CmdRetCode string `json:"cmdRetCode"`
	ReqId      string `json:"reqId"`
	RespData   struct {
		AttType   string `json:"attType"`
		OrderNo   string `json:"orderNo"`
		OrgCode   string `json:"orgCode"`
		AttFileId string `json:"attFileId"`
	} `json:"respData"`
	RetCode   string `json:"retCode"`
	RetMsg    string `json:"retMsg"`
	Timestamp int64  `json:"timestamp"`
	Ver       string `json:"ver"`
}

func (u *UploadFileResponse) SuccessOrFail() bool {
	return u.RetCode == "000000"
}
