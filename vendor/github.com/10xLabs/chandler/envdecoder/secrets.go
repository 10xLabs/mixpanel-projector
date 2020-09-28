package envdecoder

import (
	"log"
	"os"

	"github.com/10xLabs/secremanager"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var secretsID = os.Getenv("AWS_SECRETS_ID")
var loader secremanager.Loader

func variablesFromSecrets() (s map[string]string) {
	loader := secremanager.NewLoader(secretsmanager.New(session.New()))

	if len(secretsID) > 0 {
		if err := loader.Load(secretsID, &s); err != nil {
			log.Println(err)
		}
	}
	return s
}
