package abc

/*
import (
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type NacosConfigClient struct {
	//ConfigInput
	//"ip1,ip2,ip3"
	IpAddrs         string
	Port            int
	NaocesNameSpace string
	DataId          string
	Group           string
	NacosLogDir     string
	NacosCacheDir   string
	//NamingInput
	NamingServiceName string
	NamingServiceIP   string
	NamingOwner       string
	NamingPort        int
	//Output
	ConfigClient config_client.IConfigClient
	NamingClient naming_client.INamingClient
}

type OptionNacos func(msg *NacosConfigClient)

// config naming
func OptionNacosConfigWithServiceName(serviceName string) OptionNacos {
	return func(msg *NacosConfigClient) {
		msg.NamingServiceName = serviceName
	}
}
func OptionNacosConfigWithServiceIP(serviceIP string) OptionNacos {
	return func(msg *NacosConfigClient) {
		msg.NamingServiceIP = serviceIP
	}
}

func OptionNacosConfigWithOwner(owner string) OptionNacos {
	return func(msg *NacosConfigClient) {
		msg.NamingOwner = owner
	}
}

// config client
func OptionNacosConfigWithLogDir(logDir string) OptionNacos {
	return func(msg *NacosConfigClient) {
		msg.NacosLogDir = logDir
	}
}

func OptionNacosConfigWithCacheDir(cacheDir string) OptionNacos {
	return func(msg *NacosConfigClient) {
		msg.NacosCacheDir = cacheDir
	}
}

func OptionNacosConfigWithIpaddr(ipAddrr string) OptionNacos {
	return func(msg *NacosConfigClient) {
		msg.IpAddrs = ipAddrr
	}
}
func OptionNacosConfigWithNameSpace(nameSpace string) OptionNacos {
	return func(msg *NacosConfigClient) {
		msg.NaocesNameSpace = nameSpace
	}
}
func OptionNacosConfigWithPort(port int) OptionNacos {
	return func(msg *NacosConfigClient) {
		msg.Port = port
	}
}

func OptionNacosConfigDataId(dataId string) OptionNacos {
	return func(msg *NacosConfigClient) {
		msg.DataId = dataId
	}
}

func OptionNacosConfigGroup(group string) OptionNacos {
	return func(msg *NacosConfigClient) {
		msg.Group = group
	}
}

func OptionNacosWithNamingPort(port int) OptionNacos {
	return func(msg *NacosConfigClient) {
		msg.NamingPort = port
	}
}

func (client *NacosConfigClient) GetConfig() (string, error) {
	content, err := client.ConfigClient.GetConfig(vo.ConfigParam{
		DataId: client.DataId,
		Group:  client.Group,
	})
	return content, err
}
*/
