package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
	"gitlab.com/auto-staging/stagectl/helper"
	"gitlab.com/auto-staging/tower/types"
)

func GetGeneralConfig() (types.GeneralConfig, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/environments", nil)
	if err != nil {
		return types.GeneralConfig{}, err
	}

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return types.GeneralConfig{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return types.GeneralConfig{}, err
		}
		helper.PrintAPIError(body)
		return types.GeneralConfig{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	config := types.GeneralConfig{}
	err = json.Unmarshal([]byte(body), &config)
	if err != nil {
		return types.GeneralConfig{}, err
	}

	return config, nil
}

func UpdateGeneralConfiguration(body []byte) (types.GeneralConfig, error) {
	req, err := http.NewRequest("PUT", viper.GetString("tower_base_url")+"/repositories/environments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return types.GeneralConfig{}, err
	}

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return types.GeneralConfig{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return types.GeneralConfig{}, err
		}
		helper.PrintAPIError(body)
		return types.GeneralConfig{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	config := types.GeneralConfig{}
	err = json.Unmarshal([]byte(respBody), &config)
	if err != nil {
		return types.GeneralConfig{}, err
	}

	return config, nil
}
