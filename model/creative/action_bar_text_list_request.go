package creative

import (
	"encoding/json"
)

// ActionBarTextListRequest 获取行动号召按钮
type ActionBarTextListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignType 计划类型; 必填 2 - 提升应用安装;3 - 获取电商下单;4 - 推广品牌活动;5 - 收集销售线索;13 - 小店商品推广；14：直播推广
	CampaignType int `json:"campaign_type,omitempty"`
	// ConsultType 是否使用了咨询组件；0=未使用，1=使用；注，咨询组件仅在收集销售线索计划(campaign_type=5)下可用，且使用了咨询组件后，可用的行动号召按钮限于接口返回内容
	ConsultType int `json:"consult_type,omitempty"`
	// UnitMaterialType 广告标的物类型，campaign_type=32（微信小程序/小游戏）时必填，3=微信小程序；4=微信小游戏
	UnitMaterialType int `json:"unit_material_type"`
}

// Url implement GetRequest interface
func (r ActionBarTextListRequest) Url() string {
	return "v1/creative/action_bar_text/list"
}

// Encode implement GetRequest interface
func (r ActionBarTextListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
