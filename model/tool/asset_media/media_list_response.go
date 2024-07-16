package asset_media

// MediaListResponse 获取联盟定向媒体包列表
type MediaListResponse struct {
	// TotalCount 结果总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 关联广告组信息列表
	Details []Details `json:"details,omitempty"`
}

type Details struct {
	RelateUnits []RelateUnitsItem `json:"relate_units,omitempty"`
	Media       MediaRes          `json:"media,omitempty"`
}
