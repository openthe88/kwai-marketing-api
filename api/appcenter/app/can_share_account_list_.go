package app

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// CanShareAccountList 获取可共享的账号列表
func CanShareAccountList(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.CanShareAccountListRequest) (*app.CanShareAccountListResponse, error) {
	var resp app.CanShareAccountListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
