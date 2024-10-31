package tool

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// OcpxOcpxTypes 获取可选的浅度优化目标
func OcpxOcpxTypes(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.OcpxOcpxTypesRequest) (*tool.OcpxOcpxTypesResponse, error) {
	var resp tool.OcpxOcpxTypesResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
