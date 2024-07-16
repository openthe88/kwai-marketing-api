package tool

import "encoding/json"

// BehaviorInterestRequest 获取行为与兴趣"类目"
type BehaviorInterestRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

func (r BehaviorInterestRequest) Url() string {
	return "/v1/tool/label/behavior_interest"
}

func (r BehaviorInterestRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
