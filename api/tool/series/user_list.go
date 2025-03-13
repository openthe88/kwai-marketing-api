package series

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/series"
)

func UserList(ctx context.Context, clt *core.SDKClient, accessToken string, req *series.UserListRequest) ([]series.MapiSeriesUserSnake, error) {
	var resp []series.MapiSeriesUserSnake
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
