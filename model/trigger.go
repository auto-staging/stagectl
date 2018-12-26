package model

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
	"gitlab.com/auto-staging/stagectl/helper"
)

func TriggerSchedule(body []byte) error {
	req, err := http.NewRequest("POST", viper.GetString("tower_base_url")+"/triggers/schedule", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
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

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		helper.PrintAPIError(body)
		return err
	}

	return nil
}
