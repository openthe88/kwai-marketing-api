package tool

import "encoding/json"

// KeywordQueryRequest 获取行为兴趣"关键词"
type KeywordQueryRequest struct {
	// AdvertiserID 广告主ID，必填
	AdvertiserID uint64 `json:"advertiser_id"`

	// QueryStr 行为兴趣关键词名称，选填
	QueryStr string `json:"query_str,omitempty"`

	// Type 查询类型，必填。0：按照 query_str 模糊查询；1：按照 id 进行查询
	Type int `json:"type"`

	// IDs 选填的 ID 列表
	IDs []uint64 `json:"ids,omitempty"`
}

func (r KeywordQueryRequest) Url() string {
	return "v1/tool/keyword/query"
}

func (r KeywordQueryRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
