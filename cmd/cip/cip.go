package cip

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/ng2/utils"
)

// ipCmd represents the ip command
var IpCmd = &cobra.Command{
	Use:   "cip",
	Short: "查询本机 IP 地址",
	Long:  `查询IP地址`,
	Run: func(cmd *cobra.Command, args []string) {

		ipMain()
	},
}

var (
	ipPublic bool
)

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
	bodyByte := utils.HttpGet("http://api.ip.sb/ip")
	fmt.Printf("Public IP:\n    %s", bodyByte)
}
