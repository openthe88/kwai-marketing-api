package series

import (
	"context"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/series"
)

func PayModeTemplateInfo(ctx context.Context, clt *core.SDKClient, accessToken string, req *series.PayModeTemplateInfoRequest) ([]series.MapiSeriesPayModeTemplateInfoSnake, error) {
	var resp []series.MapiSeriesPayModeTemplateInfoSnake
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
