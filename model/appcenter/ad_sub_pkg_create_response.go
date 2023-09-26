package appcenter

// AdSubPkgCreateResponse 新建应用分包 API Response
type AdSubPkgCreateResponse struct {
	// Data 返回Data数据
	Data []SubPkg `json:"data"`
}

type SubPkg struct {
	// ChannelId 渠道号(分包号)
	ChannelId string `json:"channel_id,omitempty"`
	// PackageId 应用包ID
	PackageId int64 `json:"package_id,omitempty"`
	// BuildStatus 构建状态 0-创建中，1-构建中，2-构建成功，3-构建失败
	BuildStatus int `json:"build_status,omitempty"`
	// ParentPackageId 绑定的母包id
	ParentPackageId int64 `json:"parent_package_id,omitempty"`
	// Description 分包备注
	Description string `json:"description,omitempty"`
}
