package asset

// AdvCardListResponse 获取高级创意列表
type AdvCardListResponse struct {
	// TotalCount 卡片总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 卡片列表
	Details []AdvCardList `json:"details,omitempty"`
}

type AdvCardList struct {
	// AdvCardID 卡片ID
	AdvCardID uint64 `json:"adv_card_id"`

	// TemplateName 模版名
	TemplateName string `json:"template_name"`

	// UnitCount 关联广告组数
	UnitCount int `json:"unit_count"`

	// URL 图片URL
	URL string `json:"url"`

	// Title 标题
	Title string `json:"title"`

	// SubTitle 副标题
	SubTitle string `json:"sub_title"`

	// Price 原价格(单位：分)
	Price int `json:"price"`

	// SalePrice 售卖价(单位：分)
	SalePrice int `json:"sale_price"`

	// BeginTime 倒计时卡开始时间
	BeginTime int `json:"begin_time"`

	// EndTime 倒计时卡结束时间
	EndTime int `json:"end_time"`

	// SDPACardContent SDPA商品卡样式内容，card_type=131/132/133/134专用
	SDPACardContent SDPAContent `json:"sdpa_card_content"`
}

// SDPAContent SDPA商品卡样式内容的结构体定义
type SDPAContent struct {
	// Title 商品卡标题，card_type=131/132/133/134专用
	Title string `json:"title"`

	// SubTitle 商品卡副标题(描述/简介)，card_type=131/132/133/134专用
	SubTitle string `json:"sub_title"`

	// Icon 商品卡图片，card_type=131/132/133/134专用
	Icon string `json:"icon"`

	// Description 商品卡价格类型描述，card_type=131/132/133/134专用
	Description string `json:"description"`

	// Symbol 商品卡价格符号描述，card_type=131/132/133/134专用
	Symbol string `json:"symbol"`

	// Price 商品卡价格(现价)，card_type=131/132/133/134专用
	Price float64 `json:"price"`

	// TailDescription 商品卡价格单位描述，card_type=131/132/133/134专用
	TailDescription string `json:"tail_description"`

	// SalePrice 商品卡价格(原价)，card_type=131/132/133/134专用
	SalePrice float64 `json:"sale_price"`

	// TagText 商品卡角标，card_type=131/132/133/134专用
	TagText string `json:"tag_text"`

	// HotText 商品卡热度标签，card_type=131/132/133/134专用
	HotText string `json:"hot_text"`

	// City 商品卡城市标签，card_type=131/132/133/134专用
	City string `json:"city"`
}
