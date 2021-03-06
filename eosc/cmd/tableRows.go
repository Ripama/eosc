// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"encoding/json"
	"fmt"

	"os"

	"github.com/eoscanada/eos-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tableRowsCmd = &cobra.Command{
	Use:   "rows [contract] [scope] [table]",
	Short: "List the producers",
	Long:  `List the producers`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		api := api()

		response, err := api.GetTableRows(
			eos.GetTableRowsRequest{
				Code:  args[0],
				Scope: args[1],
				Table: args[2],
				JSON:  true,
				Limit: uint32(viper.GetInt("tableCmd.limit")),
			},
		)

		if err != nil {
			fmt.Printf("Get table rows , %s\n", err.Error())
			os.Exit(1)
		}

		data, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Printf("Error: json conversion , %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Println(string(data))

	},
}

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "table related commands",
}

func init() {
	// RootCmd.AddCommand(tableCmd)
	// tableCmd.AddCommand(tableRowsCmd)

	tableRowsCmd.Flags().IntP("limit", "", 50, "maximum producers that will be return")

	for _, flag := range []string{"limit"} {
		if err := viper.BindPFlag("tableCmd."+flag, tableRowsCmd.Flags().Lookup(flag)); err != nil {
			panic(err)
		}
	}

}
