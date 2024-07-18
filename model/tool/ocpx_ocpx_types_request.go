package tool

import "encoding/json"

// OcpxOcpxTypesRequest 获取可选的浅度优化目标
type OcpxOcpxTypesRequest struct {
	// AdvertiserID 广告主ID，在获取 access_token 的时候返回，必填
	AdvertiserID uint64 `json:"advertiser_id"`

	// SceneCategory 场景分类，0=普通创编、1=追踪工具、2=常规托管（智能模式）、3=搜索广告，必填
	SceneCategory int `json:"scene_category"`

	// CampaignType 计划类型，在scene_category=1即追踪工具时可不填，其他场景必填，即新建计划接口的type参数
	CampaignType int `json:"campaign_type"`

	// SceneIDs 投放资源位置，在scene_category=1即追踪工具时可不填，其他场景必填，即新建单元接口的scene_id参数
	SceneIDs []int `json:"scene_ids"`

	// TraceutilType 追踪工具的类型，在scene_category=1即追踪工具时必填，其他场景不填，1=js、2=xpath、3=sdk、7=api、12=kwai attribution
	TraceutilType int `json:"traceutil_type"`

	// AppID 应用ID，可选，追踪工具类型时不需要本参数；其他类型时，本参数会决定可拉到的优化目标列表
	AppID uint64 `json:"app_id"`

	// BidType 出价类型，10=ocpm，12=mcb（计划层级选择最大转化类型时），当scene_category=0时本字段可用
	BidType int `json:"bid_type"`

	// HostingScene 智能托管细分场景，0=常规推广、1=测书工具、2=账户预热、5=最大转化，当scene_category=2时本字段可用
	HostingScene int `json:"hosting_scene"`

	// SiteID 魔力建站的页面ID，按需传入，如果是收集销售线索（campaign_type=5）这类的落地页并且使用的是魔力建站的页面，可以传入本ID
	SiteID uint64 `json:"site_id"`

	// UnitMaterialType 小程序单元组件类型，按需传入，1=快手小程序、2=快手小游戏、3=微信小程序、4=微信小游戏
	UnitMaterialType int `json:"unit_material_type"`
}

func (r OcpxOcpxTypesRequest) Url() string {
	return "gw/dsp/v1/ocpx/ocpxTypes"
}

func (r OcpxOcpxTypesRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
