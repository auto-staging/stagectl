package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gitlab.com/auto-staging/stagectl/model"
	yaml "gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

func addRepositoryCmdFunc(cmd *cobra.Command, args []string) {
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

	output, err := model.AddRepository(byteValue)
	if err != nil {
		log.Println("Failed")
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
