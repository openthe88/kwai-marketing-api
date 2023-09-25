package appcenter

import (
	"github.com/openthe88/kwai-marketing-api/core"
	"github.com/openthe88/kwai-marketing-api/model/appcenter"
)

// AdSubPkgCreate 创建应用分包
func AdSubPkgCreate(clt *core.SDKClient, accessToken string, req *appcenter.AdSubPkgCreateRequest) (*appcenter.AdSubPkgCreateResponse, error) {
	var resp appcenter.AdSubPkgCreateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
