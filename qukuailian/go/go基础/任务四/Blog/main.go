package main

import (
	"fmt"
	"modulename/core"
	"modulename/global"
	"modulename/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Error("哈哈")
	global.Log.Warnf("呵呵")
	global.Log.Info("嘻嘻")
	//连接数据库
	global.DB = core.InitGorm()
	//连接数据库
	fmt.Println(global.DB)

	routers := routers.InitRouters()
	routers.Run()
}
