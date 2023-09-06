package preput

import (
	"github.com/openthe88/kwai-marketing-api/core"
	"github.com/openthe88/kwai-marketing-api/model/tool/preput"
)

// PredicationTaskManagement 投前预估任务管理接口
func PredicationTaskManagement(clt *core.SDKClient, accessToken string, req *preput.PredicationTaskManagementRequest) (*preput.RealTaskResult, error) {
	var resp preput.RealTaskResult
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
