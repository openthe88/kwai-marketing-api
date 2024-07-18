package tool

// OcpxOcpxTypesResponse 获取可选的浅度优化目标
type OcpxOcpxTypesResponse struct {
	//可用目标数组
	OcpxTypes []OcpxTypes `json:"ocpx_types"`
}

type OcpxTypes struct {
	//OcpxActionType 优化目标
	OcpxActionType int `json:"ocpx_action_type"`
	//Name ocpx_action_type的中文描述 在不同的campaignType下，本描述对部分优化目标会有不同的输出
	Name string `json:"name"`
}
