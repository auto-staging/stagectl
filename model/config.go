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

func GetTowerConfig() (types.TowerConfiguration, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/configuration", nil)
	if err != nil {
		return types.TowerConfiguration{}, err
	}

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return types.TowerConfiguration{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return types.TowerConfiguration{}, err
		}
		return types.TowerConfiguration{}, errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	config := types.TowerConfiguration{}
	err = json.Unmarshal([]byte(body), &config)
	if err != nil {
		return types.TowerConfiguration{}, err
	}

	return config, nil
}
