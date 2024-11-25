package app

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// ShareList 获取应用已共享账号列表（应用的共享类型为：share_type=1【账号共享】才会有数据）
func ShareList(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.ShareListRequest) (*app.ShareListResponse, error) {
	var resp app.ShareListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
