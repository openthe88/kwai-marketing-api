package appcenter

import (
	"encoding/json"
)

// AdAppReleaseListRequest 获取应用列表 API Request
type AdAppReleaseListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// ListType 列表类型:不传-全部, 1-我创建的, 2-共享给我的
	ListType uint `json:"list_type,omitempty"`
	// Platform android或ios
	Platform string `json:"platform,omitempty"`
	// AppIds 批量应用id查询:最多支持查询100个
	AppIds []uint64 `json:"app_ids,omitempty"`
	// KeyWord 关键词:支持应用ID或应用名称搜索
	KeyWord string `json:"key_word,omitempty"`
	// Page 页码，默认 1
	Page int `json:"page,omitempty"`
	// PageSize 页数，默认 10
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r AdAppReleaseListRequest) Url() string {
	return "gw/dsp/appcenter/app/release/list"
}

// Encode implement PostRequest interface
func (r AdAppReleaseListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
