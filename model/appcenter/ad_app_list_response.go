package appcenter

// AdAppListResponse 获取应用列表 API Response
type AdAppListResponse struct {
	// CurrentPage 当前页
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 分页大小
	PageSize int `json:"page_size,omitempty"`
	// TotalCount 返回条数
	TotalCount int `json:"total_count,omitempty"`
	// List 列表
	List []App `json:"list,omitempty"`
}

// App 应用
type App struct {
	// AccountId 账号ID
	AccountId int64 `json:"account_id,omitempty"`
	// AppDetailImg 应用详情图片
	AppDetailImg string `json:"app_detail_img,omitempty"`
	// AppIconUrl 应用图标链接
	AppIconUrl string `json:"app_icon_url,omitempty"`
	// AppID 应用ID
	AppID int64 `json:"app_id,omitempty"`
	// AppPrivacyUrl 应用隐私政策链接
	AppPrivacyUrl string `json:"app_privacy_url,omitempty"`
	// SourceType 应用来源:1-我创建的 2-共享给我的
	SourceType int `json:"source_type,omitempty"`
	// AppSource 应用创建者信息
	AppSource AppSource `json:"app_source,omitempty"`
	// AppStatus 应用状态:1-审核中 2-审核失败 3-待发布 4-已发布 5-已下架
	AppStatus int `json:"app_status,omitempty"`
	// IosAppId 解析出的iosAppID
	IosAppId string `json:"ios_app_id,omitempty"`
	// OfflineAppStores 下架的应用商店:"huawei","oppo","vivo","xiaomi","meizu","smartisan"
	OfflineAppStores string `json:"offline_app_stores,omitempty"`
	// PackageId 应用包ID
	PackageId int64 `json:"package_id,omitempty"`
	// PackageName 应用包名
	PackageName string `json:"package_name,omitempty"`
	// PackageSize 应用包大小
	PackageSize int64 `json:"package_size,omitempty"`
	// Platform android或ios
	Platform string `json:"platform,omitempty"`
	// PutStatus 投放状态
	PutStatus int `json:"put_status,omitempty"`
	// RealAppName 应用名称
	RealAppName string `json:"real_app_name,omitempty"`
	// RealAppVersion 应用版本信息
	RealAppVersion string `json:"real_app_version,omitempty"`
	// ReviewDetail 审核详情
	ReviewDetail string `json:"review_detail,omitempty"`
	// ReviewStatus 审核状态
	ReviewStatus int `json:"review_status,omitempty"`
	// TraceActivation 转化追踪状态
	TraceActivation int `json:"trace_activation,omitempty"`
	// UpdateTime 更新时间; 单位：毫秒
	UpdateTime int64 `json:"update_time,omitempty"`
	// URL 应用下载地址
	URL string `json:"url,omitempty"`
	// UseSDK 是否接入快手广告监测SDK; 0：未接入，1：已接入
	UseSDK int `json:"use_sdk,omitempty"`
	// VersionCode 应用版本号
	VersionCode int64 `json:"version_code,omitempty"`
	// ShareType 共享类型，0-不共享，1-账号，2-主体
	ShareType int `json:"share_type,omitempty"`
}

// AppSource 应用创建者信息
type AppSource struct {
	// AccountId 应用创建账号id
	AccountId int64 `json:"account_id,omitempty"`
	// AccountName 应用创建账号名称
	AccountName string `json:"account_name,omitempty"`
}
