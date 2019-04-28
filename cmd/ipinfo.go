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

package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// ipinfoCmd represents the ipinfo command
var ipinfoCmd = &cobra.Command{
	Use:   "ipinfo",
	Short: "查询 ip 地址详细信息",
	Long:  `使用 ip.taobao.com api 接口，查询 ip 地址详细信息`,
	Run: func(cmd *cobra.Command, args []string) {
		ipinfo(IPAddr)

	},
}

func init() {
	rootCmd.AddCommand(ipinfoCmd)

	ipinfoCmd.Flags().StringVarP(&IPAddr, "ip", "", "", "specify ipaddr")
	ipinfoCmd.Flags().BoolVarP(&Oneline, "no-trunc", "", false, "单行输出")
	ipinfoCmd.Flags().StringVarP(&IPAPI, "api", "", ApiURL, ApiURLDesc)
}

const (
	//ApiURL = "http://ip.taobao.com/service/getIpInfo.php?ip=%s"
	//ApiURL = "http://ip-api.com/json/%s"
	ApiURL = "https://api.ttt.sh/ip/qqwry/%s"

	ApiURLDesc = `指定查询 ApiURL
海外: http://ip-api.com/json/%s
淘宝: http://ip.taobao.com/service/getIpInfo.php?ip=%s 
`
)

var (
	IPAddr  string
	Oneline bool
	IPAPI   string
)

func ipinfo(ip string) {
	url := fmt.Sprintf(ApiURL, ip)
	logrus.Info(url)

	bodyByte := httpGet(url)

	//logrus.Infof("%s", bodyByte)
	js := unmarshal(bodyByte)

	// oneline
	if Oneline {
		data, _ := json.Marshal(js)
		fmt.Printf("%s", data)
		return
	}

	data, _ := json.MarshalIndent(js, "", "  ")
	fmt.Printf("%s", data)
	// pretty indent
}

func unmarshal(data []byte) interface{} {
	var js interface{}
	err := json.Unmarshal(data, &js)
	if err != nil {
		logrus.Fatalf("%s", err)
	}
	return js
}
