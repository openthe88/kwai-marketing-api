package tool

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// BehaviorInterestList 获取行为与兴趣"类目"
func BehaviorInterestList(clt *core.SDKClient, accessToken string, req *tool.BehaviorInterestRequest) (*tool.BehaviorInterestResponse, error) {
	var resp tool.BehaviorInterestResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
