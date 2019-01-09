package model

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/auto-staging/stagectl/helper"
)

func sendRequest(req *http.Request, expectedCode int) ([]byte, error) {
	helper.SignRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedCode {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return []byte{}, err
		}
		helper.PrintAPIError(body)
		log.Println(err)
		return []byte{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	return respBody, nil
}
