package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/auto-staging/tower/types"

	"github.com/auto-staging/stagectl/helper"
	"github.com/auto-staging/stagectl/model"
	"github.com/spf13/cobra"
	input "github.com/tcnksm/go-input"
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
		log.Fatal(err)
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
		log.Fatal(err)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")
}
