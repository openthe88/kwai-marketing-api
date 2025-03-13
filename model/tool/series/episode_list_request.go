package series

import "encoding/json"

// EpisodeListRequest 查询短剧的剧集列表
type EpisodeListRequest struct {
	// AdvertiserID 广告主ID（必填）
	AdvertiserID int64 `json:"advertiser_id"`

	// Cursor 游标，第一次传""，后续查询请求根据上一次返回结果的 cursor 传入, 当返回 "no_more" 时，则可结束查询
	Cursor string `json:"cursor"`

	// EpisodeName 剧集名称
	EpisodeName string `json:"episode_name"`

	// PageSize 每次请求的数量，当 cursor 和 page_size 未传入时，则最多返回 5000 条
	PageSize int `json:"page_size"`

	// SeriesID 短剧ID（必填）
	SeriesID int64 `json:"series_id"`

	// UserID 作者ID（必填）
	UserID int64 `json:"user_id"`
}

func (r EpisodeListRequest) Url() string {
	return "gw/dsp/series/episode/list"
}

func (r EpisodeListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
