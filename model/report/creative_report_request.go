package report

import "encoding/json"

// CreativeReportRequest 广告创意数据APIRequest
type CreativeReportRequest struct {
	ReportRequest
	// CampaignIDs 广告计划ID集，过滤筛选条件，单次查询数量不超过5000
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// CampaignType 计划类型，过滤筛选条件1 - 作品推广；2 - 提升应用安装；3 - 获取电商下单；4 - 推广品牌活动；5 - 收集销售线索；6 - 保量广告；7 - 提高应用活跃。
	CampaignType int `json:"campaign_type,omitempty"`
	// UnitIDs 广告组ID集，过滤筛选条件，单次查询数量不超过5000
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// CreativeIDs 广告创意ID集，过滤筛选条件，单次查询数量不超过5000
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	//获取额外信息， "photo"获取视频信息，包含视频Id与视频md5
	ExtendInfo []string `json:"extend_info,omitempty"`
}

// Url implement PostRequest interface
func (r CreativeReportRequest) Url() string {
	return "v1/report/creative_report"
}

// Encode implement PostRequest interface
func (r CreativeReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
