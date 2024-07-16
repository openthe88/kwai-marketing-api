package tool

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// IndustryCelebrityList 快手网红标签
func IndustryCelebrityList(clt *core.SDKClient, accessToken string, req *tool.IndustryCelebrityRequest) (*tool.IndustryCelebrityResponse, error) {
	var resp tool.IndustryCelebrityResponse
	err := clt.GetBody(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
