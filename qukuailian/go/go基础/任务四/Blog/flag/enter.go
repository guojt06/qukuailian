package flag

import (
	sys_flag "flag"

	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string
	ES   string
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	//user := sys_flag.String("u", "", "创建用户")
	//es := sys_flag.String("es", "", "es操作")

	sys_flag.Parse()
	return Option{
		DB: *db,
		//User: *user,
		//ES:   *es,
	}
}

func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case bool:
			if val {
				f = true
			}
		case string:
			if val != "" {
				f = true
			}
		}
	}
	return f
}
func SwitchOption(option Option) bool {
	if option.DB {
		Makemigrations()
		return true
	}
	return false
}
