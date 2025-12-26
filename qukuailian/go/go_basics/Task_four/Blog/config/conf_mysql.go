package config

import "strconv"

type Mysql struct {
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"db"`
	LogLevel string `yaml:"log_level"` //日志等级，debug就是输出全部sql，dev，release
	Config   string `yaml:"config"`
}

func (m Mysql) Dsn() string {
	// 添加 @tcp() 包裹地址和端口
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Database
}
