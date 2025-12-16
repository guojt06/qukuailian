package config

import "fmt"

type System struct {
	HOST string `yaml:"host"`
	PORT uint16 `yaml:"port"`
	ENV  string `yaml:"env"`
}

func (s System) Addr() string {
	return fmt.Sprint(s.HOST, ":", s.PORT)
}
