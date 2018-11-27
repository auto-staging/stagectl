package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
	"gitlab.com/auto-staging/stagectl/helper"
	"gitlab.com/auto-staging/tower/types"
)

// GetEnvironmentsForRepo returns an array of Environment structs and an error for all Environments of the repository
func GetEnvironmentsForRepo(repo string) ([]types.Environment, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/"+repo+"/environments", nil)
	if err != nil {
		return []types.Environment{}, err
	}

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []types.Environment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []types.Environment{}, err
		}
		return []types.Environment{}, errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	envs := []types.Environment{}
	err = json.Unmarshal([]byte(body), &envs)
	if err != nil {
		return []types.Environment{}, err
	}

	return envs, nil
}

// GetSingleEnvironmentForRepo returns an Environment struct and error for a single Environment based on repo and branch
func GetSingleEnvironmentForRepo(repo, branch string) (types.Environment, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/"+repo+"/environments/"+branch, nil)
	if err != nil {
		return types.Environment{}, err
	}

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return types.Environment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return types.Environment{}, err
		}
		return types.Environment{}, errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	env := types.Environment{}
	err = json.Unmarshal([]byte(body), &env)
	if err != nil {
		return types.Environment{}, err
	}

	return env, nil
}
