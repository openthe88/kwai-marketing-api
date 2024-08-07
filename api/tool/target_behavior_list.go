package tool

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// TargetBehaviorList 行为意向4.0，类目查询接口
func TargetBehaviorList(clt *core.SDKClient, accessToken string, req *tool.TargetBehaviorRequest) ([]tool.TargetBehaviorCategory, error) {
	var resp []tool.TargetBehaviorCategory
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
