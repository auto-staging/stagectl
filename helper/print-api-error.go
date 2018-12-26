package helper

import (
	"encoding/json"
	"log"

	"gitlab.com/auto-staging/stagectl/types"
)

// PrintApiError unmarshals the Tower API error body and prints the error message to console
func PrintApiError(body []byte) {
	apiError := types.ApiErrorResponse{}
	err := json.Unmarshal(body, &apiError)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("API Error")

	// Output raw json if message was not successfully unmarshaled
	if apiError.Message == "" {
		log.Fatal(string(body))
	}

	log.Fatal(apiError.Message)
}
