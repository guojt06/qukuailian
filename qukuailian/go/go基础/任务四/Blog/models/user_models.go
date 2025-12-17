package models

import "modulename/models/ctype"

type UserModel struct { // 用户表
	MODEL
	NickName   string           `gorm:"size:36" json:"nick_name,select(info|comment)"`
	UserName   string           `gorm:"size:36" json:"user_name"`
	Password   string           `gorm:"size:128" json:"-"`
	Avatar     string           `gorm:"size:256" json:"avatar,select(info|comment)"`
	Email      string           `gorm:"size:128" json:"email,select(info)"`
	Tel        string           `gorm:"size:18" json:"tel"`
	Addr       string           `gorm:"size:64" json:"addr,select(info|comment)"`
	Token      string           `gorm:"size:64" json:"token"`
	IP         string           `gorm:"size:20" json:"ip,select(comment)"`
	Role       ctype.Role       `gorm:"size:4;default:1" json:"role,select(info)"`
	SignStatus ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status,select(info)"`
	Signature  string           `gorm:"size" json:"signature,select(info)"`
	Link       string           `gorm:"size" json:"link,select(info)"`
}
