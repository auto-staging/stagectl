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

func GetAllStatus() ([]types.EnvironmentStatus, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/environments/status", nil)
	if err != nil {
		return []types.EnvironmentStatus{}, err
	}

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []types.EnvironmentStatus{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []types.EnvironmentStatus{}, err
		}
		return []types.EnvironmentStatus{}, errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	status := []types.EnvironmentStatus{}
	err = json.Unmarshal([]byte(body), &status)
	if err != nil {
		return []types.EnvironmentStatus{}, err
	}

	return status, nil
}

func GetSingleStatus(repo, branch string) (types.EnvironmentStatus, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/"+repo+"/environments/"+branch+"/status", nil)
	if err != nil {
		return types.EnvironmentStatus{}, err
	}

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return types.EnvironmentStatus{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return types.EnvironmentStatus{}, err
		}
		return types.EnvironmentStatus{}, errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	status := types.EnvironmentStatus{}
	err = json.Unmarshal([]byte(body), &status)
	if err != nil {
		return types.EnvironmentStatus{}, err
	}

	return status, nil
}
