package tool

import (
	"net/url"
	"strconv"
)

// IndustryCelebrityRequest 快手网红-网红分类
type IndustryCelebrityRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

func (r IndustryCelebrityRequest) Url() string {
	return "gw/dsp/v1/industry/celebrity"
}

func (r IndustryCelebrityRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}
