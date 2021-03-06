package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/auto-staging/stagectl/model"
	"github.com/auto-staging/tower/types"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

func getRepositoriesCmdFunc(cmd *cobra.Command, args []string) {
	output := cmd.Flag("output").Value.String()
	repoName := url.PathEscape(cmd.Flag("limit").Value.String())

	if repoName == "" { // All repos
		repos, err := model.GetAllRepositories()
		if err != nil {
			log.Fatal(err)
		}
		outputRepositoriesArray(repos, output)

	} else { // Single repo
		repo, err := model.GetSingleRepository(repoName)
		if err != nil {
			log.Fatal(err)
		}
		outputRepository(repo, output)
	}
}

func outputRepositoriesArray(repos []types.Repository, format string) {
	switch format {
	case "table":
		var data [][]string
		table := tablewriter.NewWriter(os.Stdout)
		for _, repo := range repos {
			data = append(data, []string{repo.Repository, repo.InfrastructureRepoURL, fmt.Sprint(repo.Webhook), fmt.Sprint(repo.Filters), fmt.Sprint(repo.ShutdownSchedules), fmt.Sprint(repo.StartupSchedules), repo.CodeBuildRoleARN})
		}
		table.SetHeader([]string{"Repository", "InfrastructureRepoURL", "Webhook", "Filters", "ShutdownSchedules", "StartupSchedules", "CodeBuildRoleARN"})
		for _, v := range data {
			table.Append(v)
		}
		fmt.Println("")
		table.SetRowLine(true)
		table.Render()
		fmt.Println("")
		return
	case "json":
		jsonBody, err := json.MarshalIndent(repos, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
		fmt.Print(string(jsonBody))
		fmt.Println("")
		return
	case "yaml":
		yamlBody, err := yaml.Marshal(repos)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
		fmt.Println(string(yamlBody))
		fmt.Println("")
		return
	}
}

func outputRepository(repo types.Repository, format string) {
	switch format {
	case "table":
		var data [][]string
		table := tablewriter.NewWriter(os.Stdout)
		data = append(data, []string{repo.Repository, repo.InfrastructureRepoURL, fmt.Sprint(repo.Webhook), fmt.Sprint(repo.Filters), fmt.Sprint(repo.ShutdownSchedules), fmt.Sprint(repo.StartupSchedules), repo.CodeBuildRoleARN})
		table.SetHeader([]string{"Repository", "InfrastructureRepoURL", "Webhook", "Filters", "ShutdownSchedules", "StartupSchedules", "CodeBuildRoleARN"})
		for _, v := range data {
			table.Append(v)
		}
		fmt.Println("")
		table.SetRowLine(true)
		table.Render()
		fmt.Println("")
		return

	case "json":
		jsonBody, err := json.MarshalIndent(repo, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
		fmt.Print(string(jsonBody))
		fmt.Println("")
		return

	case "yaml":
		yamlBody, err := yaml.Marshal(repo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
		fmt.Println(string(yamlBody))
		fmt.Println("")
		return
	}
}
