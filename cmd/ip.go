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
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "查询本机 IP 地址",
	Long:  `查询IP地址`,
	Run: func(cmd *cobra.Command, args []string) {

		ipMain()
	},
}

var (
	ipPublic bool
)

func init() {
	rootCmd.AddCommand(ipCmd)

	//ipCmd.Flags().BoolVarP(&ipPublic, "public", "p", false, "查询公网 IP")
}

func ipMain() {
	getLocalIP()
	getPublicIP()
}

func getPublicIP() {

	//resp, err := http.Get("http://ip.cip.cc")
	//if err != nil {
	//	logrus.Errorf("%v", err)
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	logrus.Errorf("%v", err)
	//}
	//fmt.Printf("Public IP:\n    %s", body)

	getPublicIP2()

}

func getLocalIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	fmt.Printf("Local IP:\n")
	for _, addr := range addrs {
		//fmt.Println(addr.String())
		fmt.Printf("    %s\n", addr.String())
	}

}

func getPublicIP2() {
	bodyByte := httpGet("http://ip.cip.cc")
	fmt.Printf("Public IP:\n    %s", bodyByte)
}
