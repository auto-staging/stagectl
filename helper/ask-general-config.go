package helper

import (
	"fmt"
	"log"
	"os"
	"strconv"

	input "github.com/tcnksm/go-input"
	"gitlab.com/auto-staging/tower/types"
)

func AskForGeneralConfigUpdateInput(config *types.GeneralConfig) {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	// SchutdownSchedules
	for choice := -1; choice != 0; {
		fmt.Printf("\nSchutdownSchedules: \n")
		for i, v := range config.ShutdownSchedules {
			fmt.Printf("SchutdownSchedule %d \n", i)
			fmt.Printf("	Cron %s \n \n", v.Cron)
		}
		fmt.Println("")

		// Reset choice
		choice = -1
		for choice > 3 || choice < 0 {
			query := "0 = End ShutdownSchedules editing \n1 = Edit single element \n2 = Delete single element \n3 = Add single element"
			cliValue, err := ui.Ask(query, &input.Options{
				Default: "0",
			})
			if err != nil {
				log.Fatal(err)
			}
			value, err := strconv.Atoi(cliValue)
			if err != nil {
				log.Fatal(err)
			}
			choice = value
		}

		switch choice {
		case 1:
			query := "Index of the SchutdownSchedule to edit"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}

			if index < len(config.ShutdownSchedules) && index >= 0 {
				query = "ShutdownSchedule (Cron)"
				name, err := ui.Ask(query, &input.Options{
					Default: config.ShutdownSchedules[index].Cron,
				})
				if err != nil {
					log.Fatal(err)
				}
				config.ShutdownSchedules[index].Cron = name
			} else {
				log.Println("Index out of range")
			}

		case 2:
			query := "Index of the SchutdownSchedule to remove"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}
			if index < len(config.ShutdownSchedules) && index >= 0 {
				config.ShutdownSchedules = append(config.ShutdownSchedules[:index], config.ShutdownSchedules[index+1:]...)
			} else {
				log.Println("Index out of range")
			}

		case 3:
			fmt.Println("Information for new SchutdownSchedule")
			query := "ShutdownSchedule (Cron)"
			name, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			schedule := types.TimeSchedule{
				Cron: name,
			}
			config.ShutdownSchedules = append(config.ShutdownSchedules, schedule)
		}
	}
	// StartupSchedules
	for choice := -1; choice != 0; {
		fmt.Printf("\nStartupSchedules: \n")
		for i, v := range config.StartupSchedules {
			fmt.Printf("StartupSchedules %d \n", i)
			fmt.Printf("	Cron %s \n \n", v.Cron)
		}
		fmt.Println("")

		// Reset choice
		choice = -1
		for choice > 3 || choice < 0 {
			query := "0 = End StartupSchedules editing \n1 = Edit single element \n2 = Delete single element \n3 = Add single element"
			cliValue, err := ui.Ask(query, &input.Options{
				Default: "0",
			})
			if err != nil {
				log.Fatal(err)
			}
			value, err := strconv.Atoi(cliValue)
			if err != nil {
				log.Fatal(err)
			}
			choice = value
		}

		switch choice {
		case 1:
			query := "Index of the StartupSchedules to edit"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}

			if index < len(config.StartupSchedules) && index >= 0 {
				query = "StartupSchedules (Cron)"
				name, err := ui.Ask(query, &input.Options{
					Default: config.StartupSchedules[index].Cron,
				})
				if err != nil {
					log.Fatal(err)
				}
				config.StartupSchedules[index].Cron = name
			} else {
				log.Println("Index out of range")
			}

		case 2:
			query := "Index of the StartupSchedules to remove"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}
			if index < len(config.StartupSchedules) && index >= 0 {
				config.StartupSchedules = append(config.StartupSchedules[:index], config.StartupSchedules[index+1:]...)
			} else {
				log.Println("Index out of range")
			}

		case 3:
			fmt.Println("Information for new StartupSchedules")
			query := "StartupSchedules (Cron)"
			name, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			schedule := types.TimeSchedule{
				Cron: name,
			}
			config.StartupSchedules = append(config.StartupSchedules, schedule)
		}
	}

	// EnvironmentVariables
	for choice := -1; choice != 0; {
		fmt.Printf("\nEnvironmentVariables: \n")
		for i, v := range config.EnvironmentVariables {
			fmt.Printf("EnvironmentVariables %d \n", i)
			fmt.Printf("	Name %s | Type %s | Value %s \n \n", v.Name, v.Type, v.Value)
		}
		fmt.Println("")

		// Reset choice
		choice = -1
		for choice > 3 || choice < 0 {
			query := "0 = End EnvironmentVariables editing \n1 = Edit single element \n2 = Delete single element \n3 = Add single element"
			cliValue, err := ui.Ask(query, &input.Options{
				Default: "0",
			})
			if err != nil {
				log.Fatal(err)
			}
			value, err := strconv.Atoi(cliValue)
			if err != nil {
				log.Fatal(err)
			}
			choice = value
		}

		switch choice {
		case 1:
			query := "Index of the EnvironmentVariables to edit"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}

			if index < len(config.EnvironmentVariables) && index >= 0 {
				query = "EnvironmentVariables (Name)"
				name, err := ui.Ask(query, &input.Options{
					Default: config.EnvironmentVariables[index].Name,
				})
				if err != nil {
					log.Fatal(err)
				}
				config.EnvironmentVariables[index].Name = name

				query = "EnvironmentVariables (Type)"
				name, err = ui.Ask(query, &input.Options{
					Default: config.EnvironmentVariables[index].Type,
				})
				if err != nil {
					log.Fatal(err)
				}
				config.EnvironmentVariables[index].Type = name

				query = "EnvironmentVariables (Value)"
				name, err = ui.Ask(query, &input.Options{
					Default: config.EnvironmentVariables[index].Value,
				})
				if err != nil {
					log.Fatal(err)
				}
				config.EnvironmentVariables[index].Value = name
			} else {
				log.Println("Index out of range")
			}

		case 2:
			query := "Index of the EnvironmentVariables to remove"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}
			if index < len(config.EnvironmentVariables) && index >= 0 {
				config.EnvironmentVariables = append(config.EnvironmentVariables[:index], config.EnvironmentVariables[index+1:]...)
			} else {
				log.Println("Index out of range")
			}

		case 3:
			fmt.Println("Information for new EnvironmentVariables")
			query := "EnvironmentVariables (Name)"
			name, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}

			query = "EnvironmentVariables (Type)"
			envType, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}

			query = "EnvironmentVariables (Value)"
			value, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			envVar := types.EnvironmentVariable{
				Name:  name,
				Type:  envType,
				Value: value,
			}
			config.EnvironmentVariables = append(config.EnvironmentVariables, envVar)
		}
	}
}
