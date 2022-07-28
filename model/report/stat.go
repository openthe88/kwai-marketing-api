package report

// Stat 数据报表
type Stat struct {
	// CampaignID 广告计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 广告计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// UnitID 广告组ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitName 广告组名称
	UnitName string `json:"unit_name,omitempty"`
	// CreativeID 广告创意ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// CreativeName 广告创意名称
	CreativeName string `json:"creative_name,omitempty"`
	// PhotoID 视频id
	PhotoID string `json:"photo_id,omitempty"`
	// PhotoUrl 视频链接
	PhotoUrl string `json:"photo_url,omitempty"`
	// ImageToken 封面id
	ImageToken string `json:"image_token,omitempty"`
	// CoverUrl 封面链接
	CoverUrl string `json:"cover_url,omitempty"`
	// Description 作品广告语
	Description string `json:"description,omitempty"`
	// StickyUrl 便利贴url地址（便利贴报表时返回）
	StickyUrl string `json:"sticky_url,omitempty"`
	// ImageUrls 图片url地址（图片报表时返回）
	ImageUrls string `json:"image_urls,omitempty"`
	// AdScene 广告投放场景
	AdScene string `json:"ad_scene,omitempty"`
	// Province 省份
	Province string `json:"province,omitempty"`
	// City 城市
	City string `json:"city,omitempty"`
	// Gender 性别
	Gender string `json:"gender,omitempty"`
	// AgeSegment 年龄
	AgeSegment string `json:"age_segment,omitempty"`
	// Client 系统
	Client string `json:"client,omitempty"`
	// BusinessInterestTags 商业兴趣
	BusinessInterestTags string `json:"business_interest_tags,omitempty"`
	// Status 1 - 投放中；2 - 已暂停；3 - 已删除
	Status int `json:"status,omitempty"`
	// StatDate 数据日期，格式：YYYY-MM-DD
	StatDate string `json:"stat_date,omitempty"`
	// StatHour 数据小时
	StatHour int `json:"stat_hour,omitempty"`
	// Charge 花费（元）
	Charge float64 `json:"charge,omitempty"`
	// Show 封面曝光数
	Show int64 `json:"show,omitempty"`
	// PhotoClick 封面点击数
	PhotoClick int64 `json:"photo_click,omitempty"`
	// Aclick 素材曝光数
	Aclick int64 `json:"aclick,omitempty"`
	// Bclick 行为数
	Bclick int64 `json:"bclick,omitempty"`
	// PhotoClickRatio 封面点击率
	PhotoClickRatio float64 `json:"photo_click_ratio,omitempty"`
	// Play3sRatio 3s播放率
	Play3sRatio float64 `json:"play_3s_ratio,omitempty"`
	// Played3s 有效播放数
	Played3s int64 `json:"played_three_seconds,omitempty"`
	// ActionRatio 行为率
	ActionRatio float64 `json:"action_ratio,omitempty"`
	// Impression1kCost 平均千次曝光花费（元）
	Impression1kCost float64 `json:"impression_1k_cost,omitempty"`
	// PhotoClickCost 平均点击单价（元）
	PhotoClickCost float64 `json:"photo_click_cost,omitempty"`
	// ActionCost 平均行为单价（元）
	ActionCost float64 `json:"action_cost,omitempty"`
	// Share 分享数
	Share int64 `json:"share,omitempty"`
	// Comment 评论数
	Comment int64 `json:"comment,omitempty"`
	// Like 点赞数
	Like int64 `json:"like,omitempty"`
	// Follow 新增关注数
	Follow int64 `json:"follow,omitempty"`
	// CancelFollow 取消关注数
	CancelFollow int64 `json:"cancel_follow,omitempty"`
	// Report 举报数
	Report int64 `json:"report,omitempty"`
	// Block 拉黑数
	Block int64 `json:"block,omitempty"`
	// Negative 减少此类作品数
	Negative int64 `json:"negative,omitempty"`
	// Submit 提交按钮点击数（历史字段，同下方“表单提交数”，预计年底删除该字段）
	Submit int64 `json:"submit,omitempty"`
	// DownloadStarted 应用下载数据-安卓下载开始数
	DownloadStarted int64 `json:"download_started,omitempty"`
	// DownloadCompleted 应用下载数据-安卓下载完成数
	DownloadCompleted int64 `json:"download_completed,omitempty"`
	// Activation 应用下载数据-激活数
	Activation int64 `json:"activation,omitempty"`
	// EventPayFirstDay 应用下载数据-首日付费次数
	EventPayFirstDay int64 `json:"event_pay_first_day,omitempty"`
	// EventPayPurchaseAmountFirstDay 应用下载数据-首日付费金额
	EventPayPurchaseAmountFirstDay float64 `json:"event_pay_purchase_amount_first_day,omitempty"`
	// EventPayFirstDayRoi 应用下载数据-首日ROI
	EventPayFirstDayRoi float64 `json:"event_pay_first_day_roi,omitempty"`
	// EventPay 应用下载数据-付费次数
	EventPay int64 `json:"event_pay,omitempty"`
	// EventPayPurchaseAmount 应用下载数据-付费金额
	EventPayPurchaseAmount float64 `json:"event_pay_purchase_amount,omitempty"`
	// EventPayRoi 应用下载数据-ROI
	EventPayRoi float64 `json:"event_pay_roi,omitempty"`
	// EventRegister 应用下载数据-注册数
	EventRegister int64 `json:"event_register,omitempty"`
	// EventRegisterCost 应用下载数据-注册成本
	EventRegisterCost float64 `json:"event_register_cost,omitempty"`
	// EventRegisterRatio 应用下载数据-注册率
	EventRegisterRatio float64 `json:"event_register_ratio,omitempty"`
	// EventJinJianApp 应用下载数据-完件数
	EventJinJianApp int64 `json:"event_jin_jian_app,omitempty"`
	// EventJinJianAppCost 应用下载数据-完件成本
	EventJinJianAppCost float64 `json:"event_jin_jian_app_cost,omitempty"`
	// EventCreditGrantApp 应用下载数据-授信数
	EventCreditGrantApp int64 `json:"event_credit_grant_app,omitempty"`
	// EventCreditGrantAppCost 应用下载数据-授信成本
	EventCreditGrantAppCost float64 `json:"event_credit_grant_app_cost,omitempty"`
	// EventCreditGrantAppRatio 应用下载数据-授信率
	EventCreditGrantAppRatio float64 `json:"event_credit_grant_app_ratio,omitempty"`
	// EventOrderPaid 应用下载数据-付款成功数
	EventOrderPaid int64 `json:"event_order_paid,omitempty"`
	// EventOrderPaidPurchaseAmount 应用下载数据-付款成功金额
	EventOrderPaidPurchaseAmount float64 `json:"event_order_paid_purchase_amount,omitempty"`
	// EventOrderPaidCost 应用下载数据-单次付款成本
	EventOrderPaidCost float64 `json:"event_order_paid_cost,omitempty"`
	// EventNextDayStay 应用下载数据-次留数（仅支持分日查询）
	EventNextDayStay int64 `json:"event_next_day_stay,omitempty"`
	// EventNextDayStayCost 应用下载数据-次留成本（仅支持分日查询）
	EventNextDayStayCost float64 `json:"event_next_day_stay_cost,omitempty"`
	// EventNextDayStayRatio 应用下载数据-次留率（仅支持分日查询）
	EventNextDayStayRatio float64 `json:"event_next_day_stay_ratio,omitempty"`
	// FormCount 落地页数据-表单提交数
	FormCount int64 `json:"form_count,omitempty"`
	// FormCost 落地页数据-表单提交单价
	FormCost float64 `json:"form_cost,omitempty"`
	// FormActionRatio 落地页数据-表单提交点击率
	FormActionRatio float64 `json:"form_action_ratio,omitempty"`
	// LivePlayed3s 直播观看数
	LivePlayed3s int64 `json:"live_played_3s,omitempty"`
	// EventMakingCalls 电话拨打数-用户点击电话按钮的次数
	EventMakingCalls int64 `json:"event_making_calls,omitempty"`
	// EventPayWeightedPurchaseAmount 加权付费金额-当日回传的付费行为所带来的加权付费金额，单位:元，当前用于保险行业
	EventPayWeightedPurchaseAmount float64 `json:"event_pay_weighted_purchase_amount,omitempty"`
	// EventPayWeightedPurchaseAmountFirstDay 首日加权付费金额-当日激活的用户在当天产生的付费行为所带来的加权付费金额 单位:元，当前用于保险行业
	EventPayWeightedPurchaseAmountFirstDay float64 `json:"event_pay_weighted_purchase_amount_first_day,omitempty"`
	// EventConsultationValidRetained 留咨咨询数
	EventConsultationValidRetained int64 `json:"event_consultation_valid_retained,omitempty"`
	// EventPreComponentConsultationValidRetained 附加咨询组件留资咨询数
	EventPreComponentConsultationValidRetained int64 `json:"event_pre_component_consultation_valid_retained,omitempty"`
	// EventWechatQrCodeLinkClick 微信小程序深度加粉数
	EventWechatQrCodeLinkClick int64 `json:"event_wechat_qr_code_link_click,omitempty"`
	// LiveEventGoodsView 直播间商品点击数
	LiveEventGoodsView int64 `json:"live_event_goods_view,omitempty"`
	// CancelLike 取消点赞数
	CancelLike int64 `json:"cancel_like,omitempty"`
	// DownloadInstalled 安卓安装完成数
	DownloadInstalled int64 `json:"download_installed,omitempty"`
	// EventAudition 首次试听到课数
	EventAudition int64 `json:"event_audition,omitempty"`
	// EventJinJianLandingPage 落地页数据-落地页完件数
	EventJinJianLandingPage int64 `json:"event_jin_jian_landing_page,omitempty"`
	// EventJinJianLandingPageCost 落地页数据-落地页完件成本
	EventJinJianLandingPageCost float64 `json:"event_jin_jian_landing_page_cost,omitempty"`
	// EventCreditGrantLandingPage 落地页数据-落地页授信数
	EventCreditGrantLandingPage int64 `json:"event_credit_grant_landing_page,omitempty"`
	// EventCreditGrantLandingPageCost 落地页数据-落地页授信成本
	EventCreditGrantLandingPageCost float64 `json:"event_credit_grant_landing_page_cost,omitempty"`
	// EventCreditGrantLandingPageRatio 落地页数据-落地页授信率
	EventCreditGrantLandingPageRatio float64 `json:"event_credit_grant_landing_page_ratio,omitempty"`
	// EventValidClues 落地页数据-有效线索数
	EventValidClues int64 `json:"event_valid_clues,omitempty"`
	// EventValidCluesCost 落地页数据-有效线索成本
	EventValidCluesCost float64 `json:"event_valid_clues_cost,omitempty"`
	// Click1KCost 平均千次素材曝光花费(元)
	Click1KCost float64 `json:"click_1k_cost,omitempty"`
	// EventCreditGrantLandingRatio 落地页授信率
	EventCreditGrantLandingRatio float64 `json:"event_credit_grant_landing_ratio,omitempty"`
	// AdProductCnt 商品成交数
	AdProductCnt int64 `json:"ad_product_cnt,omitempty"`
	// EventGoodsView 商品访问数
	EventGoodsView int64 `json:"event_goods_view,omitempty"`
	// MerchantRecoFans 涨粉量
	MerchantRecoFans int64 `json:"merchant_reco_fans,omitempty"`
	// EventOrderAmountRoi 涨粉量
	EventOrderAmountRoi float64 `json:"event_order_amount_roi,omitempty"`
	// EventGoodsViewCost 商品访问成本
	EventGoodsViewCost float64 `json:"event_goods_view_cost,omitempty"`
	// MerchantRecoFansCost 涨粉成本
	MerchantRecoFansCost float64 `json:"merchant_reco_fans_cost,omitempty"`
	// EventNewUserPay 新增付费人数
	EventNewUserPay int64 `json:"event_new_user_pay,omitempty"`
	// EventNewUserPayCost 新增付费人数成本：当日消耗 / 当日新增付费人数
	EventNewUserPayCost float64 `json:"event_new_user_pay_cost,omitempty"`
	// EventNewUserPayRatio 新增付费人数率：当日新增付费人数 / 当日行为数
	EventNewUserPayRatio float64 `json:"event_new_user_pay_ratio,omitempty"`
	// EventButtonClick 按钮点击数
	EventButtonClick int64 `json:"event_button_click,omitempty"`
	// EventButtonClickCost 按钮点击成本：当日消耗 / 按钮点击数
	EventButtonClickCost float64 `json:"event_button_click_cost,omitempty"`
	// EventButtonClickRatio 按钮点击成本：当日消耗 / 按钮点击数
	EventButtonClickRatio float64 `json:"event_button_click_ratio,omitempty"`
	// Play5sRatio 5s播放率
	Play5sRatio float64 `json:"play_5s_ratio,omitempty"`
	// PlayEndRatio 完播率
	PlayEndRatio float64 `json:"play_end_ratio,omitempty"`
	// EventOrderPaidRoi 订单支付率：订单支付数/商品访问数
	EventOrderPaidRoi float64 `json:"event_order_paid_roi,omitempty"`
	// EventNewUserJinjianApp 新增完件人数
	EventNewUserJinjianApp int64 `json:"event_new_user_jinjian_app,omitempty"`
	// EventNewUserJinjianAppCost 新增完件人数成本：消耗/新增完件人数
	EventNewUserJinjianAppCost float64 `json:"event_new_user_jinjian_app_cost,omitempty"`
	// EventNewUserJinjianAppRoi 新增完件人数率：新增完件人数/（转化+表单提交）
	EventNewUserJinjianAppRoi float64 `json:"event_new_user_jinjian_app_roi,omitempty"`
	// EventNewUserCreditGrantApp 新增授信人数
	EventNewUserCreditGrantApp int64 `json:"event_new_user_credit_grant_app,omitempty"`
	// EventNewUserCreditGrantAppCost 新增授信人数成本：消耗/新增授信人数
	EventNewUserCreditGrantAppCost float64 `json:"event_new_user_credit_grant_app_cost,omitempty"`
	// EventNewUserCreditGrantAppRoi 新增授信人数率：新增授信人数/（转化+表单提交）
	EventNewUserCreditGrantAppRoi float64 `json:"event_new_user_credit_grant_app_roi,omitempty"`
	// EventNewUserJinjianPage 落地页新增完件人数
	EventNewUserJinjianPage int64 `json:"event_new_user_jinjian_page,omitempty"`
	// EventNewUserJinjianPageCost 落地页新增完件人数成本
	EventNewUserJinjianPageCost float64 `json:"event_new_user_jinjian_page_cost,omitempty"`
	// EventNewUserJinjianPageRoi 落地页新增完件人数率
	EventNewUserJinjianPageRoi float64 `json:"event_new_user_jinjian_page_roi,omitempty"`
	// EventNewUserCreditGrantPage 落地页新增授信人数
	EventNewUserCreditGrantPage int64 `json:"event_new_user_credit_grant_page,omitempty"`
	// EventNewUserCreditGrantPageCost 落地页新增授信人数成本
	EventNewUserCreditGrantPageCost float64 `json:"event_new_user_credit_grant_page_cost,omitempty"`
	// EventNewUserCreditGrantPageRoi 落地页新增授信人数率
	EventNewUserCreditGrantPageRoi float64 `json:"event_new_user_credit_grant_page_roi,omitempty"`
	// EventAddWechat 微信复制数
	EventAddWechat int64 `json:"event_add_wechat,omitempty"`
	// EventAddWechatCost 微信复制成本
	EventAddWechatCost float64 `json:"event_add_wechat_cost,omitempty"`
	// EventAddWechatRatio 微信复制率
	EventAddWechatRatio float64 `json:"event_add_wechat_ratio,omitempty"`
	// EventGetThrough 智能电话-确认接通数
	EventGetThrough int64 `json:"event_get_through,omitempty"`
	// EventGetThroughCost 智能电话-确认接通成本
	EventGetThroughCost float64 `json:"event_get_through_cost,omitempty"`
	// EventGetThroughRatio 智能电话-确认接通率
	EventGetThroughRatio float64 `json:"event_get_through_ratio,omitempty"`
	// AdScene2 .
	AdScene2 string `json:"adScene,omitempty"`
	// PlacementType 广告范围
	PlacementType string `json:"placement_type,omitempty"`
	// EventAppointForm 预约表单数
	EventAppointForm int64 `json:"event_appoint_form,omitempty"`
	// EventAppointFormCost 预约表单点击成本
	EventAppointFormCost float64 `json:"event_appoint_form_cost,omitempty"`
	// EventAppointFormRatio 预约表单点击率
	EventAppointFormRatio float64 `json:"event_appoint_form_ratio,omitempty"`
	// EventAppointJumpClick 预约跳转点击数
	EventAppointJumpClick int64 `json:"event_appoint_jump_click,omitempty"`
	// EventAppointJumpClickCost 预约跳转点击成本
	EventAppointJumpClickCost float64 `json:"event_appoint_jump_click_cost,omitempty"`
	// EventAppointJumpClickRatio 预约跳转点击率
	EventAppointJumpClickRatio float64 `json:"event_appoint_jump_click_ratio,omitempty"`
	// UnionEventPayPurchaseAmount7d 联盟广告收入
	UnionEventPayPurchaseAmount7d float64 `json:"union_event_pay_purchase_amount_7d,omitempty"`
	// UnionEventPayPurchaseAmount7dRoi 联盟变现ROI
	UnionEventPayPurchaseAmount7dRoi float64 `json:"union_event_pay_purchase_amount_7d_roi,omitempty"`
	// EventDspGiftForm 附加组件表单提交
	EventDspGiftForm int64 `json:"event_dsp_gift_form,omitempty"`
	// EventAppInvoked 唤醒应用数
	EventAppInvoked int64 `json:"event_app_invoked,omitempty"`
	// EventAppInvokedCost 唤醒应用成本
	EventAppInvokedCost float64 `json:"event_app_invoked_cost,omitempty"`
	// EventAppInvokedRatio 唤醒应用率
	EventAppInvokedRatio float64 `json:"event_app_invoked_ratio,omitempty"`
	// EventMultiConversion 落地页多转化次数
	EventMultiConversion int64 `json:"event_multi_conversion,omitempty"`
	// EventMultiConversionRatio 落地页多转化率
	EventMultiConversionRatio float64 `json:"event_multi_conversion_ratio,omitempty"`
	// EventMultiConversionCost 落地页多转化成本
	EventMultiConversionCost float64 `json:"event_multi_conversion_cost,omitempty"`
	// EventWatchAppAd 广告观看
	EventWatchAppAd int64 `json:"event_watch_app_ad,omitempty"`
	// EventAdWatchTimes 广告观看次数
	EventAdWatchTimes int64 `json:"event_ad_watch_times,omitempty"`
	// EventAdWatchTimesRatio 广告观看次数转化率
	EventAdWatchTimesRatio float64 `json:"event_ad_watch_times_ratio,omitempty"`
	// EventAdWatchTimesCost 广告观看次数成本
	EventAdWatchTimesCost float64 `json:"event_ad_watch_times_cost,omitempty"`
	// EventAdWatch5Times 5次广告广告观看数
	EventAdWatch5Times int64 `json:"event_ad_watch_5_times,omitempty"`
	// EventAdWatch10Times 10次广告广告观看数
	EventAdWatch10Times int64 `json:"event_ad_watch_10_times,omitempty"`
	// EventAdWatch20Times 20次广告广告观看数
	EventAdWatch20Times int64 `json:"event_ad_watch_20_times,omitempty"`
	// EventAddShoppingCart 添加购物车数
	EventAddShoppingCart int64 `json:"event_add_shopping_cart,omitempty"`
	// EventAddShoppingCartCost 添加购物车成本
	EventAddShoppingCartCost float64 `json:"event_add_shopping_cart_cost,omitempty"`
	// AdPhotoPlayed75Percent 75%进度播放数
	AdPhotoPlayed75Percent int64 `json:"ad_photo_played_75percent,omitempty"`
	// AdPhotoPlayed10s 10s播放数
	AdPhotoPlayed10s int64 `json:"ad_photo_played_10s,omitempty"`
	// AdPhotoPlayed2s 2s播放数
	AdPhotoPlayed2s int64 `json:"ad_photo_played_2s,omitempty"`
	// AdPhotoPlayed75PercentRatio 75%进度播放率
	AdPhotoPlayed75PercentRatio float64 `json:"ad_photo_played_75percent_ratio,omitempty"`
	// AdPhotoPlayed10sRatio 10s播放率
	AdPhotoPlayed10sRatio float64 `json:"ad_photo_played_10s_ratio,omitempty"`
	// AdPhotoPlayed2sRatio 2s播放率
	AdPhotoPlayed2sRatio float64 `json:"ad_photo_played_2s_ratio,omitempty"`
	// EventPhoneGetThrough 电话建联数
	EventPhoneGetThrough int64 `json:"event_phone_get_through,omitempty"`
	// EventIntentionConfirmed 意向确认数
	EventIntentionConfirmed int64 `json:"event_intention_confirmed,omitempty"`
	// EventWechatConnected 微信加粉数
	EventWechatConnected int64 `json:"event_wechat_connected,omitempty"`
	// EventOrderSubmit 提交订单数
	EventOrderSubmit int64 `json:"event_order_submit,omitempty"`
	// EventOrderSuccessed 有效线索成交数
	EventOrderSuccessed int64 `json:"event_order_successed,omitempty"`
	// EventPhoneCardActivate 电话卡激活数
	EventPhoneCardActivate int64 `json:"event_phone_card_activate,omitempty"`
	// EventMeasurementHouse 量房数
	EventMeasurementHouse int64 `json:"event_measurement_house,omitempty"`
	// ActionNewRatio 行为率 新
	ActionNewRatio float64 `json:"action_new_ratio,omitempty"`
	// EventOutboundCall 电话拨打数
	EventOutboundCall int64 `json:"event_outbound_call,omitempty"`
	// EventOutboundCallCost 电话拨打成本
	EventOutboundCallCost float64 `json:"event_outbound_call_cost,omitempty"`
	// EventOutboundCallRatio 电话拨打率
	EventOutboundCallRatio float64 `json:"event_outbound_call_ratio,omitempty"`
	// KeyAction 关键行为数
	KeyAction int64 `json:"key_action,omitempty"`
	// KeyActionCost 关键行为成本
	KeyActionCost float64 `json:"key_action_cost,omitempty"`
	// KeyActionRatio 关键行为率
	KeyActionRatio float64 `json:"key_action_ratio,omitempty"`
	// EventCreditCardRecheck 信用卡核卡数
	EventCreditCardRecheck int64 `json:"event_credit_card_recheck,omitempty"`
	// EventCreditCardRecheckFirstDay 信用卡首日核卡数
	EventCreditCardRecheckFirstDay int64 `json:"event_credit_card_recheck_first_day,omitempty"`
	// EventNoIntention 用户无意向数
	EventNoIntention int64 `json:"event_no_intention,omitempty"`
	// BackflowCvLower 回流预估转化数下限
	BackflowCvLower int `json:"backflow_cv_lower,omitempty"`
	// BackflowCvUpper 回流预估转化数上限
	BackflowCvUpper int `json:"backflow_cv_upper,omitempty"`
	// ConversionCostLower 回流预估转化成本下限
	ConversionCostLower float64 `json:"conversion_cost_lower,omitempty"`
	// ConversionCostUpper 回流预估转化成本上限
	ConversionCostUpper float64 `json:"conversion_cost_upper,omitempty"`
	// BackflowPayment 首日付费总金额，单位元
	BackflowPayment float64 `json:"backflow_payment,omitempty"`
	// BackflowRoi 首日ROI
	BackflowRoi float64 `json:"backflow_roi,omitempty"`
	// BackflowSevenRoi 7日ROI
	BackflowSevenRoi float64 `json:"backflow_seven_roi,omitempty"`
	// ApproxPayCount 近似购买数
	ApproxPayCount int64 `json:"approx_pay_count,omitempty"`
	// ApproxPayCost 淘系近似购买成本
	ApproxPayCost float64 `json:"approx_pay_cost,omitempty"`
	// ApproxPayRatio 淘系近似购买率
	ApproxPayRatio float64 `json:"approx_pay_ratio,omitempty"`
	// ClickConversionRatio 点击激活成本
	ClickConversionRatio float64 `json:"click_conversion_ratio,omitempty"`
	// ConversionCost 激活单价
	ConversionCost float64 `json:"conversion_cost,omitempty"`
	// DownloadCompletedCost 安卓下载完成单价（元）
	DownloadCompletedCost float64 `json:"download_completed_cost,omitempty"`
	// DownloadCompletedRatio 安卓下载完成率
	DownloadCompletedRatio float64 `json:"download_completed_ratio,omitempty"`
	// DownloadConversionRatio 下载完成激活率
	DownloadConversionRatio float64 `json:"download_conversion_ratio,omitempty"`
	// DownloadStartedCost 安卓下载开始单价（元）
	DownloadStartedCost float64 `json:"download_started_cost,omitempty"`
	// DownloadStartedRatio 安卓下载开始率
	DownloadStartedRatio float64 `json:"download_started_ratio,omitempty"`
	// EventAdWatch5TimesCost 5次广告观看成本
	EventAdWatch5TimesCost float64 `json:"event_ad_watch_5_times_cost,omitempty"`
	// EventAdWatch5TimesRatio 5次广告观看转化率
	EventAdWatch5TimesRatio float64 `json:"event_ad_watch_5_times_ratio,omitempty"`
	// EventAdWatch10TimesCost 10次广告观看成本
	EventAdWatch10TimesCost float64 `json:"event_ad_watch_10_times_cost,omitempty"`
	// EventAdWatch10TimesRatio 10次广告观看转化率
	EventAdWatch10TimesRatio float64 `json:"event_ad_watch_10_times_ratio,omitempty"`
	// EventAdWatch20TimesCost 20次广告观看成本
	EventAdWatch20TimesCost float64 `json:"event_ad_watch_20_times_cost,omitempty"`
	// EventAdWatch20TimesRatio 20次广告观看转化率
	EventAdWatch20TimesRatio float64 `json:"event_ad_watch_20_times_ratio,omitempty"`
	// EventConsultationValidRetainedCost 留咨咨询成本
	EventConsultationValidRetainedCost float64 `json:"event_consultation_valid_retained_cost,omitempty"`
	// EventConsultationValidRetainedRatio 留咨咨询率
	EventConsultationValidRetainedRatio float64 `json:"event_consultation_valid_retained_ratio,omitempty"`
	// EventConversionClickCost 有效咨询成本
	EventConversionClickCost float64 `json:"event_conversion_click_cost,omitempty"`
	// EventConversionClickRatio 有效咨询率
	EventConversionClickRatio float64 `json:"event_conversion_click_ratio,omitempty"`
	// EventCreditGrantFirstDayAppCost 首日授信成本
	EventCreditGrantFirstDayAppCost float64 `json:"event_credit_grant_first_day_app_cost,omitempty"`
	// EventCreditGrantFirstDayAppRatio 首日授信率
	EventCreditGrantFirstDayAppRatio float64 `json:"event_credit_grant_first_day_app_ratio,omitempty"`
	// EventCreditGrantFirstDayLandingPage 落地页首日授信数
	EventCreditGrantFirstDayLandingPage int64 `json:"event_credit_grant_first_day_landing_page,omitempty"`
	// EventCreditGrantFirstDayLandingPageCost 落地页首日授信成本
	EventCreditGrantFirstDayLandingPageCost float64 `json:"event_credit_grant_first_day_landing_page_cost,omitempty"`
	// EventCreditGrantFirstDayLandingPageRatio 落地页首日授信率
	EventCreditGrantFirstDayLandingPageRatio float64 `json:"event_credit_grant_first_day_landing_page_ratio,omitempty"`
	// EventMakingCallsCost 电话拨打成本
	EventMakingCallsCost float64 `json:"event_making_calls_cost,omitempty"`
	// EventMakingCallsRatio 电话拨打率
	EventMakingCallsRatio float64 `json:"event_making_calls_ratio,omitempty"`
	// EventPayWeekByConversion 7日付费次数
	EventPayWeekByConversion int64 `json:"event_pay_week_by_conversion,omitempty"`
	// EventPayPurchaseAmountThreeDayByConversion 激活后三日付费金额
	EventPayPurchaseAmountThreeDayByConversion float64 `json:"event_pay_purchase_amount_three_day_by_conversion,omitempty"`
	// EventPayPurchaseAmountWeekByConversion 激活后7日付费金额
	EventPayPurchaseAmountWeekByConversion float64 `json:"event_pay_purchase_amount_week_by_conversion,omitempty"`
	// EventThreeDayStayByConversion 3日留存
	EventThreeDayStayByConversion int64 `json:"event_three_day_stay_by_conversion,omitempty"`
	// EventWeekStayByConversion 7日留存
	EventWeekStayByConversion int64 `json:"event_week_stay_by_conversion,omitempty"`
	// EventPayPurchaseAmountOneDayByConversion 激活后24h付费金额(激活时间)
	EventPayPurchaseAmountOneDayByConversion float64 `json:"event_pay_purchase_amount_one_day_by_conversion,omitempty"`
	// EventPayPurchaseAmountOneDay 激活后24h付费金额(回传时间)
	EventPayPurchaseAmountOneDay float64 `json:"event_pay_purchase_amount_one_day,omitempty"`
	// ConversionNum 转化数(回传时间)
	ConversionNum int64 `json:"conversion_num,omitempty"`
	// DeepConversionNum 深度转化数(回传时间)
	DeepConversionNum int64 `json:"deep_conversion_num,omitempty"`
	// ConversionNumByImpression7d 转化数(计费时间)
	ConversionNumByImpression7d int64 `json:"conversion_num_by_impression_7d,omitempty"`
	// DeepConversionNumByImpression7d 深度转化数(计费时间)
	DeepConversionNumByImpression7d int64 `json:"deep_conversion_num_by_impression_7d,omitempty"`
	// Event24Stay 活后24h次日留存数（回传时间）
	Event24Stay int64 `json:"event_24_stay,omitempty"`
	// Event24StayByConversion 激活后24h次日留存数（激活时间）
	Event24StayByConversion int64 `json:"event_24_stay_by_conversion,omitempty"`
	// EventPayWeekByConversionCost 7日付费次数成本
	EventPayWeekByConversionCost float64 `json:"event_pay_week_by_conversion_cost,omitempty"`
	// EventPayPurchaseAmountThreeDayByConversionRoi 激活后3日ROI
	EventPayPurchaseAmountThreeDayByConversionRoi float64 `json:"event_pay_purchase_amount_three_day_by_conversion_roi,omitempty"`
	// EventPayPurchaseAmountWeekByConversionRoi 激活后7日ROI
	EventPayPurchaseAmountWeekByConversionRoi float64 `json:"event_pay_purchase_amount_week_by_conversion_roi,omitempty"`
	// EventThreeDayStayByConversionCost 3日留存成本
	EventThreeDayStayByConversionCost float64 `json:"event_three_day_stay_by_conversion_cost,omitempty"`
	// EventThreeDayStayByConversionRatio 3日留存率
	EventThreeDayStayByConversionRatio float64 `json:"event_three_day_stay_by_conversion_ratio,omitempty"`
	// EventWeekStayByConversionCost 7日留存成本
	EventWeekStayByConversionCost float64 `json:"event_week_stay_by_conversion_cost,omitempty"`
	// EventWeekStayByConversionRatio 7日留存率
	EventWeekStayByConversionRatio float64 `json:"event_week_stay_by_conversion_ratio,omitempty"`
	// ConversionNumCost 转化成本（回传时间归因）
	ConversionNumCost float64 `json:"conversion_num_cost,omitempty"`
	// ConversionRatio 转化率（回传时间归因）
	ConversionRatio float64 `json:"conversion_ratio,omitempty"`
	// DeepConversionCost 深度转化成本（回传时间归因）
	DeepConversionCost float64 `json:"deep_conversion_cost,omitempty"`
	// DeepConversionRatio 深度转化率（回传时间归因）
	DeepConversionRatio float64 `json:"deep_conversion_ratio,omitempty"`
	// ConversionCostByImpression7d 转化成本（7日曝光时间归因）
	ConversionCostByImpression7d float64 `json:"conversion_cost_by_impression_7d,omitempty"`
	// ConversionRatioByImpression7d 转化率（7日曝光时间归因）
	ConversionRatioByImpression7d float64 `json:"conversion_ratio_by_impression_7d,omitempty"`
	// DeepConversionCostByImpression7d 深度转化成本（7日曝光时间归因）
	DeepConversionCostByImpression7d float64 `json:"deep_conversion_cost_by_impression_7d,omitempty"`
	// DeepConversionRatioByImpression7d 深度转化率（7日曝光时间归因）
	DeepConversionRatioByImpression7d float64 `json:"deep_conversion_ratio_by_impression7d,omitempty"`
	// EventPayPurchaseAmountOneDayByConversionRoi 激活后24小时付费ROI
	EventPayPurchaseAmountOneDayByConversionRoi float64 `json:"event_pay_purchase_amount_one_day_by_conversion_roi,omitempty"`
	// EventPayPurchaseAmountOneDayRoi 激活后24h-ROI(回传时间)
	EventPayPurchaseAmountOneDayRoi float64 `json:"event_pay_purchase_amount_one_day_roi"`
	// Event24hStayCost 激活后24h次日留存成本（回传时间）
	Event24hStayCost float64 `json:"event_24h_stay_cost,omitempty"`
	// Event24hStayRatio 激活后24h次日留存率（回传时间）
	Event24hStayRatio float64 `json:"event_24h_stay_ratio,omitempty"`
	// Event24hStayByConversionCost 激活后24h次日留存成本（激活时间）
	Event24hStayByConversionCost float64 `json:"event_24h_stay_by_conversion_cost,omitempty"`
	// Event24hStayByConversionRatio 激活后24h次日留存率（激活时间）
	Event24hStayByConversionRatio float64 `json:"event_24h_stay_by_conversion_ratio,omitempty"`
	// UnitSource 广告组来源0:常规（非托管）、1:托管
	UnitSource int `json:"unit_source,omitempty"`
}
