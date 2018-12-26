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

// GetEnvironmentsForRepo calls the Tower API - GET /repositories/{name}/environments.
// If an error occurs the error gets returned, otherwise an array of Environment structs get returned.
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
		helper.PrintAPIError(body)
		return []types.Environment{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	envs := []types.Environment{}
	err = json.Unmarshal([]byte(body), &envs)
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
		helper.PrintAPIError(body)
		return types.Environment{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	env := types.Environment{}
	err = json.Unmarshal([]byte(body), &env)
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
		helper.PrintAPIError(body)
		return types.Environment{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	env := types.Environment{}
	err = json.Unmarshal([]byte(respBody), &env)
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

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		helper.PrintAPIError(body)
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

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return types.Environment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return types.Environment{}, err
		}
		helper.PrintAPIError(body)
		return types.Environment{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	env := types.Environment{}
	err = json.Unmarshal([]byte(respBody), &env)
	if err != nil {
		return types.Environment{}, err
	}

	return env, nil
}
