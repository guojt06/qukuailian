package models

type TagModel struct {
	MODEL
	Title string `gorm:"size:16" json:"titlle"`
}
