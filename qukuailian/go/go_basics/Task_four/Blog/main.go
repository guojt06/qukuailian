package main

import (
	"modulename/core"
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
	//初始化数据库
	//optin := flag.Parse()
	//if flag.IsWebStop(optin) {
	//	flag.SwitchOption(optin)
	//	return
	//}

	//路由
	routers := routers.InitRouters()
	//global.Log.Info("server运行在: %s", global.Config.System.Addr())
	routers.Run(global.Config.System.Addr())
}
