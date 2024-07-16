package tool

import (
	"encoding/json"
)

// IndustryCelebrityRequest 快手网红-网红分类
type IndustryCelebrityRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

func (r IndustryCelebrityRequest) Url() string {
	return "gw/dsp/v1/industry/celebrity"
}

func (r IndustryCelebrityRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
