package cmd

import (
	"encoding/json"
	"fmt"
	"log"
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
		return
	}

	repoName := args[0]

	fmt.Printf("Please provide the information for the new environment: \n \n")
	envAdd := types.EnvironmentPost{}
	helper.AskForEnvironmentAddInput(&envAdd)
	yamlBody, err := yaml.Marshal(envAdd)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")

	query := "Do you want to add this environment to " + repoName + "? (yes or no)"
	decision, err := ui.Ask(query, &input.Options{
		Default: "no",
	})
	if err != nil {
		log.Fatal(err)
	}
	if decision == "no" {
		return
	}

	body, err := json.Marshal(envAdd)
	if err != nil {
		log.Fatal(err)
	}

	env, err := model.AddEnvironment(repoName, body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully added")

	yamlBody, err = yaml.Marshal(env)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")
	return
}
