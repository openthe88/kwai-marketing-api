package native

import (
	"encoding/json"
	"github.com/bububa/kwai-marketing-api/model"
)

type DetailedReportRequest struct {
	AdvertiserID uint64          `json:"advertiser_id"`          // 广告主ID
	SearchParam  *SearchParam    `json:"search_param,omitempty"` // 查询参数
	PageInfo     *model.PageInfo `json:"page_info,omitempty"`    // 分页信息
}

type SearchParam struct {
	ViewType       int   `json:"view_type"`
	ReportStartDay int64 `json:"report_start_day"`
	ReportEndDay   int64 `json:"report_end_day"`
}

// Url implement PostRequest interface
func (r DetailedReportRequest) Url() string {
	return "gw/dsp/v1/effect/native/detailedReport"
}

// Encode implement PostRequest interface
func (r DetailedReportRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type DetailedReportResponse struct {
	ResultList []AdDspNativeReportViewSnake `json:"result_list"` // 报表明细
	Sum        []AdDspNativeReportViewSnake `json:"sum"`         // 全局汇总
	PageInfo   *model.PageInfo              `json:"page_info"`   // 分页信息
}

// AdDspNativeReportViewSnake 定义广告数据结构体
type AdDspNativeReportViewSnake struct {
	CampaignID                                  int64   `json:"campaign_id" dc:"计划ID"`
	CampaignName                                string  `json:"campaign_name" dc:"计划名称"`
	UnitID                                      int64   `json:"unit_id" dc:"广告组ID"`
	UnitName                                    string  `json:"unit_name" dc:"广告组名称"`
	CreativeID                                  int64   `json:"creative_id" dc:"创意ID"`
	CreativeName                                string  `json:"creative_name" dc:"创意名称"`
	CreateTime                                  int     `json:"create_time" dc:"创建时间 时间戳 毫秒"`
	Roi                                         bool    `json:"roi" dc:"ROI"`
	IndirectSubmit0DCnt                         float64 `json:"indirect_submit_0_d_cnt" dc:"当日间接表单提交"`
	IndirectSubmit7DCnt                         float64 `json:"indirect_submit_7_d_cnt" dc:"7日间接表单提交"`
	T0IndirectConversionCnt                     float64 `json:"t_0_indirect_conversion_cnt" dc:"当日间接激活数"`
	T7IndirectConversionCnt                     float64 `json:"t_7_indirect_conversion_cnt" dc:"7日间接激活数"`
	T0IndirectPaiedCnt                          float64 `json:"t_0_indirect_paied_cnt" dc:"当日间接付费次数"`
	T7IndirectPaiedCnt                          float64 `json:"t_7_indirect_paied_cnt" dc:"7日间接付费次数"`
	T0IndirectPaiedAmt                          float64 `json:"t_0_indirect_paied_amt" dc:"当日间接付费金额"`
	T7IndirectPaiedAmt                          float64 `json:"t_7_indirect_paied_amt" dc:"7日间接付费金额"`
	T7IndirectedPaiedRoi                        float64 `json:"t_7_indirected_paied_roi" dc:"7日间接ROI"`
	TotalCharge                                 int64   `json:"total_charge" dc:"花费(厘)"`
	Impression                                  int64   `json:"impression" dc:"封面曝光数"`
	PhotoClick                                  int64   `json:"photo_click" dc:"封面点击数"`
	Click                                       int64   `json:"click" dc:"素材曝光数"`
	ActionbarClick                              int64   `json:"actionbar_click" dc:"行为数"`
	PhotoClickRatio                             float64 `json:"photo_click_ratio" dc:"封面点击率"`
	Impression1KCost                            float64 `json:"impression_1_k_cost" dc:"平均千次封面曝光花费(币)"`
	Click1KCost                                 float64 `json:"click_1_k_cost" dc:"平均千次素材曝光花费(币)"`
	PhotoClickCost                              float64 `json:"photo_click_cost" dc:"平均封面点击单价(币)"`
	ActionCost                                  float64 `json:"action_cost" dc:"平均行为单价(币)"`
	AdShow                                      int64   `json:"ad_show" dc:"曝光数"`
	ActionNewRatio                              float64 `json:"action_new_ratio" dc:"行为率"`
	AdPhotoPlayed2SRatio                        float64 `json:"ad_photo_played_2_s_ratio" dc:"2s播放率"`
	Play3SRatio                                 float64 `json:"play_3_s_ratio" dc:"3s播放率"`
	Play5SRatio                                 float64 `json:"play_5_s_ratio" dc:"5s播放率"`
	AdPhotoPlayed10SRatio                       float64 `json:"ad_photo_played_10_s_ratio" dc:"10s播放率"`
	AdPhotoPlayed75PercentRatio                 float64 `json:"ad_photo_played_75_percent_ratio" dc:"75%进度播放率"`
	PlayEndRatio                                float64 `json:"play_end_ratio" dc:"完播率"`
	Share                                       int64   `json:"share" dc:"分享数"`
	Comment                                     int64   `json:"comment" dc:"评论数"`
	Likes                                       int64   `json:"likes" dc:"点赞数"`
	Follow                                      int64   `json:"follow" dc:"新增关注数"`
	Report                                      int64   `json:"report" dc:"举报数"`
	Block                                       int64   `json:"block" dc:"拉黑数"`
	Negative                                    float64 `json:"negative" dc:"不感兴趣数(减少此类作品数)"`
	PlayedNum                                   int64   `json:"played_num" dc:"播放数"`
	PlayedThreeSeconds                          int64   `json:"played_three_seconds" dc:"3s播放数(有效播放数)"`
	AdPhotoPlayed10S                            int64   `json:"ad_photo_played_10_s" dc:"10s播放数"`
	AdPhotoPlayed75Percent                      int64   `json:"ad_photo_played_75_percent" dc:"75%播放进度数"`
	PlayedEnd                                   int64   `json:"played_end" dc:"完播数(播放完成)"`
	Conversion                                  int64   `json:"conversion" dc:"激活数"`
	ConversionCost                              float64 `json:"conversion_cost" dc:"激活单价"`
	KeyAction                                   int64   `json:"key_action" dc:"关键行为数"`
	KeyActionCost                               float64 `json:"key_action_cost" dc:"关键行为成本"`
	KeyActionRatio                              float64 `json:"key_action_ratio" dc:"关键行为率"`
	EventNewUserPay                             int64   `json:"event_new_user_pay" dc:"新增付费人数"`
	EventNewUserPayCost                         float64 `json:"event_new_user_pay_cost" dc:"新增付费人数成本"`
	EventNewUserPayRatio                        float64 `json:"event_new_user_pay_ratio" dc:"新增付费人数率"`
	EventPayFirstDay                            int64   `json:"event_pay_first_day" dc:"首日付费次数"`
	EventPayFirstDayCost                        float64 `json:"event_pay_first_day_cost" dc:"首日付费次数成本"`
	EventPayPurchaseAmountFirstDay              float64 `json:"event_pay_purchase_amount_first_day" dc:"激活当日付费金额"`
	EventPayFirstDayRoi                         float64 `json:"event_pay_first_day_roi" dc:"激活当日ROI"`
	EventPayPurchaseAmountOneDay                float64 `json:"event_pay_purchase_amount_one_day" dc:"激活后24h付费金额(回传时间)"`
	EventPayPurchaseAmountOneDayRoi             float64 `json:"event_pay_purchase_amount_one_day_roi" dc:"激活后24h-ROI(回传时间)"`
	EventPayPurchaseAmountOneDayByConversion    float64 `json:"event_pay_purchase_amount_one_day_by_conversion" dc:"激活后24h付费金额(激活时间)"`
	EventPayPurchaseAmountOneDayByConversionRoi float64 `json:"event_pay_purchase_amount_one_day_by_conversion_roi" dc:"激活后24h-ROI(激活时间)"`
	EventPay                                    int64   `json:"event_pay" dc:"付费次数"`
	EventPayCost                                float64 `json:"event_pay_cost" dc:"付费次数成本"`
	EventPayPurchaseAmount                      float64 `json:"event_pay_purchase_amount" dc:"付费金额"`
	T0DirectConversionCnt                       int64   `json:"t_0_direct_conversion_cnt" dc:"激活数(计费时间)"`
	T0ConversionCnt                             int64   `json:"t_0_conversion_cnt" dc:"当日累计激活数"`
	T7ConversionCnt                             int64   `json:"t_7_conversion_cnt" dc:"7日累计激活数"`
	T0DirectPaiedCnt                            int64   `json:"t_0_direct_paied_cnt" dc:"付费次数(计费时间)"`
	T0PaiedCnt                                  int64   `json:"t_0_paied_cnt" dc:"当日累计付费次数"`
	T7PaiedCnt                                  int64   `json:"t_7_paied_cnt" dc:"7日累计付费次数"`
	AccumulatedPaiedCost                        float64 `json:"accumulated_paied_cost" dc:"累计付费次数成本"`
	T0DirectPaiedAmt                            float64 `json:"t_0_direct_paied_amt" dc:"付费金额(计费时间)"`
	T0PaiedAmt                                  float64 `json:"t_0_paied_amt" dc:"当日累计付费金额"`
	T7PaiedAmt                                  float64 `json:"t_7_paied_amt" dc:"7日累计付费金额"`
	T0PaiedRoi                                  float64 `json:"t_0_paied_roi" dc:"当日累计ROI"`
	T7PaiedRoi                                  float64 `json:"t_7_paied_roi" dc:"7日累计ROI"`
	ConversionNum                               int64   `json:"conversion_num" dc:"转化数(回传时间)"`
	ConversionNumCost                           float64 `json:"conversion_num_cost" dc:"转化成本(回传时间)"`
	ConversionRatio                             float64 `json:"conversion_ratio" dc:"转化率(回传时间)"`
	DeepConversionNum                           int64   `json:"deep_conversion_num" dc:"深度转化数(回传时间)"`
	DeepConversionCost                          float64 `json:"deep_conversion_cost" dc:"深度转化成本(回传时间)"`
	DeepConversionRatio                         float64 `json:"deep_conversion_ratio" dc:"深度转化率(回传时间)"`
	DirectSubmit1DCnt                           int64   `json:"direct_submit_1_d_cnt" dc:"表单提交数(计费时间)"`
	Submit1DCnt                                 int64   `json:"submit_1_d_cnt" dc:"当日累计表单提交"`
	Submit7DCnt                                 int64   `json:"submit_7_d_cnt" dc:"7日累计表单提交"`
	SubmitUnitPriceCost                         float64 `json:"submit_unit_price_cost" dc:"累计表单提交单价"`
	EventValidClues                             int64   `json:"event_valid_clues" dc:"有效线索数"`
	EventValidClueCost                          float64 `json:"event_valid_clue_cost" dc:"有效线索成本"`
	MerchantRecoFans                            int64   `json:"merchant_reco_fans" dc:"涨粉数"`
	MerchantRecoFansCost                        float64 `json:"merchant_reco_fans_cost" dc:"涨粉成本"`
	IndirectEventPayCnt                         int64   `json:"indirect_event_pay_cnt" dc:"间接转化数(回传时间)"`
	KolUserTypeDesc                             string  `json:"kol_user_type_desc" dc:"原生广告类型"`
	OcpcActionType                              string  `json:"ocpc_action_type" dc:"优化目标"`
	CampaignType                                string  `json:"campaign_type" dc:"营销目标"`
	AuthorId                                    string  `json:"author_id" dc:"快手号ID"`
}
