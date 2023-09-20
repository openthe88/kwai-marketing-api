package appcenter

import (
	"encoding/json"
)

// AdSubPackageListRequest 获取分包列表 API Request
type AdSubPackageListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// ListType 列表类型:不填为分包管理列表，填2-分包回收列表
	ListType uint `json:"list_type"`
	// KeyWord 关键词
	KeyWord string `json:"key_word,omitempty"`
	// AppId 应用id
	AppId uint64 `json:"app_id,omitempty"`
	// Page 页码，默认 1
	Page int `json:"page,omitempty"`
	// PageSize 页数，默认 20
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r AdSubPackageListRequest) Url() string {
	return "gw/dsp/appcenter/subPackage/list"
}

// Encode implement PostRequest interface
func (r AdSubPackageListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
