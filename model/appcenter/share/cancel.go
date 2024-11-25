package share

import "github.com/bububa/kwai-marketing-api/model"

// CancelRequest 取消应用共享 API Request
type CancelRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppID 应用ID
	AppID uint64 `json:"app_id,omitempty"`
	// CancelShareAdvertiserIDs 要取消的账号ID列表
	CancelShareAdvertiserIDs []uint64 `json:"cancel_share_advertiser_ids,omitempty"`
	// CancelShareCorpIDs 要取消的主体ID集合
	CancelShareCorpIDs []uint64 `json:"cancel_share_corp_ids,omitempty"`
	// ShareType 共享类型，0-不共享，1-账号，2-主体
	ShareType int `json:"share_type,omitempty"`
}

func (r CancelRequest) Url() string {
	return "gw/dsp/appcenter/share/cancel"
}

// Encode implement PostRequest interface
func (r CancelRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// CancelResponse 取消应用共享响应 API Response
type CancelResponse struct {
	// Result 取消应用共享是否成功
	// true-成功，false-失败
	Result bool `json:"result"`
}
