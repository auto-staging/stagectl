package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
	"gitlab.com/auto-staging/stagectl/helper"
	"gitlab.com/auto-staging/tower/types"
)

func GetAllRepositories() ([]types.Repository, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories", nil)
	if err != nil {
		return []types.Repository{}, err
	}

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []types.Repository{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []types.Repository{}, err
		}
		return []types.Repository{}, errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	repos := []types.Repository{}
	err = json.Unmarshal([]byte(body), &repos)
	if err != nil {
		return []types.Repository{}, err
	}

	return repos, nil
}

func GetSingleRepository(repoName string) (types.Repository, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/"+repoName, nil)
	if err != nil {
		return types.Repository{}, err
	}

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return types.Repository{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return types.Repository{}, err
		}
		return types.Repository{}, errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	repo := types.Repository{}
	err = json.Unmarshal([]byte(body), &repo)
	if err != nil {
		return types.Repository{}, err
	}

	return repo, nil
}

func AddRepository(body []byte) (types.Repository, error) {
	req, err := http.NewRequest("POST", viper.GetString("tower_base_url")+"/repositories", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return types.Repository{}, err
	}

	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return types.Repository{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return types.Repository{}, err
		}
		return types.Repository{}, errors.New(string(body))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	repo := types.Repository{}
	err = json.Unmarshal([]byte(respBody), &repo)
	if err != nil {
		return types.Repository{}, err
	}

	return repo, nil
}
