package series

import "encoding/json"

// UserListRequest 获取授权的短剧作者列表
type UserListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
}

func (r UserListRequest) Url() string {
	return "gw/dsp/series/auth/user/list"
}

func (r UserListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
