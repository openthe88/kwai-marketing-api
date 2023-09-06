package dmp

import (
	"github.com/openthe88/kwai-marketing-api/core"
	"github.com/openthe88/kwai-marketing-api/model/dmp"
)

// PopulationUpdateV2 人群包更新接口
func PopulationUpdateV2(clt *core.SDKClient, accessToken string, req *dmp.PopulationUpdateRequestv2) (*dmp.PopulationV2, error) {
	var resp dmp.PopulationV2
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
