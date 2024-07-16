package tool

// BehaviorInterestResponse 获取行为与兴趣"类目"
type BehaviorInterestResponse struct {
	//BehaviorInterest 行为兴趣类目,详情见下方
	BehaviorInterest BehaviorInterest `json:"behavior_interest"`
}

type BehaviorInterest struct {
	//Behavior 行为类目，详情见下方
	Behavior []Behavior `json:"behavior"`
	//Interest 兴趣类目，详情见下方
	Interest []Interest `json:"interest"`
}

type Interest struct {
	// ID 兴趣类目的ID
	ID uint64 `json:"id,omitempty"`

	// Name 兴趣类目的名称
	Name string `json:"name,omitempty"`

	// Level 类目所在的层数
	Level int `json:"level,omitempty"`

	// Des 类目的介绍
	Des string `json:"des,omitempty"`

	// Children 当前类目的子类目，父类目与子类目结构相同
	Children []Interest `json:"children,omitempty"`

	// Count 当前类目的人群预估数量
	Count uint64 `json:"count,omitempty"`
}

type Behavior struct {
	// ID 行为类目的ID
	ID uint64 `json:"id,omitempty"`

	// Name 行为类目的名称
	Name string `json:"name,omitempty"`

	// Level 类目所在的层数
	Level int `json:"level,omitempty"`

	// Des 类目的介绍
	Des string `json:"des,omitempty"`

	// Children 当前类目的子类目，父类目与子类目结构相同
	Children []Behavior `json:"children,omitempty"`

	// Count 当前类目的人群预估数量
	Count uint64 `json:"count,omitempty"`
}
