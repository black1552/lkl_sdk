package model

type BaseModel[T any] struct {
	ReqTime string `json:"req_time"`
	Version string `json:"version"`
	ReqData T      `json:"req_data"`
}
