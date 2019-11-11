package ipinfo

// 查询 IP 地址信息

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/tangx/ng2/utils"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	// IpinfoCmd represents the ipinfo command
	IpinfoCmd = &cobra.Command{
		Use:   "ip",
		Short: "查询 ip 地址详细信息",
		Long:  `使用 ip.taobao.com api 接口，查询 ip 地址详细信息`,
		Run: func(cmd *cobra.Command, args []string) {
			//ipinfo(IPAddr)
			//ipinfo(args[0])
			ipsInfo(args)

		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a color argument")
			}
			return nil
			//return fmt.Errorf("invalid color specified: %s", args[0])
		},
	}
)

func init() {
	// ApiURLDesc 帮助描述
	ApiURLDesc := ""
	for k, v := range ApiURL {
		ApiURLDesc += fmt.Sprintf("%v: %v\n", k, v)
	}

	//IpinfoCmd.Flags().StringVarP(&IPAddr, "ip", "", "", "specify ipaddr")
	IpinfoCmd.Flags().BoolVarP(&Oneline, "no-trunc", "", false, "单行输出")
	IpinfoCmd.Flags().StringVarP(&IPAPI, "api", "", "default", ApiURLDesc)
}

var (
	// IPAddr IP 地址
	IPAddr string
	// Oneline 单行输出
	Oneline bool
	// IPAPI api 接口地址
	IPAPI string
)

// ApiURL api 接口地址
var ApiURL = map[string]string{
	"taobao":  "http://ip.taobao.com/service/getIpInfo.php?ip=%s",
	"ipapi":   "http://ip-api.com/json/%s",
	"default": "https://www.mxnzp.com/api/ip/aim_ip?ip=%s",
	"ipsb":    "https://api.ip.sb/geoip/%s",
}

func ipinfo(ip string) {

	url := fmt.Sprintf(ApiURL[IPAPI], ip)
	logrus.Info(url)

	bodyByte := utils.HttpGet(url)

	//logrus.Infof("%s", bodyByte)
	js := unmarshal(bodyByte)

	// oneline
	if Oneline {
		data, _ := json.Marshal(js)
		fmt.Printf("%s\n", data)
		return
	}

	data, _ := json.MarshalIndent(js, "", "  ")
	fmt.Printf("%s\n", data)
	// pretty indent
}

func unmarshal(data []byte) interface{} {
	var js interface{}
	err := json.Unmarshal(data, &js)
	if err != nil {
		logrus.Fatal("%s", err)
	}
	return js
}

// 使用 net package 判断 ip 地址是否合法
func ipValid(ipaddr string) bool {
	if net.ParseIP(ipaddr) != nil {
		return true
	}
	return false
}

func ipsInfo(ips []string) {
	for _, ip := range ips {
		// ip valid check
		if ipValid(ip) {
			ipinfo(ip)
		} else {
			logrus.Error("Invalid ip address: ", ip)
		}
	}
}
