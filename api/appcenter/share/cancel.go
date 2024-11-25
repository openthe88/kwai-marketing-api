package share

import (
	"context"
	"github.com/bububa/kwai-marketing-api/model/appcenter/share"

	"github.com/bububa/kwai-marketing-api/core"
)

// Cancel 【应用共享】取消应用共享
func Cancel(ctx context.Context, clt *core.SDKClient, accessToken string, req *share.CancelRequest) (share.CancelResponse, error) {
	var resp share.CancelResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}
