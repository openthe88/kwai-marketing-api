package tool

// OcpxDeepTypesResponse 获取可选的深度优化目标
type OcpxDeepTypesResponse struct {
	//DeepTypes 可用目标数组
	DeepTypes []DeepTypes `json:"deep_types"`
}

type DeepTypes struct {
	// EventType 深度优化目标
	EventType int `json:"event_type"`

	// Name event_type的中文描述
	Name string `json:"name"`
}
