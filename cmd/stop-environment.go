package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"gitlab.com/auto-staging/stagectl/model"

	"gitlab.com/auto-staging/tower/types"

	"github.com/spf13/cobra"
)

func stopEnvironmentCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Please specify the repository and branch you want to stop the environment for, check 'stagectl stop environment -h' for more info")
		return
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
		log.Fatal(err)
	}

	err = model.TriggerSchedule(body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully invoked scheduler for stop")
}
