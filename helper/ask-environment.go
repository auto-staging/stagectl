package helper

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/auto-staging/tower/types"
	input "github.com/tcnksm/go-input"
)

// AskForEnvironmentUpdateInput sets the current Environment configuration as default and asks the user for new (updated) values.
// The updated values will be written back to the EnvironmentPut struct (call by reference)
func AskForEnvironmentUpdateInput(env *types.EnvironmentPut) {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	query := "CodeBuildRoleARN"
	name, err := ui.Ask(query, &input.Options{
		Default: env.CodeBuildRoleARN,
	})
	if err != nil {
		log.Fatal(err)
	}
	env.CodeBuildRoleARN = name

	query = "InfrastructureRepoURL"
	name, err = ui.Ask(query, &input.Options{
		Default: env.InfrastructureRepoURL,
	})
	if err != nil {
		log.Fatal(err)
	}
	env.InfrastructureRepoURL = name

	// SchutdownSchedules
	for choice := -1; choice != 0; {
		fmt.Printf("\nSchutdownSchedules: \n")
		for i, v := range env.ShutdownSchedules {
			fmt.Printf("SchutdownSchedule %d \n", i)
			fmt.Printf("	Cron %s \n \n", v.Cron)
		}
		fmt.Println("")

		// Reset choice
		choice = -1
		for choice > 3 || choice < 0 {
			query = "0 = End ShutdownSchedules editing \n1 = Edit single element \n2 = Delete single element \n3 = Add single element"
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
			query = "Index of the SchutdownSchedule to edit"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}

			if index < len(env.ShutdownSchedules) && index >= 0 {
				query = "ShutdownSchedule (Cron)"
				name, err = ui.Ask(query, &input.Options{
					Default: env.ShutdownSchedules[index].Cron,
				})
				if err != nil {
					log.Fatal(err)
				}
				env.ShutdownSchedules[index].Cron = name
			} else {
				log.Println("Index out of range")
			}

		case 2:
			query = "Index of the SchutdownSchedule to remove"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}
			if index < len(env.ShutdownSchedules) && index >= 0 {
				env.ShutdownSchedules = append(env.ShutdownSchedules[:index], env.ShutdownSchedules[index+1:]...)
			} else {
				log.Println("Index out of range")
			}

		case 3:
			fmt.Println("Information for new SchutdownSchedule")
			query = "ShutdownSchedule (Cron)"
			name, err = ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			schedule := types.TimeSchedule{
				Cron: name,
			}
			env.ShutdownSchedules = append(env.ShutdownSchedules, schedule)
		}
	}
	// StartupSchedules
	for choice := -1; choice != 0; {
		fmt.Printf("\nStartupSchedules: \n")
		for i, v := range env.StartupSchedules {
			fmt.Printf("StartupSchedules %d \n", i)
			fmt.Printf("	Cron %s \n \n", v.Cron)
		}
		fmt.Println("")

		// Reset choice
		choice = -1
		for choice > 3 || choice < 0 {
			query = "0 = End StartupSchedules editing \n1 = Edit single element \n2 = Delete single element \n3 = Add single element"
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
			query = "Index of the StartupSchedules to edit"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}

			if index < len(env.StartupSchedules) && index >= 0 {
				query = "StartupSchedules (Cron)"
				name, err = ui.Ask(query, &input.Options{
					Default: env.StartupSchedules[index].Cron,
				})
				if err != nil {
					log.Fatal(err)
				}
				env.StartupSchedules[index].Cron = name
			} else {
				log.Println("Index out of range")
			}

		case 2:
			query = "Index of the StartupSchedules to remove"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}
			if index < len(env.StartupSchedules) && index >= 0 {
				env.StartupSchedules = append(env.StartupSchedules[:index], env.StartupSchedules[index+1:]...)
			} else {
				log.Println("Index out of range")
			}

		case 3:
			fmt.Println("Information for new StartupSchedules")
			query = "StartupSchedules (Cron)"
			name, err = ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			schedule := types.TimeSchedule{
				Cron: name,
			}
			env.StartupSchedules = append(env.StartupSchedules, schedule)
		}
	}

	// EnvironmentVariables
	for choice := -1; choice != 0; {
		fmt.Printf("\nEnvironmentVariables: \n")
		for i, v := range env.EnvironmentVariables {
			fmt.Printf("EnvironmentVariables %d \n", i)
			fmt.Printf("	Name %s | Type %s | Value %s \n \n", v.Name, v.Type, v.Value)
		}
		fmt.Println("")

		// Reset choice
		choice = -1
		for choice > 3 || choice < 0 {
			query = "0 = End EnvironmentVariables editing \n1 = Edit single element \n2 = Delete single element \n3 = Add single element"
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
			query = "Index of the EnvironmentVariables to edit"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}

			if index < len(env.EnvironmentVariables) && index >= 0 {
				query = "EnvironmentVariables (Name)"
				name, err = ui.Ask(query, &input.Options{
					Default: env.EnvironmentVariables[index].Name,
				})
				if err != nil {
					log.Fatal(err)
				}
				env.EnvironmentVariables[index].Name = name

				query = "EnvironmentVariables (Type)"
				name, err = ui.Ask(query, &input.Options{
					Default: env.EnvironmentVariables[index].Type,
				})
				if err != nil {
					log.Fatal(err)
				}
				env.EnvironmentVariables[index].Type = name

				query = "EnvironmentVariables (Value)"
				name, err = ui.Ask(query, &input.Options{
					Default: env.EnvironmentVariables[index].Value,
				})
				if err != nil {
					log.Fatal(err)
				}
				env.EnvironmentVariables[index].Value = name
			} else {
				log.Println("Index out of range")
			}

		case 2:
			query = "Index of the EnvironmentVariables to remove"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}
			if index < len(env.EnvironmentVariables) && index >= 0 {
				env.EnvironmentVariables = append(env.EnvironmentVariables[:index], env.EnvironmentVariables[index+1:]...)
			} else {
				log.Println("Index out of range")
			}

		case 3:
			fmt.Println("Information for new EnvironmentVariables")
			query = "EnvironmentVariables (Name)"
			name, err = ui.Ask(query, &input.Options{})
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
			env.EnvironmentVariables = append(env.EnvironmentVariables, envVar)
		}
	}
}

// AskForEnvironmentAddInput asks the user for values used in the new Environment.
// The updated values will be written back to the EnvironmentPost struct (call by reference)
func AskForEnvironmentAddInput(env *types.EnvironmentPost) {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	query := "Branch"
	name, err := ui.Ask(query, &input.Options{})
	if err != nil {
		log.Fatal(err)
	}
	env.Branch = name

	query = "InfrastructureRepoURL"
	name, err = ui.Ask(query, &input.Options{})
	if err != nil {
		log.Fatal(err)
	}
	env.InfrastructureRepoURL = name

	query = "CodeBuildRoleARN"
	name, err = ui.Ask(query, &input.Options{})
	if err != nil {
		log.Fatal(err)
	}
	env.CodeBuildRoleARN = name

	// SchutdownSchedules
	for choice := -1; choice != 0; {
		fmt.Printf("\nSchutdownSchedules: \n")
		for i, v := range env.ShutdownSchedules {
			fmt.Printf("SchutdownSchedule %d \n", i)
			fmt.Printf("	Cron %s \n \n", v.Cron)
		}
		fmt.Println("")

		// Reset choice
		choice = -1
		for choice > 3 || choice < 0 {
			query = "0 = End ShutdownSchedules editing \n1 = Edit single element \n2 = Delete single element \n3 = Add single element"
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
			query = "Index of the SchutdownSchedule to edit"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}

			if index < len(env.ShutdownSchedules) && index >= 0 {
				query = "ShutdownSchedule (Cron)"
				name, err = ui.Ask(query, &input.Options{
					Default: env.ShutdownSchedules[index].Cron,
				})
				if err != nil {
					log.Fatal(err)
				}
				env.ShutdownSchedules[index].Cron = name
			} else {
				log.Println("Index out of range")
			}

		case 2:
			query = "Index of the SchutdownSchedule to remove"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}
			if index < len(env.ShutdownSchedules) && index >= 0 {
				env.ShutdownSchedules = append(env.ShutdownSchedules[:index], env.ShutdownSchedules[index+1:]...)
			} else {
				log.Println("Index out of range")
			}

		case 3:
			fmt.Println("Information for new SchutdownSchedule")
			query = "ShutdownSchedule (Cron)"
			name, err = ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			schedule := types.TimeSchedule{
				Cron: name,
			}
			env.ShutdownSchedules = append(env.ShutdownSchedules, schedule)
		}
	}
	// StartupSchedules
	for choice := -1; choice != 0; {
		fmt.Printf("\nStartupSchedules: \n")
		for i, v := range env.StartupSchedules {
			fmt.Printf("StartupSchedules %d \n", i)
			fmt.Printf("	Cron %s \n \n", v.Cron)
		}
		fmt.Println("")

		// Reset choice
		choice = -1
		for choice > 3 || choice < 0 {
			query = "0 = End StartupSchedules editing \n1 = Edit single element \n2 = Delete single element \n3 = Add single element"
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
			query = "Index of the StartupSchedules to edit"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}

			if index < len(env.StartupSchedules) && index >= 0 {
				query = "StartupSchedules (Cron)"
				name, err = ui.Ask(query, &input.Options{
					Default: env.StartupSchedules[index].Cron,
				})
				if err != nil {
					log.Fatal(err)
				}
				env.StartupSchedules[index].Cron = name
			} else {
				log.Println("Index out of range")
			}

		case 2:
			query = "Index of the StartupSchedules to remove"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}
			if index < len(env.StartupSchedules) && index >= 0 {
				env.StartupSchedules = append(env.StartupSchedules[:index], env.StartupSchedules[index+1:]...)
			} else {
				log.Println("Index out of range")
			}

		case 3:
			fmt.Println("Information for new StartupSchedules")
			query = "StartupSchedules (Cron)"
			name, err = ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			schedule := types.TimeSchedule{
				Cron: name,
			}
			env.StartupSchedules = append(env.StartupSchedules, schedule)
		}
	}

	// EnvironmentVariables
	for choice := -1; choice != 0; {
		fmt.Printf("\nEnvironmentVariables: \n")
		for i, v := range env.EnvironmentVariables {
			fmt.Printf("EnvironmentVariables %d \n", i)
			fmt.Printf("	Name %s | Type %s | Value %s \n \n", v.Name, v.Type, v.Value)
		}
		fmt.Println("")

		// Reset choice
		choice = -1
		for choice > 3 || choice < 0 {
			query = "0 = End EnvironmentVariables editing \n1 = Edit single element \n2 = Delete single element \n3 = Add single element"
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
			query = "Index of the EnvironmentVariables to edit"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}

			if index < len(env.EnvironmentVariables) && index >= 0 {
				query = "EnvironmentVariables (Name)"
				name, err = ui.Ask(query, &input.Options{
					Default: env.EnvironmentVariables[index].Name,
				})
				if err != nil {
					log.Fatal(err)
				}
				env.EnvironmentVariables[index].Name = name

				query = "EnvironmentVariables (Type)"
				name, err = ui.Ask(query, &input.Options{
					Default: env.EnvironmentVariables[index].Type,
				})
				if err != nil {
					log.Fatal(err)
				}
				env.EnvironmentVariables[index].Type = name

				query = "EnvironmentVariables (Value)"
				name, err = ui.Ask(query, &input.Options{
					Default: env.EnvironmentVariables[index].Value,
				})
				if err != nil {
					log.Fatal(err)
				}
				env.EnvironmentVariables[index].Value = name
			} else {
				log.Println("Index out of range")
			}

		case 2:
			query = "Index of the EnvironmentVariables to remove"
			indexStr, err := ui.Ask(query, &input.Options{})
			if err != nil {
				log.Fatal(err)
			}
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal(err)
			}
			if index < len(env.EnvironmentVariables) && index >= 0 {
				env.EnvironmentVariables = append(env.EnvironmentVariables[:index], env.EnvironmentVariables[index+1:]...)
			} else {
				log.Println("Index out of range")
			}

		case 3:
			fmt.Println("Information for new EnvironmentVariables")
			query = "EnvironmentVariables (Name)"
			name, err = ui.Ask(query, &input.Options{})
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
			env.EnvironmentVariables = append(env.EnvironmentVariables, envVar)
		}
	}
}
