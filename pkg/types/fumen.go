package types

type Fumen struct {
	FumenID     int `json:"fumen_id"`
	ArtistName  int `json:"artist_name"`
	MusicName   int `json:"music_name"`
	BPMMax      int `json:"bpm_max"`
	BPMMin      int `json:"bpm_min"`
	BasicLevel  int `json:"basic_level"`
	MediumLevel int `json:"medium_level"`
	HardLevel   int `json:"hard_level"`
}

// GetFumenRequest ...
type GetFumenRequest struct {
	*Header
	Params struct {
		Keyword string `json:"key_word"`
		Page    int    `json:"page"`
		Start   int    `json:"start"`
	} `json:"params"`
}

// GetFumenResponse ...
type GetFumenResponse struct {
	*Header
	Fumens []*Fumen `json:"fumens"`
}
