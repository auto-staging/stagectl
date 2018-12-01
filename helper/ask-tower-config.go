package helper

import (
	"log"
	"os"
	"strconv"

	input "github.com/tcnksm/go-input"
	"gitlab.com/auto-staging/tower/types"
)

func AskForTowerConfigUpdateInput(config *types.TowerConfiguration) {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	query := "LogLevel"
	level, err := ui.Ask(query, &input.Options{
		Default: strconv.Itoa(config.LogLevel),
	})
	if err != nil {
		log.Fatal(err)
	}
	levelValue, err := strconv.Atoi(level)
	if err != nil {
		log.Fatal(err)
	}
	config.LogLevel = levelValue
}
