package dmp

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dmp"
)

// PopulationSystemList 系统推荐定向/排除人群包
func PopulationSystemList(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationSystemRequest) (*dmp.PopulationSystemResponse, error) {
	var resp *dmp.PopulationSystemResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
