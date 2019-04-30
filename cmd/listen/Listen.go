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

package listen

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// listenCmd represents the listen command
var ListenCmd = &cobra.Command{
	Use:   "listen",
	Short: "查看本机监听",
	Long:  `查看本机端口监听`,
	Run: func(cmd *cobra.Command, args []string) {
		listen()
	},
}

func listen() {

	var cmd *exec.Cmd

	// 获取操作系统
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("lsof", "-nP", "-iTCP", "-sTCP:LISTEN")
	case "linux":
		cmd = exec.Command("netstat", "-tunpl")
	default:
		return
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Errorf("%v", err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		logrus.Fatalf("%v", err)
	}

	str, err := ioutil.ReadAll(stdout)
	if err != nil {
		logrus.Fatalf("%v", err)
	}
	fmt.Printf("%s", string(str))

}
