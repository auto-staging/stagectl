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

func getEnvironmentCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please specify the repository you want to get the environments for, check 'stagectl get environments -h' for more info")
		os.Exit(1)
	}
	output := cmd.Flag("output").Value.String()
	repo := args[0]

	if cmd.Flag("limit").Value.String() == "" {
		envs, err := model.GetEnvironmentsForRepo(repo)
		if err != nil {
			log.Fatal(err)
		}
		outputEnvironmentsArray(envs, output)

	} else {
		branch := url.PathEscape(cmd.Flag("limit").Value.String())

		env, err := model.GetSingleEnvironmentForRepo(repo, branch)
		if err != nil {
			log.Fatal(err)
		}
		outputEnvironment(env, output)

	}

}

func outputEnvironmentsArray(envs []types.Environment, format string) {
	switch format {
	case "table":
		var data [][]string
		table := tablewriter.NewWriter(os.Stdout)
		for _, env := range envs {
			data = append(data, []string{env.Branch, env.CreationDate, env.Status, fmt.Sprint(env.ShutdownSchedules), fmt.Sprint(env.StartupSchedules)})
		}
		table.SetHeader([]string{"Branch", "Creation_Date", "Status", "Shutdown_Schedules", "Startup_Schedules"})
		for _, v := range data {
			table.Append(v)
		}
		fmt.Println("")
		table.SetRowLine(true)
		table.Render()
		fmt.Println("")
		return

	case "json":
		jsonBody, err := json.MarshalIndent(envs, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
		fmt.Print(string(jsonBody))
		fmt.Println("")
		return
	case "yaml":
		yamlBody, err := yaml.Marshal(envs)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
		fmt.Println(string(yamlBody))
		fmt.Println("")
		return
	}
}

func outputEnvironment(env types.Environment, format string) {
	switch format {
	case "table":
		var data [][]string
		table := tablewriter.NewWriter(os.Stdout)
		data = append(data, []string{env.Branch, env.CreationDate, env.Status, fmt.Sprint(env.ShutdownSchedules), fmt.Sprint(env.StartupSchedules)})
		table.SetHeader([]string{"Branch", "Creation_Date", "Status", "Shutdown_Schedules", "Startup_Schedules"})
		for _, v := range data {
			table.Append(v)
		}
		fmt.Println("")
		table.SetRowLine(true)
		table.Render()
		fmt.Println("")
		return

	case "json":
		jsonBody, err := json.MarshalIndent(env, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
		fmt.Print(string(jsonBody))
		fmt.Println("")
		return

	case "yaml":
		yamlBody, err := yaml.Marshal(env)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
		fmt.Println(string(yamlBody))
		fmt.Println("")
		return
	}
}
