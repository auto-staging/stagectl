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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new resource",
}

var addRepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Add a repository",
	Long:  "Uses the file .auto-staging.json in the current directory as fallback if not otherwise specified.",
	Run:   addRepositoryCmdFunc,
}

var addEnvironmentCmd = &cobra.Command{
	Use:     "environment",
	Short:   "Add an environmemt for repository",
	Example: "stagectl add environment my-repository",
	Run:     addEnvironmentCmdFunc,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addRepositoryCmd)
	addCmd.AddCommand(addEnvironmentCmd)

	addRepositoryCmd.Flags().StringP("input-file", "i", ".auto-staging.json", "Filename of the repository definition")
}
