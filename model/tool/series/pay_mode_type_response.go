package series

// PayModeTypeResponse 短剧付费模式查询接口
type PayModeTypeResponse struct {
}

// MapiSeriesPayModeInfoSnake 短剧付费模式信息
type MapiSeriesPayModeInfoSnake struct {
	// PayMode 付费模式，目前仅支持 payMode=1，即“打包”
	PayMode int `json:"pay_mode"`

	// PayModeDesc 付费模式描述
	PayModeDesc string `json:"pay_mode_desc"`
}
