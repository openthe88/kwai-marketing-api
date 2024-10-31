package creative

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/creative"
)

// ElementReviewDetails 获取创意审核详情
func ElementReviewDetails(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.ElementReviewDetailsRequest) (*[]creative.ElementReviewDetailsResponse, error) {
	var resp []creative.ElementReviewDetailsResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
