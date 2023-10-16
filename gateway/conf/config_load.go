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
var server = &Server{}

func LoadConfig() {
	configData, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	// 解析 YAML 配置并序列化到配置结构体
	//fmt.Println("string string(configData", string(configData))
	var configs = map[string]interface{}{}
	if err := yaml.Unmarshal(configData, configs); err != nil {
		fmt.Println("UUU", err)
	}
	serverKV, _ := configs["server"].(map[interface{}]interface{})
	//elasticsearchKV, _ := configs["elasticsearch"].(map[string]interface{})
	Port, _ := serverKV["port"].(string)
	fmt.Println("serverkv", serverKV, Port)
	//server.Port = int32(Port)
	server.Address, _ = serverKV["address"].(string)
	// 现在我把这个配置怎么序列化位对应的结构体，使用的protobuf 协议
	// 然后使用这个服务封装 es 和 mysql 的交互
	fmt.Println("server.Address", server.Address)
	fmt.Println("server.Port", server.Port)
}
