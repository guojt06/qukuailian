package flag

import (
	"fmt"
	"modulename/global"
	"modulename/models"
	"modulename/models/ctype"
)

func CreateUser(permission string) {
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称：")
	fmt.Scan(&nickName)
	fmt.Printf("请输入密码：")
	fmt.Scan(&password)
	fmt.Printf("请再次输入密码：")
	fmt.Scan(&rePassword)

	// 判断用户名是否存在
	var userModel models.UserModel
	count := global.DB.Take(&userModel, "user_name = ?", userName).RowsAffected

	if count != 0 {
		// 存在用户
		global.Log.Errorf("用户 %s 已存在", userName)
		return
	}
	role := ctype.PermissionUser
	if permission == "admin" {
		role = ctype.PermissionAdmin
	}

	avatar := "http://192.168.31.194/uploads/avatar/image.png"

	// 入库
	err := global.DB.Create(&models.UserModel{
		UserName:   userName,
		NickName:   nickName,
		Password:   password,
		Role:       role,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error

	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户 %s 创建成功", userName)
}
