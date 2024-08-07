package tool

// TargetBehaviorResponse 行为意向4.0，类目查询接口
type TargetBehaviorResponse struct {
}

type TargetBehaviorCategory struct {
	// Children 子类目列表
	Children []TargetBehaviorCategory `json:"children"`

	// Count 类目覆盖人群数
	Count uint64 `json:"count"`

	// ID 类目ID，可用于广告组创编
	ID uint64 `json:"id"`

	// Level 类目层级
	Level int `json:"level"`

	// Name 类目名称
	Name string `json:"name"`
}
