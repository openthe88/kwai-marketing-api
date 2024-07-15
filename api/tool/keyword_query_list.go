package tool

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

func KeywordQueryList(clt *core.SDKClient, accessToken string, req *tool.KeywordQueryRequest) (*tool.KeywordQueryResponse, error) {
	var resp tool.KeywordQueryResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
