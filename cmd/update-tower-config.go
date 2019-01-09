package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/auto-staging/stagectl/helper"
	"github.com/auto-staging/stagectl/model"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

func updateTowerConfigurationCmdFunc(cmd *cobra.Command, args []string) {

	config, err := model.GetTowerConfig()
	if err != nil {
		log.Fatal(err)
	}

	helper.AskForTowerConfigUpdateInput(&config)

	body, err := json.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	config, err = model.UpdateTowerConfiguration(body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully updated")

	yamlBody, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")
}
