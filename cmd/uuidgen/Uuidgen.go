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

package uuidgen

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// Flags
var (
	Upper  bool
	Number int
)

// 这里可以不用通道 ch 的，使用的目的，只是为了练习
// 如果使用循环， 看 ./Uuidgen.go.tmp
// 通道一定要使用 make 进行初始化。
// 最开始声明 ch  `var ch chan uuid.UUID` ; 但是没有初始化,  在 printer 中始终拿不到信息，实际上应该是， uuidgen 就没有传入。
var ch = make(chan uuid.UUID, 3)

// UuidCmd represents the uuid command
var UuidCmd = &cobra.Command{
	Use:   "uuidgen",
	Short: "generate uuid",
	Long:  `generate uuid for you`,
	Run: func(cmd *cobra.Command, args []string) {

		uuidMain()

	},
}

func init() {

	UuidCmd.Flags().BoolVarP(&Upper, "upper", "u", false, "Generate UPPER CASE uuid, (default lower)")
	UuidCmd.Flags().IntVarP(&Number, "num", "n", 1, "number of UUID to be generated")
}

func uuidgen(num int) {

	// https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/16.4.md
	// - 切片、映射和通道，使用make
	// - 数组、结构体和所有的值类型，使用new
	// var ch1 chan string
	for i := 0; i < num; i++ {
		id := uuid.New()
		ch <- id
	}
	close(ch)

}

func printer(upper bool) {

	for id := range ch {
		s := id.String()
		if upper {
			s = strings.ToUpper(s)
		}

		fmt.Println(s)
	}
	//wg.Done()
}

func uuidMain() {
	go uuidgen(Number)
	printer(Upper)
}
