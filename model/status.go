package model

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
	"gitlab.com/auto-staging/tower/types"
)

// GetAllStatus calls the Tower API - GET /repositories/environments/status.
// If an error occurs the error gets returned, otherwise an array of EnvironmentStatus structs gets returned.
func GetAllStatus() ([]types.EnvironmentStatus, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/environments/status", nil)
	if err != nil {
		return []types.EnvironmentStatus{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return []types.EnvironmentStatus{}, err
	}

	status := []types.EnvironmentStatus{}
	err = json.Unmarshal(result, &status)
	if err != nil {
		return []types.EnvironmentStatus{}, err
	}

	return status, nil
}

// GetSingleStatus calls the Tower API - GET /repositories/{name}/environments/{branch}/status.
// If an error occurs the error gets returned, otherwise an EnvironmentStatus struct gets returned.
func GetSingleStatus(repo, branch string) (types.EnvironmentStatus, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/"+repo+"/environments/"+branch+"/status", nil)
	if err != nil {
		return types.EnvironmentStatus{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return types.EnvironmentStatus{}, err
	}

	status := types.EnvironmentStatus{}
	err = json.Unmarshal(result, &status)
	if err != nil {
		return types.EnvironmentStatus{}, err
	}

	return status, nil
}
