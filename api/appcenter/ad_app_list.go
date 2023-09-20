package appcenter

import (
	"github.com/openthe88/kwai-marketing-api/core"
	"github.com/openthe88/kwai-marketing-api/model/appcenter"
)

// AdAppList 获取应用列表
func AdAppList(clt *core.SDKClient, accessToken string, req *appcenter.AdAppListRequest) (*appcenter.AdAppListResponse, error) {
	var resp appcenter.AdAppListResponse
	err := clt.GetOnBody(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
