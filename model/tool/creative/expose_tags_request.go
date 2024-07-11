package creative

import "encoding/json"

// ExposeTagsRequest 获取创意推荐理由 API Request
type ExposeTagsRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	//CampaignType 计划营销目标类型
	CampaignType int `json:"campaign_type,omitempty"`
}

func (r ExposeTagsRequest) Url() string {
	return "v1/tool/expose_tags/list"
}

func (r ExposeTagsRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
