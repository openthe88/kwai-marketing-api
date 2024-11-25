package app

// ShareCorp 账号列表
type ShareCorp struct {
	// CorpID 主体ID
	CorpID uint64 `json:"corp_id,omitempty"`
	// CorpName 主体名称
	CorpName string `json:"corp_name,omitempty"`
	// TotalAccountCnt 主体共享，挂载的Account数量
	TotalAccountCnt int `json:"total_account_cnt,omitempty"`
	// ShareAccountVos 账号列表
	ShareAccountVos []ShareAccount `json:"share_account_vos,omitempty"`
}

// ShareAccount 账号列表
type ShareAccount struct {
	// AccountID 账号ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountName 账号名称
	AccountName string `json:"account_name,omitempty"`
}
