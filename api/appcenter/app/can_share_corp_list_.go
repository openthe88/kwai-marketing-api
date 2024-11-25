package app

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// CanShareCorpList 获取可共享的主体列表(接口只有在跨主体共享白名单内才会有数据返回，如没有相关场景，则可忽略此接口)
func CanShareCorpList(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.CanShareCorpListRequest) (*app.CanShareCorpListResponse, error) {
	var resp app.CanShareCorpListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
