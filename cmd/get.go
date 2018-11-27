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
	"fmt"
	"log"
	"net/url"
	"os"

	"gitlab.com/auto-staging/stagectl/model"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information about a specific resource",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get the available subcommands / resources with 'stagectl get -h'")
	},
}

var getEnvironmentCmd = &cobra.Command{
	Use:   "environments",
	Short: "Get all environments for a repository",
	Long: `Usage:

'stagectl get environments demo-app', where demo-app is your repository`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify the repository you want to get the environments for, check 'stagectl get environments -h' for more info")
			return
		}
		repo := args[0]

		if cmd.Flag("limit").Value.String() == "" {
			envs, err := model.GetEnvironmentsForRepo(repo)
			if err != nil {
				log.Println(err)
				return
			}

			data := [][]string{}
			table := tablewriter.NewWriter(os.Stdout)

			if cmd.Flag("wide").Value.String() == "true" {
				for _, env := range envs {
					data = append(data, []string{env.Branch, env.CreationDate, env.Status, fmt.Sprint(env.ShutdownSchedules), fmt.Sprint(env.StartupSchedules), fmt.Sprint(env.EnvironmentVariables)})
				}
				table.SetHeader([]string{"Branch", "Creation_Date", "Status", "Shutdown_Schedules", "Startup_Schedules", "environment_Variables"})
				for _, v := range data {
					table.Append(v)
				}
			} else {
				for _, env := range envs {
					data = append(data, []string{env.Branch, env.CreationDate, env.Status, fmt.Sprint(env.ShutdownSchedules), fmt.Sprint(env.StartupSchedules)})
				}
				table.SetHeader([]string{"Branch", "Creation_Date", "Status", "Shutdown_Schedules", "Startup_Schedules"})
				for _, v := range data {
					table.Append(v)
				}
			}

			fmt.Println("")
			table.SetRowLine(true)
			table.Render()
			fmt.Println("")
		} else {
			branch := url.PathEscape(cmd.Flag("limit").Value.String())

			env, err := model.GetSingleEnvironmentForRepo(repo, branch)
			if err != nil {
				log.Println(err)
				return
			}

			data := [][]string{}
			table := tablewriter.NewWriter(os.Stdout)

			if cmd.Flag("wide").Value.String() == "true" {

				data = append(data, []string{env.Branch, env.CreationDate, env.Status, fmt.Sprint(env.ShutdownSchedules), fmt.Sprint(env.StartupSchedules), fmt.Sprint(env.EnvironmentVariables)})

				table.SetHeader([]string{"Branch", "Creation_Date", "Status", "Shutdown_Schedules", "Startup_Schedules", "environment_Variables"})
				for _, v := range data {
					table.Append(v)
				}
			} else {
				data = append(data, []string{env.Branch, env.CreationDate, env.Status, fmt.Sprint(env.ShutdownSchedules), fmt.Sprint(env.StartupSchedules)})

				table.SetHeader([]string{"Branch", "Creation_Date", "Status", "Shutdown_Schedules", "Startup_Schedules"})
				for _, v := range data {
					table.Append(v)
				}
			}

			fmt.Println("")
			table.SetRowLine(true)
			table.Render()
			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getEnvironmentCmd)

	getEnvironmentCmd.Flags().BoolP("wide", "w", false, "Enhanced output")
	getEnvironmentCmd.Flags().StringP("limit", "l", "", "Limit output to a specific environment by branch - example: '--limit feat/new-ui'")
}
