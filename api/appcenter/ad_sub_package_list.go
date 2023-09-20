package appcenter

import (
	"github.com/openthe88/kwai-marketing-api/core"
	"github.com/openthe88/kwai-marketing-api/model/appcenter"
)

// AdSubPackageList 获取分包列表
func AdSubPackageList(clt *core.SDKClient, accessToken string, req *appcenter.AdSubPackageListRequest) (*appcenter.AdSubPackageListResponse, error) {
	var resp appcenter.AdSubPackageListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
