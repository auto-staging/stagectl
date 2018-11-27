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
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gitlab.com/auto-staging/stagectl/helper"
	"gitlab.com/auto-staging/tower/types"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status of all environments",
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/environments/status", nil)
		if err != nil {
			log.Println(err)
		}

		helper.SignRequest(req)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
			log.Println(string(body))
		}

		body, err := ioutil.ReadAll(resp.Body)
		envs := []types.EnvironmentStatus{}
		err = json.Unmarshal([]byte(body), &envs)
		if err != nil {
			log.Panicln(err)
		}

		data := [][]string{}

		for _, env := range envs {
			data = append(data, []string{env.Repository, env.Branch, env.Status})
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
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
