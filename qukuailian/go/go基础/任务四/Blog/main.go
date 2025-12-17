package main

import (
	"modulename/core"
	"modulename/flag"
	"modulename/global"
	"modulename/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()
	//连接数据库
	//fmt.Println(global.DB)
	//路由之前执行
	optin := flag.Parse()
	if flag.IsWebStop(optin) {
		flag.SwitchOption(optin)
		return
	}

	routers := routers.InitRouters()
	//global.Log.Info("server运行在: %s", global.Config.System.Addr())
	routers.Run(global.Config.System.Addr())
}
