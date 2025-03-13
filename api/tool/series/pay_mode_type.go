package series

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/series"
)

func PayModeType(ctx context.Context, clt *core.SDKClient, accessToken string, req *series.PayModeTypeRequest) ([]series.MapiSeriesPayModeInfoSnake, error) {
	var resp []series.MapiSeriesPayModeInfoSnake
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
