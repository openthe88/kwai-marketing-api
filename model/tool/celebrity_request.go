package tool

import (
	"encoding/json"
)

// CelebrityRequest 搜索网红
type CelebrityRequest struct {
	// AdvertiserID 账号ID，必填
	AdvertiserID uint64 `json:"advertiser_id"`

	// Keyword 查询关键词，可选
	Keyword string `json:"keyword"`

	// SecondLabel secondLabel，可选
	SecondLabel string `json:"second_label"`

	// AuthorID authorId，可选
	AuthorID uint64 `json:"author_id"`

	// Test 测试，可选
	Test uint64 `json:"test"`
}

func (r CelebrityRequest) Url() string {
	return "gw/dsp/v1/celebrity/list"
}

func (r CelebrityRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
