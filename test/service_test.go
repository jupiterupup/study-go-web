package test

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/yitter/idgenerator-go/idgen"
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	maps := make(map[string]string)
	maps["one"] = "one"
	fmt.Println(maps["one"])
}

func TestSnowFlake(t *testing.T) {
	fmt.Println(idgen.NextId())
}

func TestNacos(t *testing.T) {
	serverConfigs := []constant.ServerConfig{
		{IpAddr: "139.9.50.212", Port: 8848},
	}

	nacosClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &constant.ClientConfig{Username: "reader", Password: "Nacos@yf2021read", TimeoutMs: 5000, NamespaceId: "develop"},
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Fatal("nacos初始化错误:", err)
	}

	content, err := nacosClient.GetConfig(vo.ConfigParam{DataId: "fire-press.properties", Group: "1.0.0", Content: "dsn"})
	if err != nil {
		log.Fatalln("nacos读取配置错误:" + content)
	}

	fmt.Println("content", content)
}
