package appcenter

// AdAppReleaseListResponse 获取应用列表 API Response
type AdAppReleaseListResponse struct {
	// CurrentPage 当前页
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 分页大小
	PageSize int `json:"page_size,omitempty"`
	// TotalCount 返回条数
	TotalCount int `json:"total_count,omitempty"`
	// List 列表
	List []App `json:"list,omitempty"`
}
