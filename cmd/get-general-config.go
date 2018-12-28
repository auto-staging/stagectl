package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/auto-staging/stagectl/model"
	yaml "gopkg.in/yaml.v2"
)

func getGeneralConfigurationCmdFunc(cmd *cobra.Command, args []string) {
	config, err := model.GetGeneralConfig()
	if err != nil {
		os.Exit(1)
	}

	switch cmd.Flag("output").Value.String() {
	case "yaml":
		yamlBody, err := yaml.Marshal(config)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println("")
		fmt.Println(string(yamlBody))
		fmt.Println("")
		return
	case "json":
		jsonBody, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			os.Exit(1)
		}
		fmt.Println("")
		fmt.Print(string(jsonBody))
		fmt.Println("")
		return
	}
}
