package dmp

import "encoding/json"

// PopulationSystemRequest 系统推荐定向/排除人群包
type PopulationSystemRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	//Type 系统人群包类型，1:定向，2:排除
	Type int `json:"type,omitempty"`
}

// Url implement GetRequest interface
func (r PopulationSystemRequest) Url() string {
	return "gw/dsp/v1/tool/population/system/recommend"
}

// Encode implement GetRequest interface

func (r PopulationSystemRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
