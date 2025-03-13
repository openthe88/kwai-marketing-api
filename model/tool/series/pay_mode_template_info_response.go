package series

// PayModeTemplateInfoResponse 查询短剧付费模板信息接口
type PayModeTemplateInfoResponse struct {
}

type MapiSeriesPayModeTemplateInfoSnake struct {
	SeriesPayMode       int    `json:"series_pay_mode"`        // 付费模式类型
	SeriesPayTemplateId int64  `json:"series_pay_template_id"` // 付费模板ID
	TemplateName        string `json:"template_name"`          // 模板名称
	TemplateStatus      int    `json:"template_status"`        // 1 - 有效, 2 - 已删除
}
