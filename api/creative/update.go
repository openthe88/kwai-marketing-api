package creative

import (
	"github.com/openthe88/kwai-marketing-api/core"
	"github.com/openthe88/kwai-marketing-api/model/creative"
)

// Update 修改创意
// 【注】联盟广告不支持便利贴图片素材。
func Update(clt *core.SDKClient, accessToken string, req *creative.UpdateRequest) (uint64, error) {
	var resp creative.UpdateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.CreativeID, nil
}
