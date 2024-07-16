package tool

// KeywordQueryResponse 获取行为兴趣"关键词"
type KeywordQueryResponse struct {
	//行为兴趣关键词，详情见下方
	Keyword []Keyword `json:"keyword,omitempty"`
}

type Keyword struct {
	// ID 关键词ID
	ID uint64 `json:"id"`

	// Name 关键词名称
	Name string `json:"name"`

	// Count 当前关键词的人群预估数量
	Count uint64 `json:"count"`
}
