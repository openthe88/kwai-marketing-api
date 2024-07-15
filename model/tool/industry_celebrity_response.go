package tool

// IndustryCelebrityResponse 快手网红-网红分类
type IndustryCelebrityResponse struct {
	//快手网红标签
	CelebrityTree []IndustryCelebrityNode `json:"celebrity_tree,omitempty"`
}

type IndustryCelebrityNode struct {
	//Id 标签ID
	Id int64 `json:"id,omitempty"`
	//Name 标签名字
	Name string `json:"name,omitempty"`
	//Children 树级标签
	Children []IndustryCelebrityNode `json:"children,omitempty"`
}
