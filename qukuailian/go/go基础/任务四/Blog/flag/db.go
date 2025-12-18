package flag

import (
	"modulename/global"
	"modulename/models"
)

func Makemigrations() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_golang", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			//&models.AdverModel{},
			&models.UserModel{},
			&models.CommentModel{},
			&models.ArticleModel{},
			&models.UserCollectModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			//&models.FeedBackModel{},
			&models.LoginDataModel{},
			//&log_stash.LogStashModel{},
		)
	if err != nil {
		global.Log.Error("[ error ]: 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ]: 生成数据库表结构成功")
}
