package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"gitlab.com/auto-staging/tower/types"

	"github.com/spf13/cobra"
	input "github.com/tcnksm/go-input"
	"gitlab.com/auto-staging/stagectl/helper"
	"gitlab.com/auto-staging/stagectl/model"
	yaml "gopkg.in/yaml.v2"
)

func addEnvironmentCmdFunc(cmd *cobra.Command, args []string) {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	if len(args) == 0 {
		fmt.Println("Please specify the repository you want to add the environment for, check 'stagectl add environment -h' for more info")
		os.Exit(1)
	}

	repoName := args[0]

	fmt.Printf("Please provide the information for the new environment: \n \n")
	envAdd := types.EnvironmentPost{}
	helper.AskForEnvironmentAddInput(&envAdd)
	yamlBody, err := yaml.Marshal(envAdd)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")

	query := "Do you want to add this environment to " + repoName + "? (yes or no)"
	decision, err := ui.Ask(query, &input.Options{
		Default: "no",
	})
	if err != nil {
		os.Exit(1)
	}
	if decision == "no" {
		return
	}

	body, err := json.Marshal(envAdd)
	if err != nil {
		os.Exit(1)
	}

	env, err := model.AddEnvironment(repoName, body)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Successfully added")

	yamlBody, err = yaml.Marshal(env)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")
}
