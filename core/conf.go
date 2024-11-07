package core

import (
	"EStarExchange/config"
	"EStarExchange/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// 首先需要安装yaml包
// go get gopkg.in/yaml.v2

func InitConf() {
	const ConfigFile = "setting.yaml"
	c := &config.Config{}
	yamlFile, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println("config yamlFile is successfully.\n")
	global.Config = c
}
