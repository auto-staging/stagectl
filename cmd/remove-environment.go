package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/auto-staging/stagectl/model"
	"github.com/spf13/cobra"
)

func removeEnvironmentCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Please specify the repository and branch you want to delete the environment for, check 'stagectl delete environment -h' for more info")
		os.Exit(1)
	}

	repoName := args[0]
	branchName := args[1]

	err := model.DeleteSingleEnvironment(repoName, url.PathEscape(branchName))
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Successfully started deletion of environment")
}
