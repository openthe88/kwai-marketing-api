package tool

// CelebrityResponse 搜索快手网红
type CelebrityResponse struct {
	CelebrityList []CelebrityList `json:"celebrity_list"`
}

type CelebrityList struct {
	// FirstLabelID firstLabelId
	FirstLabelID int `json:"first_label_id"`

	// FirstLabel firstLabel
	FirstLabel string `json:"first_label"`

	// SecondLabelID secondLabelId
	SecondLabelID int `json:"second_label_id"`

	// SecondLabel secondLabel
	SecondLabel string `json:"second_label"`

	// AuthorID authorId
	AuthorID string `json:"author_id"`

	// KwaiID kwaiId
	KwaiID string `json:"kwai_id"`

	// AuthorName authorName
	AuthorName string `json:"author_name"`

	// FansNum fansNum
	FansNum uint64 `json:"fans_num"`
}
