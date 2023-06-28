package main

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func getNacosConfig() (result int) {
	// 从控制台命名空间管理的"命名空间详情"中拷贝 End Point、命名空间 ID
	var endpoint = "172.18.10.36:80/nacos"
	var namespaceId = "1cc67a8d-f2f3-45f5-aef7-f7dd4b451264"
	// 推荐使用 RAM 用户的 accessKey、secretKey
	//var accessKey = "${accessKey}"
	//var secretKey = "${secretKey}"
	clientConfig := constant.ClientConfig{
		Endpoint:    endpoint,
		NamespaceId: namespaceId,
		//AccessKey:      accessKey,
		//SecretKey:      secretKey,
		Username:       "nacos",
		Password:       "heavengifts@2023",
		TimeoutMs:      10 * 1000, //http请求超时时间，单位毫秒
		ListenInterval: 30 * 1000, //监听间隔时间，单位毫秒（仅在ConfigClient中有效）
		BeatInterval:   5 * 1000,  //心跳间隔时间，单位毫秒（仅在ServiceClient中有效）
	}

	// 至少一个
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "172.18.10.36",
			ContextPath: "/nacos",
			Port:        80,
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	var dataId = "ops-phone"
	var group = "DEFAULT_GROUP"

	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	fmt.Println("Get config：" + content)
	if len(content) == 0 {
		fmt.Println("获取配置为空")
	}
	return len(content)
}

func main() {
	re := getNacosConfig()
	fmt.Println(re)
}
