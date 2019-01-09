package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/auto-staging/tower/types"
	yaml "gopkg.in/yaml.v2"

	"github.com/auto-staging/stagectl/helper"
	"github.com/auto-staging/stagectl/model"
	"github.com/spf13/cobra"
)

func updateEnvironmentCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Please specify the repository and branch you want to edit the environment for, check 'stagectl update environment -h' for more info")
		os.Exit(1)
	}

	repoName := args[0]
	branchName := args[1]

	env, err := model.GetSingleEnvironmentForRepo(repoName, url.PathEscape(branchName))
	if err != nil {
		os.Exit(1)
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
		os.Exit(1)
	}

	env, err = model.UpdateSingleEnvironment(repoName, url.PathEscape(branchName), body)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Successfully updated")

	yamlBody, err := yaml.Marshal(env)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")
}
