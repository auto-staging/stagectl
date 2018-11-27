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
