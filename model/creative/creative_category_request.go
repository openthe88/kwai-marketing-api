package creative

import "encoding/json"

// CategoryRequest 获取创意分类查询接口 API Request
type CategoryRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

func (r CategoryRequest) Url() string {
	return "v1/creative/creative_category/list"
}

func (r CategoryRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
