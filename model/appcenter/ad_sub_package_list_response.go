package appcenter

// AdSubPackageListResponse 获取应用列表 API Response
type AdSubPackageListResponse struct {
	// CurrentPage 当前页
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 分页大小
	PageSize int `json:"page_size,omitempty"`
	// TotalCount 返回条数
	TotalCount int `json:"total_count,omitempty"`
	// List 列表
	List []SubPackage `json:"list,omitempty"`
}

// SubPackage 应用分包
type SubPackage struct {
	// PackageId 应用包ID
	PackageId int64 `json:"package_id,omitempty"`
	// ChannelId 渠道号(分包号)
	ChannelId string `json:"channel_id,omitempty"`
	// Description 分包描述
	Description string `json:"description,omitempty"`
	// RealAppVersion 应用版本信息
	RealAppVersion string `json:"real_app_version,omitempty"`
	// SubPackageStatus 应用分包状态:1-审核中，2-审核失败，3-待发布，4-已发布，5-已下架 6-创建中，7-更新中，8-构建失败
	SubPackageStatus int `json:"sub_package_status,omitempty"`
	// CanUpdate 是否可更新:仅分包管理列表时有效，表示应用分包是否可以更新
	CanUpdate bool `json:"can_update,omitempty"`
	// UpdateTime 更新时间:仅分包管理列表时有效，表示应用分包的更新时间
	UpdateTime int64 `json:"update_time,omitempty"`
	// CanRecycle 是否可恢复:仅分包回收站列表时有效，表示应用分包是否可以恢复
	CanRecycle bool `json:"can_recycle,omitempty"`
	// DeleteTime 删除时间:仅分包回收站列表时有效，表示应用分包的删除时间
	DeleteTime int64 `json:"delete_time,omitempty"`
	// ParentPackageId 分包的母包ID
	ParentPackageId int64 `json:"parent_package_id,omitempty"`
	// URL 应用下载地址
	URL string `json:"url,omitempty"`
}
