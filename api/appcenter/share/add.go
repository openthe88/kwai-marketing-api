package share

import (
	"context"
	"github.com/bububa/kwai-marketing-api/model/appcenter/share"

	"github.com/bububa/kwai-marketing-api/core"
)

// Add 【应用共享】添加应用共享
func Add(ctx context.Context, clt *core.SDKClient, accessToken string, req *share.AddRequest) (share.AddResponse, error) {
	var resp share.AddResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}
