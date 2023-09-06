package oauth

import (
	"strings"

	"github.com/openthe88/kwai-marketing-api/core"
	"github.com/openthe88/kwai-marketing-api/model/oauth"
)

// Url 生成授权链接
func Url(clt *core.SDKClient, req *oauth.UrlRequest) string {
	req.AppID = clt.AppID()
	var builder strings.Builder
	builder.WriteString(core.OAUTH_URL)
	builder.WriteString("/")
	builder.WriteString(req.Url())
	builder.WriteString("?")
	builder.WriteString(req.Encode())
	return builder.String()
}
