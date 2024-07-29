package tool

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// CelebrityList 搜索快手网红
func CelebrityList(clt *core.SDKClient, accessToken string, req *tool.CelebrityRequest) (*tool.CelebrityResponse, error) {
	var resp tool.CelebrityResponse
	err := clt.GetBody(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
