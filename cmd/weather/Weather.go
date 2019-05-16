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

package weather

import (
	"fmt"

	"github.com/tangx/ng2/utils"

	"github.com/spf13/cobra"
)

// WeatherCmd represents the Weather command
var WeatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		weather(args[0])
	},
}

var pretty bool
var NoTrunc bool

func init() {
	//WeatherCmd.Flags().BoolVarP(&pretty, "pretty", "p", true, "单行输出")
	WeatherCmd.Flags().BoolVarP(&NoTrunc, "no-trunc", "", false, "单行输出")

}
func weather(city string) {

	//data := utils.HttpGet("http://www.mxnzp.com/api/weather/current/" + city)

	//data := utils.HttpGet("http://www.mxnzp.com/api/weather/current/%E6%B7%B1%E5%9C%B3")
	data := utils.HttpGet("http://www.mxnzp.com/api/weather/current/" + city)

	data = utils.PrettyJson(data, !NoTrunc)

	fmt.Printf("%s", data)

}
