package series

import "encoding/json"

// ListRequest 查询授权短剧列表
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	//SeriesTitle 短剧标题
	SeriesTitle string `json:"series_title,omitempty"`
	//UserId 作者id
	UserId int64 `json:"user_id,omitempty"`
}

func (r ListRequest) Url() string {
	return "gw/dsp/series/list"
}

func (r ListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
