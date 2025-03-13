package series

// EpisodeListResponse 查询短剧的剧集列表
type EpisodeListResponse struct {
	// Cursor 游标
	Cursor string `json:"cursor"`

	// Episodes MAPI 短剧剧集信息
	Episodes []MapiEpisodeInfoSnake `json:"episodes"`
}

// MapiEpisodeInfoSnake MAPI 剧集信息
type MapiEpisodeInfoSnake struct {
	// ID 剧集ID
	ID int64 `json:"id"`

	// Name 剧集名称
	Name string `json:"name"`

	// Description 剧集描述
	Description string `json:"description"`

	// OrderNo 剧集顺序
	OrderNo int `json:"order_no"`

	// SerialID 短剧ID（旧），合集ID，与短剧ID不同
	SerialID int64 `json:"serial_id"`
}
