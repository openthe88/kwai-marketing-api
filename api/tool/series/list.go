package series

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/series"
)

func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *series.ListRequest) (*series.ListResponse, error) {
	var resp series.ListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
