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

package pwgen

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// pwgenCmd represents the pwgen command
var PwgenCmd = &cobra.Command{
	Use:   "pwgen",
	Short: "产生多种随机密码",
	Long:  `产生多种随机密码`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pwgen called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pwgenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pwgenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Pwgen() {

}

const (
	LETTER_UPPER = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	DIGIT        = `1234567890`
	SYMBAL       = `!@#$%^&*(){})_+:?`
)

// 使用字典产生随机密码
func pwgen01(n int) {

	pool := LETTER_UPPER + strings.ToLower(LETTER_UPPER) + DIGIT + SYMBAL

	pw := []string{}
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())

		s := string(pool[rand.Intn(len(pool))])
		//fmt.Println(s)
		pw = append(pw, s)

	}
	fmt.Println(strings.Join(pw, ``))
}

func pwgenDevZero(n int) {

	fobj, err := os.Open("/dev/random")
	if err != nil {
		panic(err)
	}
	defer fobj.Close()

	buf := make([]byte, n)

	_, err = fobj.Read(buf)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%x", buf)
	fmt.Println(string(buf[:]))
}
