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
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information about a specific resource",
}

var getEnvironmentCmd = &cobra.Command{
	Use:     "environments",
	Short:   "Get all environments for a repository",
	Example: "stagectl get environments my-repository",
	Run:     getEnvironmentCmdFunc,
}

var getRepositoriesCmd = &cobra.Command{
	Use:     "repositories",
	Short:   "Get all repositories",
	Example: "stagectl get repositories",
	Run:     getRepositoriesCmdFunc,
}

var getTowerConfigurationCmd = &cobra.Command{
	Use:     "tower-configuration",
	Short:   "Get current Tower configuration",
	Example: "stagectl get tower-configuration",
	Run:     getTowerConfigurationCmdFunc,
}

var getGeneralConfigurationCmd = &cobra.Command{
	Use:     "general-configuration",
	Short:   "Get the general default configuration for new repositories",
	Example: "stagectl get general-configuration",
	Run:     getGeneralConfigurationCmdFunc,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getEnvironmentCmd)
	getCmd.AddCommand(getRepositoriesCmd)
	getCmd.AddCommand(getTowerConfigurationCmd)
	getCmd.AddCommand(getGeneralConfigurationCmd)

	getEnvironmentCmd.Flags().StringP("limit", "l", "", "Limit output to a specific environment by branch - example: '--limit feat/new-ui'")
	getEnvironmentCmd.Flags().StringP("output", "o", "table", "Format of the output, options are table (limited output) / yaml / json")

	getRepositoriesCmd.Flags().StringP("limit", "l", "", "Limit output to a specific repository - example: '--limit demo-app'")
	getRepositoriesCmd.Flags().StringP("output", "o", "table", "Format of the output, options are table (limited output) / yaml / json")

	getGeneralConfigurationCmd.Flags().StringP("output", "o", "yaml", "Format of the output, options are yaml / json")

	getTowerConfigurationCmd.Flags().StringP("output", "o", "yaml", "Format of the output, options are yaml / json")
}
