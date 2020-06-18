package wifipasswd

// 查看 mac wifi 密码
import (
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// WpCmd represents the wp command
var WpCmd = &cobra.Command{
	Use:   "wp",
	Short: "查看当前 wifi password",
	Long:  `查看当前 Wi-Fi 密码， 目前只支持 MacOS`,
	Run: func(cmd *cobra.Command, args []string) {
		GetSSID()
	},
}

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
