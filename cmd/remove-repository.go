package cmd

import (
	"fmt"
	"os"

	"gitlab.com/auto-staging/stagectl/model"

	"github.com/spf13/cobra"
)

func removeRepositoryCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please specify the repository you want to delete, check 'stagectl remove repository -h' for more info")
		os.Exit(1)
	}

	err := model.DeleteRepository(args[0])
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Successfully deleted")
}
