package series

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/series"
)

func EpisodeList(ctx context.Context, clt *core.SDKClient, accessToken string, req *series.EpisodeListRequest) (*series.EpisodeListResponse, error) {
	var resp series.EpisodeListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
