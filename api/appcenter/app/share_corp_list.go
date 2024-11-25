package app

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// ShareCorpList 获取应用已共享主体列表（应用的共享类型为：share_type=2【主体共享】才会有数据）
func ShareCorpList(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.ShareCorpListRequest) (*app.ShareCorpListResponse, error) {
	var resp app.ShareCorpListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
