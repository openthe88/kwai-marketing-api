package asset_media

// MediaReq media请求结构体
type MediaReq struct {
	// MediaID 媒体包ID
	MediaID uint64 `json:"media_id"`

	// Name 媒体包名字，修改媒体包名
	Name string `json:"name"`

	// SourceType 媒体包来源，0-不限，未指定，1-行业优质流量包，2-广告主自定义，默认 0
	SourceType int `json:"source_type"`

	// PosIDs 广告位列表，修改媒体包广告位
	PosIDs []string `json:"pos_ids"`

	// Status 状态，删除媒体包，0-下线，1-上线，默认 1
	Status int `json:"status"`
}

// RelateUnitsItem 关联广告组信息列表
type RelateUnitsItem struct {
	// UnitID 广告组ID
	UnitID uint64 `json:"unit_id"`

	// UnitName 广告组名字
	UnitName string `json:"unit_name"`

	// CampaignID 广告计划ID
	CampaignID uint64 `json:"campaign_id"`

	// CampaignType 广告计划类型，计划类型，过滤筛选条件：1 - 作品推广；2 - 提升应用安装；3 - 获取电商下单；4 - 推广品牌活动；5 - 收集销售线索；6 - 保量广告；7 - 提高应用活跃
	CampaignType int `json:"campaign_type"`

	// PutStatus 在投状态，UNKNOWN_PUT_STATUS = 0; PUT_STATUS_OPEN = 1; PUT_STATUS_PAUSE = 2; PUT_STATUS_DELETED = 3
	PutStatus int `json:"put_status"`

	// ReviewStatus 审核状态，UNKNOWN_REVIEW_STATUS = 0; REVIEW_WAIT_AUDIT = 1; REVIEW_THROUGH = 2; REVIEW_NOT_THROUGH = 3; BANNED = 4
	ReviewStatus int `json:"review_status"`

	// MediaUsing 媒体包是否在使用
	MediaUsing bool `json:"media_using"`

	// MediaType 媒体包的使用类型，0-未知，1-定向，2-排除
	MediaType int `json:"media_type"`
}

type MediaRes struct {
	// MediaID 媒体包ID
	MediaID uint64 `json:"mediaId"`

	// Name 媒体包名字
	Name string `json:"name"`

	// AccountId 广告主ID
	AccountId uint64 `json:"accountId"`

	// SourceType 媒体包来源，0-不限，未指定，1-行业优质流量包，2-广告主自定义，默认 0
	SourceType int `json:"sourceType"`

	// PosIDs 广告位列表，修改媒体包广告位
	PosIDs []string `json:"posIds"`

	// Status 状态，删除媒体包，0-下线，1-上线，默认 1
	Status int `json:"status"`

	// CreateTime 创建时间，时间戳
	CreateTime int64 `json:"createTime"`

	// UpdateTime 更新时间，时间戳
	UpdateTime int64 `json:"updateTime"`
}
