package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/auto-staging/stagectl/model"
	yaml "gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

func addRepositoryCmdFunc(cmd *cobra.Command, args []string) {
	inputFileName := cmd.Flag("input-file").Value.String()

	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Using definition file: " + inputFileName)
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	output, err := model.AddRepository(byteValue)
	if err != nil {
		log.Println("Failed")
		log.Fatal(err)
	}

	yamlBody, err := yaml.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")
}
