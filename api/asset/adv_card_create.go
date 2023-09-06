package asset

import (
	"github.com/openthe88/kwai-marketing-api/core"
	"github.com/openthe88/kwai-marketing-api/model/asset"
)

// AdvCardCreate 创建高级创意接口
func AdvCardCreate(clt *core.SDKClient, accessToken string, req *asset.AdvCardCreateRequest) ([]uint64, error) {
	var resp asset.AdvCardCreateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.AdvList, nil
}
