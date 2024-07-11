package creative

// ExposeTagsResponse 获取创意推荐理由 API Request
type ExposeTagsResponse struct {
	// Details 推荐理由 list
	Details []ExposeTagView `json:"details,omitempty"`
}

type ExposeTagView struct {
	Text string `json:"text,omitempty"`
}
