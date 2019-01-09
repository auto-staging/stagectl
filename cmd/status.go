// Copyright © 2018 Jan Ritter <git@janrtr.de>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/auto-staging/stagectl/model"
	yaml "gopkg.in/yaml.v2"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Get the status of all environments",
	Example: "stagectl status",
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flag("repo").Value.String() != "" && cmd.Flag("branch").Value.String() != "" {
			singleStatus, err := model.GetSingleStatus(cmd.Flag("repo").Value.String(), url.PathEscape(cmd.Flag("branch").Value.String()))
			if err != nil {
				os.Exit(1)
			}

			switch cmd.Flag("output").Value.String() {
			case "table":
				var data [][]string

				data = append(data, []string{singleStatus.Repository, singleStatus.Branch, singleStatus.Status})

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Repository", "Branch", "Status"})

				for _, v := range data {
					table.Append(v)
				}

				table.SetColumnColor(tablewriter.Colors{tablewriter.FgWhiteColor},
					tablewriter.Colors{tablewriter.FgWhiteColor},
					tablewriter.Colors{tablewriter.FgWhiteColor})

				fmt.Println("")
				table.SetRowLine(true)
				table.Render()
				fmt.Println("")
				return
			case "yaml":
				yamlBody, err := yaml.Marshal(singleStatus)
				if err != nil {
					os.Exit(1)
				}
				fmt.Println("")
				fmt.Println(string(yamlBody))
				fmt.Println("")
				return
			case "json":
				jsonBody, err := json.MarshalIndent(singleStatus, "", "  ")
				if err != nil {
					os.Exit(1)
				}
				fmt.Println("")
				fmt.Print(string(jsonBody))
				fmt.Println("")
				return
			}
			return
		}

		status, err := model.GetAllStatus()
		if err != nil {
			os.Exit(1)
		}

		switch cmd.Flag("output").Value.String() {
		case "table":
			var data [][]string

			for _, singleStatus := range status {
				data = append(data, []string{singleStatus.Repository, singleStatus.Branch, singleStatus.Status})
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Repository", "Branch", "Status"})

			for _, v := range data {
				table.Append(v)
			}

			table.SetColumnColor(tablewriter.Colors{tablewriter.FgWhiteColor},
				tablewriter.Colors{tablewriter.FgWhiteColor},
				tablewriter.Colors{tablewriter.FgWhiteColor})

			fmt.Println("")
			table.SetRowLine(true)
			table.Render()
			fmt.Println("")
			return
		case "yaml":
			yamlBody, err := yaml.Marshal(status)
			if err != nil {
				os.Exit(1)
			}
			fmt.Println("")
			fmt.Println(string(yamlBody))
			fmt.Println("")
			return
		case "json":
			jsonBody, err := json.MarshalIndent(status, "", "  ")
			if err != nil {
				os.Exit(1)
			}
			fmt.Println("")
			fmt.Print(string(jsonBody))
			fmt.Println("")
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringP("repo", "r", "", "Limit output to a specific environment by repo, also requires branch - example: '--repo demo-app'")
	statusCmd.Flags().StringP("branch", "b", "", "Limit output to a specific environment by branch, also requires repo - example: '--branch feat/new-ui'")
	statusCmd.Flags().StringP("output", "o", "table", "Format of the output, options are table / yaml / json")
}
