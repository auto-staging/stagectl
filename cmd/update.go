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

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing resource",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		cmd.Help()
	},
}

var updateRepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Update an existing repository",
	Run:   updateRepositoryCmdFunc,
}

var updateEnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Update an existing environment",
	Run:   updateEnvironmentCmdFunc,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updateRepositoryCmd)
	updateCmd.AddCommand(updateEnvironmentCmd)

	updateRepositoryCmd.Flags().StringP("input-file", "i", ".auto-staging.json", "Filename of the repository definition")
}
