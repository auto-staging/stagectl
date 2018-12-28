package model

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
	"gitlab.com/auto-staging/tower/types"
)

// GetAllRepositories calls the Tower API - GET /repositories.
// If an error occurs the error gets returned, otherwise an array of Repository structs gets returned.
func GetAllRepositories() ([]types.Repository, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories", nil)
	if err != nil {
		return []types.Repository{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return []types.Repository{}, err
	}

	var repos []types.Repository
	err = json.Unmarshal(result, &repos)
	if err != nil {
		return []types.Repository{}, err
	}

	return repos, nil
}

// GetSingleRepository calls the Tower API - GET /repositories/{name}.
// If an error occurs the error gets returned, otherwise a Repository struct gets returned.
func GetSingleRepository(repoName string) (types.Repository, error) {
	req, err := http.NewRequest("GET", viper.GetString("tower_base_url")+"/repositories/"+repoName, nil)
	if err != nil {
		return types.Repository{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return types.Repository{}, err
	}

	repo := types.Repository{}
	err = json.Unmarshal(result, &repo)
	if err != nil {
		return types.Repository{}, err
	}

	return repo, nil
}

// AddRepository calls the Tower API - POST /repositories.
// If an error occurs the error gets returned, otherwise a Repository struct gets returned.
func AddRepository(body []byte) (types.Repository, error) {
	req, err := http.NewRequest("POST", viper.GetString("tower_base_url")+"/repositories", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return types.Repository{}, err
	}

	result, err := sendRequest(req, 201)
	if err != nil {
		return types.Repository{}, err
	}

	repo := types.Repository{}
	err = json.Unmarshal(result, &repo)
	if err != nil {
		return types.Repository{}, err
	}

	return repo, nil
}

// DeleteRepository calls the Tower API - DELETE /repositories/{name}.
// If an error occurs the error gets returned, otherwise a Repository struct gets returned.
func DeleteRepository(repoName string) error {
	req, err := http.NewRequest("DELETE", viper.GetString("tower_base_url")+"/repositories/"+repoName, nil)
	if err != nil {
		return err
	}

	_, err = sendRequest(req, 204)

	return err
}

// UpdateRepository calls the Tower API - PUT /repositories/{name}.
// If an error occurs the error gets returned, otherwise a Repository struct gets returned.
func UpdateRepository(body []byte, repoName string) (types.Repository, error) {
	req, err := http.NewRequest("PUT", viper.GetString("tower_base_url")+"/repositories/"+repoName, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return types.Repository{}, err
	}

	result, err := sendRequest(req, 200)
	if err != nil {
		return types.Repository{}, err
	}

	repo := types.Repository{}
	err = json.Unmarshal(result, &repo)
	if err != nil {
		return types.Repository{}, err
	}

	return repo, nil
}
