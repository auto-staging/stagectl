package model

import (
	"io/ioutil"
	"net/http"

	"github.com/auto-staging/stagectl/helper"
)

func sendRequest(req *http.Request, expectedCode int) ([]byte, error) {
	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedCode {
		body, err := ioutil.ReadAll(resp.Body)
		helper.PrintAPIError(body)
		return []byte{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return respBody, nil
}
