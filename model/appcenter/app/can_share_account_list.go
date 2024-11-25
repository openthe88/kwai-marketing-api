package app

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// CanShareAccountListRequest 获取可共享的账号列表 API Request
type CanShareAccountListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SearchAccountID 精确查找的账号
	SearchAccountID []uint64 `json:"search_account_id,omitempty"`
	// CurrentPage 当前页
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 分页大小
	PageSize int `json:"page_size,omitempty"`
}

func (r CanShareAccountListRequest) Url() string {
	return "gw/dsp/appcenter/app/canShare/accountList"
}

// Encode implement PostRequest interface
func (r CanShareAccountListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// CanShareAccountListResponse 获取可共享的账号列表 API Response
type CanShareAccountListResponse struct {
	// CurrentPage 当前页
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 分页大小
	PageSize int `json:"page_size,omitempty"`
	// TotalCount 总数
	TotalCount int64 `json:"total_count,omitempty"`
	// List 账号列表
	List []ShareAccount `json:"list,omitempty"`
}
