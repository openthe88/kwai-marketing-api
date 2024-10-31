package tool

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// OcpxDeepTypes 获取可选的深度优化目标
func OcpxDeepTypes(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.OcpxDeepTypesRequest) (*tool.OcpxDeepTypesResponse, error) {
	var resp tool.OcpxDeepTypesResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
