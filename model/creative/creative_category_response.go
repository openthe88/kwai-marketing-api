package creative

// CategoryResponse 获取创意分类查询接口 API Request
type CategoryResponse struct {
	Details []CategoryView `json:"details"`
}

type CategoryView struct {
	//分类 id
	CategoryId int `json:"category_id"`
	//父节点分类 id
	ParentId int `json:"parent_id"`
	//分类名称
	CategoryName string `json:"category_name"`
	//层级
	Level string `json:"level"`
}
