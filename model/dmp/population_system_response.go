package dmp

// PopulationSystemResponse 系统推荐定向/排除人群包
type PopulationSystemResponse struct {
	//interest_target 店铺外兴趣人群
	InterestTarget []PopulationSystem `json:"interest_target"`
	//industry_target 行业人群
	IndustryTarget []PopulationSystem `json:"industry_target"`
	//negative_exclude LA负反馈人群
	NegativeExclude []PopulationSystem `json:"negative_exclude"`
	//react_target 店铺内互动人群
	ReactTarget []PopulationSystem `json:"react_target"`
	//consume_target 店铺消费者人群
	ConsumeTarget []PopulationSystem `json:"consume_target"`
	//common_exclude 通用排除人群
	CommonExclude []PopulationSystem `json:"common_exclude"`
}

// PopulationSystem 系统人群包字段结构体
type PopulationSystem struct {
	// OrientationID 定向ID
	OrientationID uint64 `json:"orientation_id"`

	// OrientationName 定向名称
	OrientationName string `json:"orientation_name"`

	// AccountID 账户ID
	AccountID uint64 `json:"account_id"`

	// Type 类型
	Type int `json:"type"`

	// PopulationType 人群类型
	PopulationType int `json:"population_type"`

	// PopulationTypeName 人群类型名称
	PopulationTypeName string `json:"population_type_name"`

	// RecordSize 记录大小
	RecordSize uint64 `json:"record_size"`

	// MatchSize 匹配大小
	MatchSize uint64 `json:"match_size"`

	// CoverNum 覆盖数量
	CoverNum uint64 `json:"cover_num"`

	// Status 状态
	Status int `json:"status"`

	// SrcID 来源ID
	SrcID uint64 `json:"src_id"`

	// CreateTime 创建时间
	CreateTime int64 `json:"create_time"`

	// VerifyTime 审核时间
	VerifyTime int64 `json:"verify_time"`

	// ProfileTags 配置标签
	ProfileTags string `json:"profile_tags"`

	// Unbind 解绑
	Unbind string `json:"unbind"`

	// UnbindType 解绑类型
	UnbindType string `json:"unbind_type"`

	// SuccessUnbind 成功解绑
	SuccessUnbind string `json:"success_unbind"`

	// FailUnbind 失败解绑
	FailUnbind string `json:"fail_unbind"`

	// TpCode TP代码
	TpCode int `json:"tp_code"`

	// IsExcludePopulation 是否排除人群
	IsExcludePopulation bool `json:"is_exclude_population"`

	// UpdateTime 更新时间
	UpdateTime int64 `json:"update_time"`

	// CategoryType 类别类型
	CategoryType int `json:"category_type"`

	// CanExclude 是否可以排除
	CanExclude int `json:"can_exclude"`

	// CanTarget 是否可以定向
	CanTarget int `json:"can_target"`

	// SrcType 人群包来源，0：DMP平台（磁力万象）；1：MAPI平台；2：CDP平台（磁力方舟）；3：CDP投放共建类型
	SrcType int `json:"src_type"`
}
