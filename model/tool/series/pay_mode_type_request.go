package series

import "encoding/json"

// PayModeTypeRequest 短剧付费模式查询接口
type PayModeTypeRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	//SeriesId 短剧id
	SeriesId uint64 `json:"series_id,omitempty"`
	//UserId 作者id
	UserId uint64 `json:"user_id,omitempty"`
}

func (r PayModeTypeRequest) Url() string {
	return "gw/dsp/series/payModeType"
}

func (r PayModeTypeRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
