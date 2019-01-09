package helper

import (
	"log"
	"os"
	"strconv"

	"github.com/auto-staging/tower/types"
	input "github.com/tcnksm/go-input"
)

// AskForTowerConfigUpdateInput sets the current Tower configuration as default and asks the user for new (updated) values.
// The updated values will be written back to the TowerConfiguration struct (call by reference)
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
