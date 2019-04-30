//
// 使用 cobra add command -p parentCmd 添加子命令(deault: parentCmd = rootCmd)
//
package ip

import (
	"github.com/spf13/cobra"
)

// publicCmd represents the public command
var publicCmd = &cobra.Command{
	Use:   "public",
	Short: "获取 公网IP",
	Long:  `访问 http://ip.cip.cc 获取公网 IP`,
	Run: func(cmd *cobra.Command, args []string) {
		getPublicIP()
	},
}

func init() {
	//
	// 此处: 将命令注册到上级父命令
	IpCmd.AddCommand(publicCmd)
}

//
// 复用 ip.go 中的 getPublicIP() 函数
//
//func PublicIP() {
//	body := httpGet("http://ip.cip.cc")
//	fmt.Printf("%s", body)
//}
