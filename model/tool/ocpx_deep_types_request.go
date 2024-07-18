package tool

import "encoding/json"

// OcpxDeepTypesRequest 获取可选的深度优化目标
type OcpxDeepTypesRequest struct {
	// AdvertiserID 广告主ID，在获取 access_token 的时候返回，必填
	AdvertiserID uint64 `json:"advertiser_id"`

	// SceneCategory 场景分类，0=普通创编、1=追踪工具、2=常规托管、3=搜索广告，必填
	SceneCategory int `json:"scene_category"`

	// CampaignType 计划类型，在scene_category=1即追踪工具时可不填，其他场景必填，2：提升应用安装 3：获取电商下单 4：推广品牌活动 5：收集销售线索 7：提高应用活跃 9：商品库推广 16：粉丝/直播推广
	CampaignType int `json:"campaign_type"`

	// SceneIDs 投放资源位置，在scene_category=1即追踪工具时可不填，其他场景必填，1：优选广告位 5：联盟广告 6：上下滑大屏广告 10：联盟场景 24：激励视频 27：开屏广告位 39：搜索广告
	SceneIDs []int `json:"scene_ids"`

	// TraceutilType 追踪工具的类型，在scene_category=1即追踪工具时必填，其他场景不填，1=js、2=xpath、3=sdk、7=api、12=kwai attribution
	TraceutilType int `json:"traceutil_type"`

	// AppID 应用ID，可选，追踪工具类型时不需要本参数；其他类型时，本参数会决定可拉到的优化目标列表
	AppID uint64 `json:"app_id"`

	// OcpxActionType 深度目标是否可用是关联与浅度目标的，在scene_category=1即追踪工具时不填，其他场景可填，即单元创建时的ocpx_action_type字段
	OcpxActionType int `json:"ocpx_action_type"`

	// BidType 出价方式，10 - OCPM, 12 - MCB，可选
	BidType int `json:"bid_type"`

	// HostingScene 智能托管细分场景，0=常规推广、1=测书工具、2=账户预热、5=最大转化，当scene_category=2时本字段可用
	HostingScene int `json:"hosting_scene"`

	// SiteID 魔力建站的页面ID，按需传入，如果是收集销售线索（campaign_type=5）这类的落地页并且使用的是魔力建站的页面，可以传入本ID
	SiteID uint64 `json:"site_id"`
}

func (r OcpxDeepTypesRequest) Url() string {
	return "gw/dsp/v1/ocpx/deepTypes"
}

func (r OcpxDeepTypesRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
