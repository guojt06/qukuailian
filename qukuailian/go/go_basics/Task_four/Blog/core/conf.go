package core

import (
	"fmt"
	"io/fs"
	ioutil "io/ioutil"
	"log"
	"modulename/config"
	"modulename/global"

	"gopkg.in/yaml.v3"
)

const ConfigFile = "settings.yaml"

// 读取yaml文件的配置
func InitConf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
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

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
