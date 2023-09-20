package appcenter

import (
	"encoding/json"
)

// AdSubPkgCreateRequest 新建应用分包 API Request
type AdSubPkgCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// ParentPackageId 应用（母）包id:仅支持android应用的包id新建分包。
	ParentPackageId int64 `json:"parent_package_id"`
	// Type 分包方式:1-系统自动分包，2-上传渠道号列表
	Type int `json:"type"`
	// Count 分包数量:当type=1时填写，单次最多100
	Count int `json:"count,omitempty"`
	// ChannelId 上传的渠道号列表:当type=2时填写，单次最多填写100个。同一应用包下填写的渠道号不可重复
	ChannelId []string `json:"channel_id,omitempty"`
	// ChannelColumns 带有分包备注的渠道号列表:与 channel_id 功能类似，但不可同时传递。
	ChannelColumns []ChannelColumn `json:"channel_columns,omitempty"`
}

// ChannelColumn 渠道号列表 API Request
type ChannelColumn struct {
	// ChannelName 渠道号(分包号)
	ChannelName string `json:"channel_name,omitempty"`
	// Description 分包备注
	Description string `json:"description,omitempty"`
}

// Url implement PostRequest interface
func (r AdSubPkgCreateRequest) Url() string {
	return "gw/dsp/appcenter/subpkg/add"
}

// Encode implement PostRequest interface
func (r AdSubPkgCreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
