package app

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// ShareListRequest 获取应用已共享账号列表 API Request
type ShareListRequest struct {
	// AdvertiserID 广告主ID，必填
	AdvertiserID uint64 `json:"advertiser_id"`
	// AppID 应用ID，必填
	AppID uint64 `json:"app_id"`
	// KeyWord 搜索关键词，用于账号搜索
	KeyWord string `json:"key_word,omitempty"`
	// Page 当前页码，页码，默认为1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小，个数，默认为20
	PageSize int `json:"page_size,omitempty"`
}

// Url implements PostRequest interface
func (r ShareListRequest) Url() string {
	return "gw/dsp/appcenter/app/share/list"
}

// Encode implement PostRequest interface
func (r ShareListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ShareListResponse 获取应用已共享账号列表 API Response
type ShareListResponse struct {
	// CurrentPage 当前页
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 分页大小
	PageSize int `json:"page_size,omitempty"`
	// TotalCount 总数
	TotalCount int64 `json:"total_count,omitempty"`
	// List 账号列表
	List []ShareAccount `json:"list,omitempty"`
}
