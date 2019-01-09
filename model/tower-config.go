package model

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/auto-staging/tower/types"
	"github.com/spf13/viper"
)

// GetTowerConfig calls the Tower API - GET /configuration.
// If an error occurs the error gets returned, otherwise a TowerConfiguration struct gets returned.
func GetTowerConfig() (types.TowerConfiguration, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/configuration", nil)
	if err != nil {
		return types.TowerConfiguration{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return types.TowerConfiguration{}, err
	}

	config := types.TowerConfiguration{}
	err = json.Unmarshal(result, &config)
	if err != nil {
		return types.TowerConfiguration{}, err
	}

	return config, nil
}

// UpdateTowerConfiguration calls the Tower API - PUT /configuration.
// If an error occurs the error gets returned, otherwise a TowerConfiguration struct gets returned.
func UpdateTowerConfiguration(body []byte) (types.TowerConfiguration, error) {
	req, err := http.NewRequest("PUT", viper.GetString("tower_base_url")+"/configuration", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return types.TowerConfiguration{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return types.TowerConfiguration{}, err
	}

	config := types.TowerConfiguration{}
	err = json.Unmarshal(result, &config)
	if err != nil {
		return types.TowerConfiguration{}, err
	}

	return config, nil
}
