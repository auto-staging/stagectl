package helper

import (
	"encoding/json"
	"log"

	"github.com/auto-staging/stagectl/types"
)

// PrintAPIError unmarshals the Tower API error body and prints the error message to console
func PrintAPIError(body []byte) {
	apiError := types.APIErrorResponse{}
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
