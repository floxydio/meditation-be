package models

type Music struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Url       string `json:"url"`
	Category  string `json:"category"`
	ColorCard string `json:"color_card"`
	ImageCard string `json:"image_card"`
}

func (Music) TableName() string {
	return "music"
}
