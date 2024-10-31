package tool

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

func KeywordQueryList(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.KeywordQueryRequest) (*tool.KeywordQueryResponse, error) {
	var resp tool.KeywordQueryResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
