package app

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// ShareListCorpAccountRequest 获取单个主体下共享账号列表 API Request
type ShareListCorpAccountRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppID 应用id
	AppID uint64 `json:"app_id,omitempty"`
	// CorpID 主体ID
	CorpID uint64 `json:"corp_id,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

func (r ShareListCorpAccountRequest) Url() string {
	return "gw/dsp/appcenter/app/share/listCorpAccount"
}

// Encode implement PostRequest interface
func (r ShareListCorpAccountRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ShareListCorpAccountResponse 获取单个主体下共享账号列表 API Response
type ShareListCorpAccountResponse struct {
	// List 主体列表
	List []ShareCorp `json:"list,omitempty"`
}
