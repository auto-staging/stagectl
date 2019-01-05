package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/auto-staging/tower/types"

	"github.com/spf13/cobra"
	"gitlab.com/auto-staging/stagectl/model"
	yaml "gopkg.in/yaml.v2"
)

func updateRepositoryCmdFunc(cmd *cobra.Command, args []string) {
	inputFileName := cmd.Flag("input-file").Value.String()

	file, err := os.Open(inputFileName)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Using definition file: " + inputFileName)
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		os.Exit(1)
	}

	inputRepo := types.Repository{}
	err = json.Unmarshal(byteValue, &inputRepo)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Updating " + inputRepo.Repository)

	output, err := model.UpdateRepository(byteValue, inputRepo.Repository)
	if err != nil {
		os.Exit(1)
	}

	yamlBody, err := yaml.Marshal(output)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")
}
