package asset_media

import "encoding/json"

// MediaListRequest 获取联盟定向媒体包列表
type MediaListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Media 媒体包参数结构，用于支持搜索
	Media *MediaReq `json:"media,omitempty"`
	// IsRelatedUnit 是否关联广告组信息，默认 false
	IsRelatedUnit bool `json:"is_related_unit,omitempty"`
	// Page 当前页码，默认为 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小，默认 20，最大不超过 50
	PageSize int `json:"page_size,omitempty"`
}

func (r MediaListRequest) Url() string {
	return "v1/asset/media/list"
}

func (r MediaListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
