package app

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// ShareCorpListRequest 获取应用已共享主体列表 API Request
type ShareCorpListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppID 应用id
	AppID uint64 `json:"app_id,omitempty"`
}

func (r ShareCorpListRequest) Url() string {
	return "gw/dsp/appcenter/app/shareCorpList"
}

// Encode implement PostRequest interface
func (r ShareCorpListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ShareCorpListResponse 获取应用已共享主体列表 API Response
type ShareCorpListResponse struct {
	// List 主体列表
	List []ShareCorp `json:"list,omitempty"`
}
