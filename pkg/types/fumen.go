package types

const (
	// CategoryPublished [...]
	CategoryPublished int = 0
	// CategoryPublishing [...]
	CategoryPublishing int = 1
	// CategoryOfficial [...]
	CategoryOfficial int = 2
	// CategoryNotPublished [...]
	CategoryNotPublished int = 3
	// CategoryGift [...]
	CategoryGift = 4
	// CategorySpecial [...]
	CategorySpecial = 5
)

// GetFumenRequest [...]
type GetFumenRequest struct {
	*Header
	Params struct {
		Keyword  string `json:"key_word"`
		Creator  string `json:creator`
		Artist   string `json:artist`
		Category int    `json:category`
		Page     int    `json:"page"`
		Start    int    `json:"start"`
	} `json:"params"`
}

// GetFumenResult [...]
type GetFumenResult struct {
	Data  []*Songs `json:"data"`
	Total int      `json:"total"`
}

// GetFumenResponse [...]
type GetFumenResponse struct {
	*Header
	Result *GetFumenResult `json:"result"`
}
