package app

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// ShareListCorpAccount 获取单个主体下共享账号列表(应用的共享类型为：share_type=2【主体共享】才会有数据)
func ShareListCorpAccount(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.ShareListCorpAccountRequest) (*app.ShareListCorpAccountResponse, error) {
	var resp app.ShareListCorpAccountResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
