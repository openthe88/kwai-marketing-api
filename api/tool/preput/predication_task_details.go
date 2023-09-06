package preput

import (
	"github.com/openthe88/kwai-marketing-api/core"
	"github.com/openthe88/kwai-marketing-api/model/tool/preput"
)

// PredicationTaskDetails 投前预估详情
func PredicationTaskDetails(clt *core.SDKClient, accessToken string, req *preput.PredicationTaskCreateRequest) (*preput.AdPredicationTaskDetail, error) {
	var resp preput.AdPredicationTaskDetail
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
