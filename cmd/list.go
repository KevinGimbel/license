// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"github.com/kevingimbel/license/lib"
	"github.com/spf13/cobra"
	"os"
	"encoding/json"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List known licenses",
	Long: `The list command will list all known licenses read from the local cache file. To update the local cache, run the update command`,
	Run: func(cmd *cobra.Command, args []string) {
		data := []lib.OSLJSONFormat{}

		f, err := os.Open(lib.GetOutputFilePath())
		defer f.Close()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		json.NewDecoder(f).Decode(&data)

		for _, license := range data {
			fmt.Println(license.Name, "(id: ", license.ID, ")")
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
