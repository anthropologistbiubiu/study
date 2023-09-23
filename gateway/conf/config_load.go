package conf

import (
	"fmt"
	protos "gateway/conf"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

// 可以在这里定义结构体完成初始化
type Configs struct {
}

var ConfigData = &Configs{}
var server = protos.Server{}

func LoadConfig() {
	configData, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	// 解析 YAML 配置并序列化到配置结构体
	if err := proto.UnmarshalText(string(configData), server); err != nil {
		panic(err)
	}
	// 现在我把这个配置怎么序列化位对应的结构体，使用的protobuf 协议
	// 然后使用这个服务封装 es 和 mysql 的交互
	fmt.Printf("%v \n", ConfigData)
}
