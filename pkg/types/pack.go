package types

// GetPackRequest [...]
type GetPackRequest struct {
	*Header
	Params struct {
		Keyword  string `json:"key_word"`
		Category int    `json:"category"`
		Page     int    `json:"page"`
		Start    int    `json:"start"`
	} `json:"params"`
}

// GetPackResult [...]
type GetPackResult struct {
	Data  []*Packs `json:"data"`
	Total int      `json:"total"`
}

// GetPackResponse [...]
type GetPackResponse struct {
	*Header
	Result *GetPackResult `json:"result"`
}
