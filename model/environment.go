package model

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
	"gitlab.com/auto-staging/tower/types"
)

// GetEnvironmentsForRepo calls the Tower API - GET /repositories/{name}/environments.
// If an error occurs the error gets returned, otherwise an array of Environment structs get returned.
func GetEnvironmentsForRepo(repo string) ([]types.Environment, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/"+repo+"/environments", nil)
	if err != nil {
		return []types.Environment{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return []types.Environment{}, err
	}

	envs := []types.Environment{}
	err = json.Unmarshal(result, &envs)
	if err != nil {
		return []types.Environment{}, err
	}

	return envs, nil
}

// GetSingleEnvironmentForRepo calls the Tower API - GET /repositories/{name}/environments/{branch}.
// If an error occurs the error gets returned, otherwise an Environment struct gets returned.
func GetSingleEnvironmentForRepo(repo, branch string) (types.Environment, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/"+repo+"/environments/"+branch, nil)
	if err != nil {
		return types.Environment{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return types.Environment{}, err
	}

	env := types.Environment{}
	err = json.Unmarshal(result, &env)
	if err != nil {
		return types.Environment{}, err
	}

	return env, nil
}

// UpdateSingleEnvironment calls the Tower API - PUT /repositories/{name}/environments/{branch}.
// If an error occurs the error gets returned, otherwise an Environment struct gets returned.
func UpdateSingleEnvironment(repo, branch string, body []byte) (types.Environment, error) {
	req, err := http.NewRequest("PUT", viper.GetString("tower_base_url")+"/repositories/"+repo+"/environments/"+branch, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return types.Environment{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return types.Environment{}, err
	}

	env := types.Environment{}
	err = json.Unmarshal(result, &env)
	if err != nil {
		return types.Environment{}, err
	}

	return env, nil
}

// DeleteSingleEnvironment calls the Tower API - DELETE /repositories/{name}/environments/{branch}.
// If an error occurs the error gets returned.
func DeleteSingleEnvironment(repo, branch string) error {
	req, err := http.NewRequest("DELETE", viper.GetString("tower_base_url")+"/repositories/"+repo+"/environments/"+branch, nil)
	if err != nil {
		return err
	}

	_, err = sendRequest(req, 202)
	if err != nil {
		return err
	}

	return nil
}

// AddEnvironment calls the Tower API - POST /repositories/{name}/environments/{branch}.
// If an error occurs the error gets returned, otherwise an Environment struct gets returned.
func AddEnvironment(repo string, body []byte) (types.Environment, error) {
	req, err := http.NewRequest("POST", viper.GetString("tower_base_url")+"/repositories/"+repo+"/environments/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return types.Environment{}, err
	}

	result, err := sendRequest(req, 201)
	if err != nil {
		return types.Environment{}, err
	}

	env := types.Environment{}
	err = json.Unmarshal(result, &env)
	if err != nil {
		return types.Environment{}, err
	}

	return env, nil
}
