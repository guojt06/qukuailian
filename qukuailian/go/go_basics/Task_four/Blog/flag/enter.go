package flag

import (
	sys_flag "flag"
)

type Option struct {
	DB   bool
	User string
	ES   string
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	//es := sys_flag.String("es", "", "es操作")

	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		//ES:   *es,
	}
}

// 是否需要停止wed项目
func IsWebStop(option Option) (f bool) {
	if option.DB {
		return true
	}
	return true
}
func SwitchOption(option Option) bool {
	if option.DB {
		Makemigrations()
		return true
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
	}
	return false
}
