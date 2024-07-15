package asset_media

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/asset_media"
)

func MediaList(clt *core.SDKClient, accessToken string, req *asset_media.MediaListRequest) (*asset_media.MediaListResponse, error) {
	var resp asset_media.MediaListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
