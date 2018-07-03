// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

// +build windows

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/svc"
)

// winbgCmd represents the winbg command
var winbgCmd = &cobra.Command{
	Use:   "winbg",
	Short: "not implemented yet",
	Long: `Not
implemented
yet.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("winbg is not implemented yet")

		const svcName = "myservice"

		isIntSess, err := svc.IsAnInteractiveSession()
		if err != nil {
			fmt.Println("boo")
		}
		if !isIntSess {
			fmt.Println("bar")
		}
	},
}

func init() {
	rootCmd.AddCommand(winbgCmd)
}
