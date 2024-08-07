package tool

import "encoding/json"

// TargetBehaviorRequest 行为意向4.0，类目查询接口
type TargetBehaviorRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

func (r TargetBehaviorRequest) Url() string {
	return "/gw/dsp/target/option/behavior_interest"
}

func (r TargetBehaviorRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
