package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/auto-staging/stagectl/model"

	"github.com/auto-staging/tower/types"

	"github.com/spf13/cobra"
)

func startEnvironmentCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Please specify the repository and branch you want to start the environment for, check 'stagectl start environment -h' for more info")
		os.Exit(1)
	}

	repository := args[0]
	branch := args[1]

	trigger := types.TriggerSchedulePost{
		Action:     "start",
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

	fmt.Println("Successfully invoked scheduler for start")
}
