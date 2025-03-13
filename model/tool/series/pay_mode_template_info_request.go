package series

import "encoding/json"

// PayModeTemplateInfoRequest 查询短剧付费模板信息接口
type PayModeTemplateInfoRequest struct {
	// AdvertiserID 广告主ID（必填）
	AdvertiserID int64 `json:"advertiser_id"`

	// SeriesID 短剧ID（必填）
	SeriesID int64 `json:"series_id"`

	// SeriesPayMode 付费模式类型（必填，目前仅支持“打包”，即 payMode=1）
	SeriesPayMode int `json:"series_pay_mode"`

	// SeriesPayTemplateID 付费模版ID（非必填，填写则代表查询该模板ID对应信息）注:指针参数表示非必传
	SeriesPayTemplateID *int64 `json:"series_pay_template_id"`

	// UserID 短剧作者ID（必填）
	UserID int64 `json:"user_id"`
}

func (r PayModeTemplateInfoRequest) Url() string {
	return "gw/dsp/series/payModeTemplateInfo"
}

func (r PayModeTemplateInfoRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
