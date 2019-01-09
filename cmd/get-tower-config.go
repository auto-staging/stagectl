package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/auto-staging/stagectl/model"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

func getTowerConfigurationCmdFunc(cmd *cobra.Command, args []string) {
	config, err := model.GetTowerConfig()
	if err != nil {
		log.Fatal(err)
	}

	switch cmd.Flag("output").Value.String() {
	case "yaml":
		yamlBody, err := yaml.Marshal(config)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
		fmt.Println(string(yamlBody))
		fmt.Println("")
		return
	case "json":
		jsonBody, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
		fmt.Print(string(jsonBody))
		fmt.Println("")
		return
	}
}
