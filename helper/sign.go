package helper

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

func SignRequest(req *http.Request) {
	signer := v4.NewSigner(credentials.NewSharedCredentials("", "default"))
	signer.Sign(req, nil, "execute-api", "eu-central-1", time.Now())
}
