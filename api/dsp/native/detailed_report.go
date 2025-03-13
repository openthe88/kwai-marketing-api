package native

import (
	"context"
	"github.com/bububa/kwai-marketing-api/model/dsp/native"

	"github.com/bububa/kwai-marketing-api/core"
)

func DetailedReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *native.DetailedReportRequest) (*native.DetailedReportResponse, error) {
	var resp native.DetailedReportResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
