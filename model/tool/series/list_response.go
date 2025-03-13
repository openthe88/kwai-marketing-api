package series

// ListResponse 查询授权短剧列表
type ListResponse struct {
	Series []MapiSeriesInfoSnake `json:"series"`
}

type MapiSeriesInfoSnake struct {
	// CoverImg 短剧封面
	CoverImg string `json:"cover_img"`

	// Description 短剧描述
	Description string `json:"description"`

	// EpisodeAmount 剧集数量
	EpisodeAmount int `json:"episode_amount"`

	// ID 短剧ID
	ID int64 `json:"id"`

	// Title 短剧标题
	Title string `json:"title"`
}
