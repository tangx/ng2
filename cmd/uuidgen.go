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
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"strings"
)

// Flags
var (
	Upper bool
	Number int
)



// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuidgen",
	Short: "generate uuid",
	Long: `generate uuid for you`,
	Run: func(cmd *cobra.Command, args []string) {


		uuidgen(Number)

	},
}


func init() {
	rootCmd.AddCommand(uuidCmd)


	uuidCmd.Flags().BoolVarP(&Upper,"upper","u",false,"Generate UPPER CASE uuid, (default lower)")
	uuidCmd.Flags().IntVarP(&Number,"num","n",1,"number of UUID to be generated")
}

func uuidgen(num int){

	for i:=0 ; i<=num; i++ {
		id:=uuid.New()
		s:=id.String()
		if Upper {
			s=strings.ToUpper(s)
		}
		fmt.Println(s)
	}

}
