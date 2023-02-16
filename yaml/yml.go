package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
	}
	MemCache struct {
		Enable bool     `yaml:"enable"`
		List   []string `yaml:"list"`
	}
	Mysql struct {
		User     string `yaml:"user"`
		PassWord string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int32  `yaml:"port"`
		DbName   string `yaml:"dbname"`
	}
}

func main() {
	//var config Config
	File, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("读取配置文件失败 %+v", err)
		return
	}
	fmt.Printf("File %+v", string(File))
	var mp map[string]interface{}
	err = yaml.Unmarshal(File, &mp)
	if err != nil {
		log.Fatalf("解析失败: %v", err)
		return
	}
	fmt.Printf("%+v", mp)
}
