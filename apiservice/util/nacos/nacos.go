package nacos

import (
	"fire-press/apiservice/util/viperhelper"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

const (
	ip       = "nacos.ip"
	port     = "nacos.port"
	username = "nacos.username"
	password = "nacos.password"
	appId    = "nacos.appId"
	appVer   = "nacos.appVer"
	appStage = "nacos.appStage"
)

var client config_client.IConfigClient

func init() {
	port, _ := strconv.Atoi(viperhelper.GetLocalConfIfPresent(port))
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: viperhelper.GetLocalConfIfPresent(ip),
			Port:   uint64(port),
		},
	}

	clientConfigs, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				Username:    viperhelper.GetLocalConfIfPresent(username),
				Password:    viperhelper.GetLocalConfIfPresent(password),
				NamespaceId: viperhelper.GetLocalConfIfPresent(appStage),
				TimeoutMs:   5000,
			},
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(fmt.Errorf("Fatal init naos: %s \n", err))
	}

	client = clientConfigs
}

func GetConfIfPresent(key string) string {
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: viperhelper.GetLocalConfIfPresent(appId),
		Group:  viperhelper.GetLocalConfIfPresent(appVer),
	})
	if err != nil {
		panic(fmt.Errorf("Fatal read naos: %s \n", err))
	}
	// TODO: 解析nacos的配置。注意配置改变是否同步生效？
	v := viper.New()
	err = v.ReadConfig(strings.NewReader(content))
	if err != nil {
		panic(fmt.Errorf("viper failed to resolve configuration: %s \n", err))
	}

	return v.GetString(key)
}
