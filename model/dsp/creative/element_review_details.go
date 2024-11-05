package creative

import "encoding/json"

// ElementReviewDetailsRequest 获取创意审核详情 API Request
type ElementReviewDetailsRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// IDs 自定义创意时为创意ID 程序化创意时为组ID
	IDs []int64 `json:"ids,omitempty"`
	// CreativeMold 创意类型; 1 自定义创意 2程序化创意
	CreativeMold int `json:"creative_mold,omitempty"`
}

// Url implement PostRequest interface
func (r ElementReviewDetailsRequest) Url() string {
	return "gw/dsp/creative/element/reviewDetails"
}

// Encode implement PostRequest interface
func (r ElementReviewDetailsRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// ElementReviewDetailsResponse 获取创意审核详情 API Response
type ElementReviewDetailsResponse struct {
	//社区审核状态 1审核中 2审核通过 3审核拒绝 5基本通过审核(指视频通过审核)
	CommunityReviewStatus int `json:"community_review_status,omitempty"`
	//商业审核状态 1审核中 2审核通过 3审核拒绝 5基本通过审核(指视频通过审核)
	ReviewStatus int `json:"review_status,omitempty"`
	//商业审核拒绝详情
	ReviewReason []ReviewReason `json:"review_reason,omitempty"`
	//限流详情
	LimitingReason []LimitingReason `json:"limiting_reason,omitempty"`
	//程序化创意社区审核拒绝详情
	AdvCreativeCommunityReviewDetail []AdvCreativeCommunityReviewDetail `json:"adv_creative_community_review_detail,omitempty"`
	//自定义创意社区审核拒绝详情
	CustomCreativeCommunityReviewDetail *CustomCreativeCommunityReviewDetail `json:"custom_creative_community_review_detail,omitempty"`
	//程序化创意组id，自定义创意创意id
	ID int64 `json:"id,omitempty"`
}

// ReviewReason 商业审核拒绝详情
type ReviewReason struct {
	//视频ID或者创意ID
	ID string `json:"id,omitempty"`
	//审核元素 1视频 2封面 3广告语 4图片 6创意 7图集
	Type int `json:"type,omitempty"`
	//限流类型 1低质 2降级 3负向 4封面没过
	NegativeType int `json:"negative_type,omitempty"`
	//审核拒绝/限流原因和建议
	ReasonAndModify []ReasonAndModify `json:"reason_and_modify,omitempty"`
}

// ReasonAndModify 审核拒绝/限流原因和建议
type ReasonAndModify struct {
	//审核拒绝原因
	Reason string `json:"reason,omitempty"`
	//修改建议
	Modify []string `json:"modify,omitempty"`
}

// LimitingReason 限流详情
type LimitingReason struct {
	//视频ID或者创意ID
	ID string `json:"id,omitempty"`
	//审核元素 1视频 2封面 3广告语 4图片 6创意 7图集
	Type int `json:"type,omitempty"`
	//限流类型 1低质 2降级 3负向 4封面没过
	NegativeType int `json:"negative_type,omitempty"`
	//审核拒绝/限流原因和建议
	ReasonAndModify []ReasonAndModify `json:"reason_and_modify,omitempty"`
}

// AdvCreativeCommunityReviewDetail 程序化创意社区审核拒绝详情
type AdvCreativeCommunityReviewDetail struct {
	// 创意ID
	ID int64 `json:"id,omitempty"`
	// 视频ID
	PhotoID string `json:"photo_id,omitempty"`
	// 封面ID
	CoverID int64 `json:"cover_id,omitempty"`
	// 封面URL
	CoverURL string `json:"cover_url,omitempty"`
	// 视频标题
	Caption string `json:"caption,omitempty"`
	// 审核拒绝理由
	CommunityReviewDetail string `json:"community_review_detail,omitempty"`
}

// CustomCreativeCommunityReviewDetail 自定义创意社区审核拒绝详情
type CustomCreativeCommunityReviewDetail struct {
	// 创意ID
	ID int64 `json:"id,omitempty"`
	// 视频ID
	PhotoID string `json:"photo_id,omitempty"`
	// 封面ID
	CoverID int64 `json:"cover_id,omitempty"`
	// 封面URL
	CoverURL string `json:"cover_url,omitempty"`
	// 视频标题
	Caption string `json:"caption,omitempty"`
	// 审核拒绝理由
	CommunityReviewDetail string `json:"community_review_detail,omitempty"`
}
