package creative

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/creative"
)

// CategoryList 获取创意分类查询接口
func CategoryList(clt *core.SDKClient, accessToken string, req *creative.CategoryRequest) (*creative.CategoryResponse, error) {
	var resp creative.CategoryResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}