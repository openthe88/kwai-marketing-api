package oauth

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/oauth"
)

// ApprovalList 拉取token下授权广告账户接口
func ApprovalList(clt *core.SDKClient, accessToken string, pageNo int, pageSize int) ([]byte, error) {
	req := &oauth.ApprovalListRequest{
		AppID:       clt.AppID(),
		Secret:      clt.Secret(),
		AccessToken: accessToken,
		PageNo:      pageNo,
		PageSize:    pageSize,
	}
	res, err := clt.OauthPost("", req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
