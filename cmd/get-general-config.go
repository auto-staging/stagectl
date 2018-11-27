package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/auto-staging/stagectl/model"
	yaml "gopkg.in/yaml.v2"
)

func getGeneralConfigurationCmdFunc(cmd *cobra.Command, args []string) {
	config, err := model.GetGeneralConfig()
	if err != nil {
		log.Println(err)
	}

	switch cmd.Flag("output").Value.String() {
	case "yaml":
		yamlBody, err := yaml.Marshal(config)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("")
		fmt.Println(string(yamlBody))
		fmt.Println("")
		return
	case "json":
		jsonBody, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Println(err)
		}
		fmt.Println("")
		fmt.Print(string(jsonBody))
		fmt.Println("")
		return
	}
}
