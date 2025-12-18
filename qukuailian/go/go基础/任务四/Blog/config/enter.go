package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Jwt      Jwt      `yaml:"jwt"`
	JwtConf  JwtConf  `yaml:"jwtConf"`
}
