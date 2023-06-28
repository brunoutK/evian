
/*
import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "172.18.10.36",
			Port:   8848,
		},
	}

	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "1cc67a8d-f2f3-45f5-aef7-f7dd4b451264", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           3000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// 创建动态配置客户端的另一种方式 (推荐)

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}

	//获取配置信息
	content, err := configClient.GetConfig(vo.ConfigParam{DataId: "ops-phone", Group: "DEFAULT_GROUP"})
	if err != nil {
		fmt.Println("GetConfig err: ", err)
	}
	fmt.Print(content)

	//监听配置
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "ops-phone",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		return
	}
	//	time.Sleep(time.Second * 1000)
}

*/