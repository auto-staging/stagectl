package model

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
	"gitlab.com/auto-staging/tower/types"
)

// GetGeneralConfig calls the Tower API - GET /repositories/environments.
// If an error occurs the error gets returned, otherwise an GeneralConfig struct gets returned.
func GetGeneralConfig() (types.GeneralConfig, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/environments", nil)
	if err != nil {
		return types.GeneralConfig{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return types.GeneralConfig{}, err
	}

	config := types.GeneralConfig{}
	err = json.Unmarshal(result, &config)
	if err != nil {
		return types.GeneralConfig{}, err
	}

	return config, nil
}

// UpdateGeneralConfiguration calls the Tower API - PUT /repositories/environments.
// If an error occurs the error gets returned, otherwise an GeneralConfig struct gets returned.
func UpdateGeneralConfiguration(body []byte) (types.GeneralConfig, error) {
	req, err := http.NewRequest("PUT", viper.GetString("tower_base_url")+"/repositories/environments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return types.GeneralConfig{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return types.GeneralConfig{}, err
	}

	config := types.GeneralConfig{}
	err = json.Unmarshal(result, &config)
	if err != nil {
		return types.GeneralConfig{}, err
	}

	return config, nil
}
