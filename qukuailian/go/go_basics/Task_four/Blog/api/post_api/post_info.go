package post_api

import (
	"fmt"
	"modulename/global"
	"modulename/models"
	"modulename/models/res"
	"modulename/plugins/log_stash"

	"github.com/gin-gonic/gin"
)

//文章管理功能
//实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
//实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
//实现文章的更新功能，只有文章的作者才能更新自己的文章。
//实现文章的删除功能，只有文章的作者才能删除自己的文章。
//评论功能
//实现评论的创建功能，已认证的用户可以对文章发表评论。
//实现评论的读取功能，支持获取某篇文章的所有评论列表。

type CreateArticleRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (PostApi) Insert(ctx *gin.Context) {
	//实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
	var cr CreateArticleRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	log := log_stash.NewLogByGin(ctx)

	// 从上下文中获取用户ID
	//userID, _ := ctx.Get("userID")

	//获取文章的模型
	var articleModel models.ArticleModel
	articleModel.Title = cr.Title
	articleModel.Content = cr.Content
	//articleModel.UserID = string(userID)
	err = global.DB.Save(&articleModel).Error
	if err != nil {
		log.Warn(fmt.Sprintf("%s 文章", cr.Title))
		return
	}
}
