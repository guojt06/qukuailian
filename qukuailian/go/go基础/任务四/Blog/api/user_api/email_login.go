package user_api

import (
	"fmt"
	"modulename/global"
	"modulename/models"
	"modulename/models/res"
	"modulename/plugins/log_stash"
	"modulename/utils/jwts"

	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (UserApi) EmailLoginView(ctx *gin.Context) {
	var cr EmailLoginRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	log := log_stash.NewLogByGin(ctx)

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在")
		res.FailWithMessage("用户名或密码错误", ctx)
		log.Warn(fmt.Sprintf("%s 用户名不存在", cr.UserName))
		return
	}
	// TODO 校验密码
	if cr.Password != userModel.Password {
		global.Log.Warn("密码错误")
		res.FailWithMessage("用户名或密码错误", ctx)
		log.Warn(fmt.Sprintf("用户名 %s 或密码 %s 错误", cr.UserName, cr.Password))
		return
	}

	// 返回token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserId:   userModel.ID,
		Avatar:   userModel.Avatar,
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("生成 token 失败", ctx)
		log.Error(fmt.Sprintf("生成 token 失败 %s", err.Error()))

		return
	}
	log = log_stash.New(ctx.ClientIP(), token)
	log.Info("登录成功")
	res.OKWithData(token, ctx)
}
