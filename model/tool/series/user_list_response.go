package series

// UserListResponse 获取授权的短剧作者列表
type UserListResponse struct {
}

// MapiSeriesUserSnake MAPI 短剧授权用户信息
type MapiSeriesUserSnake struct {
	// HeadURL 短剧作者快手号头像
	HeadURL string `json:"head_url"`

	// UserID 短剧作者快手号 ID
	UserID int64 `json:"user_id"`

	// UserName 短剧作者快手号名称
	UserName string `json:"user_name"`

	// UserSex 短剧作者快手号性别，男性 M，女性 F，U 未知
	UserSex string `json:"user_sex"`
}
