package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/auto-staging/stagectl/model"

	"github.com/auto-staging/tower/types"

	"github.com/spf13/cobra"
)

func stopEnvironmentCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Please specify the repository and branch you want to stop the environment for, check 'stagectl stop environment -h' for more info")
		os.Exit(1)
	}

	repository := args[0]
	branch := args[1]

	trigger := types.TriggerSchedulePost{
		Action:     "stop",
		Branch:     branch,
		Repository: repository,
	}

	body, err := json.Marshal(trigger)
	if err != nil {
		os.Exit(1)
	}

	err = model.TriggerSchedule(body)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Successfully invoked scheduler for stop")
}
