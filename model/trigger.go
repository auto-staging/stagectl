package model

import (
	"bytes"
	"net/http"

	"github.com/spf13/viper"
)

// TriggerSchedule calls the Tower API - POST /triggers/schedule.
// If an error occurs the error gets returned.
func TriggerSchedule(body []byte) error {
	req, err := http.NewRequest("POST", viper.GetString("tower_base_url")+"/triggers/schedule", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	_, err = sendRequest(req, 200)

	return err
}
