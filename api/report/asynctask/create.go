package asynctask

import (
	"github.com/openthe88/kwai-marketing-api/core"
	"github.com/openthe88/kwai-marketing-api/model/report/asynctask"
)

// Create 创建历史数据查询任务
func Create(clt *core.SDKClient, accessToken string, req *asynctask.CreateRequest) (*asynctask.CreateResponse, error) {
	var resp asynctask.CreateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
