package tool

// AppSearchResponse 获取可选的应用定向 API Request
type AppSearchResponse struct {
	TargetingApp []TargetingApp `json:"targeting_app"`
}
