package config

type System struct {
	HOST string `yaml:"host"`
	PORT uint16 `yaml:"port"`
	ENV  string `yaml:"env"`
}
