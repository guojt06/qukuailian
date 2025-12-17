package models

import (
	"modulename/global"
	"modulename/models/ctype"
	"os"

	"gorm.io/gorm"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`
	Hash      string          `json:"hash"`
	Name      string          `gorm:"size:38" json:"name"`
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片存储类型
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		err := os.Remove(b.Path[1:])
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return
}
