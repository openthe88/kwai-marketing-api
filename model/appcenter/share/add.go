package share

import "github.com/bububa/kwai-marketing-api/model"

// AddRequest 添加应用共享 API Request
type AddRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppID 应用ID
	AppID uint64 `json:"app_id,omitempty"`
	// ShareAdvertiserIDs 要共享的账号ID列表
	ShareAdvertiserIDs []uint64 `json:"share_advertiser_ids,omitempty"`
	// ShareCorpIDs 要共享的主体ID列表
	ShareCorpIDs []uint64 `json:"share_corp_ids,omitempty"`
	// ShareType 共享类型，0-不共享，1-账号，2-主体
	ShareType int `json:"share_type,omitempty"`
}

func (r AddRequest) Url() string {
	return "gw/dsp/appcenter/share/add"
}

// Encode implement PostRequest interface
func (r AddRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// AddResponse 添加应用共享响应 API Response
type AddResponse struct {
	// Result 添加应用共享是否成功
	// true-成功，false-失败
	Result bool `json:"result"`
}
