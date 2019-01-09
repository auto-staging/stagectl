package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/auto-staging/stagectl/helper"
	"github.com/auto-staging/stagectl/model"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

func updateGeneralConfigurationCmdFunc(cmd *cobra.Command, args []string) {

	config, err := model.GetGeneralConfig()
	if err != nil {
		os.Exit(1)
	}

	helper.AskForGeneralConfigUpdateInput(&config)

	body, err := json.Marshal(config)
	if err != nil {
		os.Exit(1)
	}

	config, err = model.UpdateGeneralConfiguration(body)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Successfully updated")

	yamlBody, err := yaml.Marshal(config)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")
}
