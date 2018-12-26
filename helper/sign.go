package helper

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

// SignRequest signs the Tower API request with AWS Signature v4 by using the local AWS IAM access key and secret access key.
func SignRequest(req *http.Request) {
	signer := v4.NewSigner(credentials.NewSharedCredentials("", "default"))

	// Sign without body
	if req.Body == nil {
		signer.Sign(req, nil, "execute-api", "eu-central-1", time.Now())
		return
	}

	// Convert request body to io.ReadSeeker
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Signing error")
		log.Println(err)
	}
	signer.Sign(req, bytes.NewReader(body), "execute-api", "eu-central-1", time.Now())
}
