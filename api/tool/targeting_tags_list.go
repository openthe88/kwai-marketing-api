package tool

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// TargetingTagsList 获取可选的定向标签
func TargetingTagsList(clt *core.SDKClient, accessToken string, req *tool.TargetingTagsListRequest) (*tool.TargetingTagsListResponse, error) {
	var resp tool.TargetingTagsListResponse
	err := clt.GetBody(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
