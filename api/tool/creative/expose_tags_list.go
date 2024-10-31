package creative

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/creative"
)

// ExposeTags 获取创意推荐理由
func ExposeTags(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.ExposeTagsRequest) (*creative.ExposeTagsResponse, error) {
	var resp creative.ExposeTagsResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
