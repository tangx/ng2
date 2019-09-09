// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipinfo

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
	ApiURLDesc := ""
	for k, v := range ApiURL {
		ApiURLDesc += fmt.Sprintf("%v: %v\n", k, v)
	}

	//IpinfoCmd.Flags().StringVarP(&IPAddr, "ip", "", "", "specify ipaddr")
	IpinfoCmd.Flags().BoolVarP(&Oneline, "no-trunc", "", false, "单行输出")
	IpinfoCmd.Flags().StringVarP(&IPAPI, "api", "", "default", ApiURLDesc)
}

//
//const (
//	//ApiURL = "http://ip.taobao.com/service/getIpInfo.php?ip=%s"
//	//ApiURL = "http://ip-api.com/json/%s"
//	//ApiURL = "https://api.ttt.sh/ip/qqwry/%s"
//
//	ApiURLDesc = `指定查询 ApiURL
//ipapi: http://ip-api.com/json/%s
//taobao: http://ip.taobao.com/service/getIpInfo.php?ip=%s
//`
//)

var (
	IPAddr  string
	Oneline bool
	IPAPI   string
)

var ApiURL = map[string]string{
	"taobao":  "http://ip.taobao.com/service/getIpInfo.php?ip=%s",
	"ipapi":   "http://ip-api.com/json/%s",
	"default": "https://api.ttt.sh/ip/qqwry/%s",
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

//// 使用 re 正则表达式判断包是否合法
//func ipValid(ipAddress string) bool {
//	ipAddress = strings.Trim(ipAddress, " ")
//
//	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
//	if re.MatchString(ipAddress) {
//		return true
//	}
//	return false
//}

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
