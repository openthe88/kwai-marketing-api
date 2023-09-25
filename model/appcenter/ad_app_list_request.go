package appcenter

import (
	"encoding/json"
)

// AdAppListRequest 获取应用列表 API Request
type AdAppListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// ListType 列表类型
	ListType uint `json:"list_type"`
	// Platform android或ios
	Platform string `json:"platform,omitempty"`
	// KeyWord 关键词
	KeyWord string `json:"key_word,omitempty"`
	// AppStatus 应用状态
	AppStatus uint `json:"app_status,omitempty"`
	// StartDate 发布时间范围-起始
	StartDate string `json:"start_date,omitempty"`
	// EndDate 发布时间范围-截止
	EndDate string `json:"end_date,omitempty"`
	// Page 页码，默认 1
	Page int `json:"page,omitempty"`
	// PageSize 页数，默认 20
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r AdAppListRequest) Url() string {
	return "gw/dsp/appcenter/app/list"
}

// Encode implement PostRequest interface
func (r AdAppListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
