package core

import (
	"fmt"
	ioutil "io/ioutil"
	"log"
	"modulename/config"
	"modulename/global"

	"gopkg.in/yaml.v3"
)

// 读取yaml文件的配置
func InitConf() {
	const name = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(name)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error : %v", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("yamlConf Unmarshal error : %v", err)
	}
	log.Println("config yamlFile load Init success")
	fmt.Printf("%#v\n", c)
	global.Config = c
}
