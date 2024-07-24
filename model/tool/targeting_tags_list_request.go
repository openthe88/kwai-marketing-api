package tool

import (
	"encoding/json"
)

// TargetingTagsListRequest 获取可选的定向标签
type TargetingTagsListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Type 类型; BUSINESS_INTEREST：商业兴趣APP_INTEREST：APP行为 FANS_STAR：网红类别INTEREST_VIDEO：兴趣视频 APP_INTEREST_ID:APP行为 (新）
	Type string `json:"type,omitempty"`
}

// Url implement GetRequest interface
func (r TargetingTagsListRequest) Url() string {
	return "v1/tool/targeting_tags/list"
}

// Encode implement GetRequest interface
func (r TargetingTagsListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
