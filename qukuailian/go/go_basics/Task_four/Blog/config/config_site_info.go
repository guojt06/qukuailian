package config

type SiteInfo struct {
	CreatedAt   string `yaml:"created_at" json:"created_at"`
	BeiAn       string `yaml:"bei_an" json:"bei_an"`
	Title       string `yaml:"title" json:"title"`
	QQImage     string `yaml:"qq_image" json:"qq_image"`
	Version     string `yaml:"version" json:"version"`
	Email       string `yaml:"email" json:"email"`
	WechatImage string `yaml:"wechat_image" json:"wechat_image"`
	Name        string `yaml:"name" json:"name"`
	Jog         string `yaml:"jog" json:"jog"`
	Addr        string `yaml:"addr" json:"addr"`
	Slogan      string `yaml:"slogan" json:"slogan"`
	Web         string `yaml:"web" json:"web"`
}
