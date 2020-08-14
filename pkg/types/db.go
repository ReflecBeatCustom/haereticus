package types

import "time"

// Packs [...]
type Packs struct {
	Category      int       `gorm:"primary_key;column:Category;type:int(11) unsigned;not null" json:"-"`
	PackID        int       `gorm:"primary_key;column:PackID;type:int(11);not null" json:"-"`
	Title         string    `gorm:"column:Title;type:text" json:"title"`
	Comment       string    `gorm:"column:Comment;type:text" json:"comment"`
	HasPromotion  int       `gorm:"column:HasPromotion;type:int(11)" json:"has_promotion"`
	PreviewSongID int       `gorm:"column:PreviewSongID;type:int(11) unsigned" json:"preview_song_id"`
	CreateTime    time.Time `gorm:"column:CreateTime;type:timestamp" json:"create_time"`
}

// Playrecords [...]
type Playrecords struct {
	AccountName string    `gorm:"column:AccountName;type:varchar(50)" json:"account_name"`
	SongID      int       `gorm:"column:SongID;type:int(11)" json:"song_id"`
	Difficulty  int       `gorm:"column:Difficulty;type:int(11)" json:"difficulty"`
	LogTime     time.Time `gorm:"column:LogTime;type:timestamp" json:"log_time"`
	Score       int       `gorm:"column:Score;type:int(11)" json:"score"`
	JR          int       `gorm:"column:JR;type:int(11)" json:"jr"`
	Note        int       `gorm:"column:Note;type:int(11)" json:"note"`
}

// Songs [...]
type Songs struct {
	SongID            int       `gorm:"primary_key;column:SongID;type:int(11);not null" json:"-"`
	Category          int       `gorm:"column:Category;type:int(11)" json:"category"`
	PackID            int       `gorm:"column:PackID;type:int(11)" json:"pack_id"`
	AccessLevel       int       `gorm:"column:AccessLevel;type:int(11)" json:"access_level"`
	Creator           string    `gorm:"column:Creator;type:text" json:"creator"`
	Title             string    `gorm:"column:Title;type:text" json:"title"`
	Artist            string    `gorm:"column:Artist;type:text" json:"artist"`
	ChartAuthor       string    `gorm:"column:ChartAuthor;type:text" json:"chart_author"`
	DiffB             int       `gorm:"column:diffB;type:int(11)" json:"diff_b"`
	DiffM             int       `gorm:"column:diffM;type:int(11)" json:"diff_m"`
	DiffH             int       `gorm:"column:diffH;type:int(11)" json:"diff_h"`
	DiffSP            int       `gorm:"column:diffSP;type:int(11)" json:"diff_sp"`
	SubdiffB          int       `gorm:"column:subdiffB;type:int(11)" json:"subdiff_b"`
	SubdiffM          int       `gorm:"column:subdiffM;type:int(11)" json:"subdiff_m"`
	SubdiffH          int       `gorm:"column:subdiffH;type:int(11)" json:"subdiff_h"`
	SubdiffSP         int       `gorm:"column:subdiffSP;type:int(11)" json:"subdiff_sp"`
	HasSpecial        int       `gorm:"column:HasSpecial;type:int(11)" json:"has_special"`
	SpecialID         int       `gorm:"column:SpecialID;type:int(11)" json:"special_id"`
	SequencePID       int       `gorm:"column:SequencePID;type:int(11)" json:"sequence_p_id"`
	CreateTime        time.Time `gorm:"column:CreateTime;type:timestamp" json:"create_time"`
	UpdateTime        time.Time `gorm:"column:UpdateTime;type:timestamp" json:"update_time"`
	Description       string    `gorm:"column:Description;type:text" json:"description"`
	AuthorDescription string    `gorm:"column:AuthorDescription;type:text" json:"author_description"`
}

// Fumen [...]
type Fumen struct {
	FumenID     int    `gorm:"unique;index:fumen_id_index;column:fumen_id;type:int(10) unsigned;not null" json:"fumen_id"`
	ArtistName  string `gorm:"index:artist_name_index;column:artist_name;type:varchar(50);not null" json:"artist_name"`
	MusicName   string `gorm:"index:music_name_index;column:music_name;type:varchar(50);not null" json:"music_name"`
	Version     int    `gorm:"column:version;type:int(10) unsigned;not null" json:"version"`
	BpmMax      int    `gorm:"column:bpm_max;type:int(10) unsigned;not null" json:"bpm_max"`
	BpmMin      int    `gorm:"column:bpm_min;type:int(10) unsigned;not null" json:"bpm_min"`
	BasicLevel  int    `gorm:"column:basic_level;type:int(10) unsigned;not null" json:"basic_level"`
	MediumLevel int    `gorm:"column:medium_level;type:int(10) unsigned;not null" json:"medium_level"`
	HardLevel   int    `gorm:"column:hard_level;type:int(10) unsigned;not null" json:"hard_level"`
}
