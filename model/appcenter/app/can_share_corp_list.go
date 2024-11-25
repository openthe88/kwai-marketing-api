package app

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// CanShareCorpListRequest 获取可共享的主体列表 API Request
type CanShareCorpListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

func (r CanShareCorpListRequest) Url() string {
	return "gw/dsp/appcenter/app/canShare/corpList"
}

// Encode implement PostRequest interface
func (r CanShareCorpListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// CanShareCorpListResponse 获取可共享的主体列表 API Response
type CanShareCorpListResponse struct {
	// List 主体列表
	List []ShareCorp `json:"list,omitempty"`
}
