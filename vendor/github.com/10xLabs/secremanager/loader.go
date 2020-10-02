package secremanager

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

// Loader ...
type Loader interface {
	// Load parses the JSON-encoded data from AWS Secrets Manager and stores the result
	// in the value pointed to by v. If v is nil or not a pointer,
	// Load returns error.
	Load(secretID string, v interface{}) error
}

type loader struct {
	client secretsmanageriface.SecretsManagerAPI
}

// NewLoader ...
func NewLoader(client secretsmanageriface.SecretsManagerAPI) Loader {
	return &loader{client: client}
}

func (l *loader) getSecretValue(secretID string) (*secretsmanager.GetSecretValueOutput, error) {
	in := &secretsmanager.GetSecretValueInput{SecretId: aws.String(secretID)}

	return l.client.GetSecretValue(in)
}

func (l *loader) Load(secretID string, v interface{}) error {
	res, err := l.getSecretValue(secretID)
	if err != nil {
		return err
	}

	if res.SecretString == nil {
		return NilSecretStringError{}
	}

	return json.Unmarshal([]byte(*res.SecretString), &v)
}
