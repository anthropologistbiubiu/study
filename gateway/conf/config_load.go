package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// 可以在这里定义结构体完成初始化
type Configs struct {
}

var ConfigData = &Configs{}

func LoadConfig() {
	configData, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	// 解析 YAML 配置并序列化到配置结构体
	fmt.Println("NNNNNNNN", string(configData))
	if err := yaml.Unmarshal(configData, ConfigData); err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", ConfigData)
}
