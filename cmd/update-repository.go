package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		fmt.Println(err)
		return
	}
	fmt.Println("Using definition file: " + inputFileName)
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	inputRepo := types.Repository{}
	err = json.Unmarshal(byteValue, &inputRepo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Updating " + inputRepo.Repository)

	output, err := model.UpdateRepository(byteValue, inputRepo.Repository)
	if err != nil {
		log.Println("Failed")
		log.Println(err)
		return
	}

	yamlBody, err := yaml.Marshal(output)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("")
	fmt.Println(string(yamlBody))
	fmt.Println("")
	return
}
