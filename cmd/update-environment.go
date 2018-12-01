package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"gitlab.com/auto-staging/tower/types"
	yaml "gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
	"gitlab.com/auto-staging/stagectl/helper"
	"gitlab.com/auto-staging/stagectl/model"
)

func updateEnvironmentCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Please specify the repository and branch you want to edit the environment for, check 'stagectl update environment -h' for more info")
		return
	}

	repoName := args[0]
	branchName := args[1]

	env, err := model.GetSingleEnvironmentForRepo(repoName, url.PathEscape(branchName))
	if err != nil {
		log.Println(err)
	}

	envUpdate := types.EnvironmentPut{
		CodeBuildRoleARN:      env.CodeBuildRoleARN,
		EnvironmentVariables:  env.EnvironmentVariables,
		InfrastructureRepoURL: env.InfrastructureRepoURL,
		ShutdownSchedules:     env.ShutdownSchedules,
		StartupSchedules:      env.StartupSchedules,
	}

	helper.AskForEnvironmentUpdateInput(&envUpdate)

	body, err := json.Marshal(envUpdate)
	if err != nil {
		log.Fatal(err)
	}

	env, err = model.UpdateSingleEnvironment(repoName, url.PathEscape(branchName), body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully updated")

	yamlBody, err := yaml.Marshal(env)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")
	return
}
