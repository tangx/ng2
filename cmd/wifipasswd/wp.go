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

package wifipasswd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// WpCmd represents the wp command
var WspCmd = &cobra.Command{
	Use:   "wp",
	Short: "查看当前 wifi password",
	Long:  `查看当前 Wi-Fi 密码， 目前只支持 Darwin`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wp called")
	},
}

//func init() {
//	//rootCmd.AddCommand(wpCmd)
//
//	// Here you will define your flags and configuration settings.
//
//	// Cobra supports Persistent Flags which will work for this command
//	// and all subcommands, e.g.:
//	// wpCmd.PersistentFlags().String("foo", "", "A help for foo")
//
//	// Cobra supports local flags which will only run when this command
//	// is called directly, e.g.:
//	// wpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//}
//

// GetSSID return current Wifi SSID
func GetSSID() (ssid string) {

	// 执行系统命令
	// 第一个参数是命令名称
	// 后面参数可以有多个，命令参数
	cmd := exec.Command("/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", "-I")
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(opBytes))
	//return string(opBytes)

	re := regexp.MustCompile(` SSID: ([a-zA-Z0-9].+)`)

	s := re.FindString(string(opBytes))

	return s
}
