package cmd

import (
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/cobra"
	"gitlab.com/auto-staging/stagectl/model"
)

func removeEnvironmentCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Please specify the repository and branch you want to delete the environment for, check 'stagectl delete environment -h' for more info")
		return
	}

	repoName := args[0]
	branchName := args[1]

	err := model.DeleteSingleEnvironment(repoName, url.PathEscape(branchName))
	if err != nil {
		log.Println("Error")
		log.Println(err)
		return
	}

	fmt.Println("Successfully started deletion of environment")
}
