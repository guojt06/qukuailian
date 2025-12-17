package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id_list" form:"id_list" binding:"required"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key" json:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
