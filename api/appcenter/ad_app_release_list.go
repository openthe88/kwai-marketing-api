package appcenter

import (
	"github.com/openthe88/kwai-marketing-api/core"
	"github.com/openthe88/kwai-marketing-api/model/appcenter"
)

// AdAppReleaseList 获取应用列表
func AdAppReleaseList(clt *core.SDKClient, accessToken string, req *appcenter.AdAppReleaseListRequest) (*appcenter.AdAppReleaseListResponse, error) {
	var resp appcenter.AdAppReleaseListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
