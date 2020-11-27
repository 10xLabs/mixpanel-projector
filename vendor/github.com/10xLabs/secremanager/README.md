# secremanager
Avoids "boilerplate" code when using AWS Secrets Manager


## Install

```
go get -u github.com/10xLabs/secremanager
```

## Usage

To get started, import the `secremanager` package, create a `secremanager.Loader`:

```go
import (
    "github.com/10xLabs/secremanager"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/secretsmanager"
)

var loader = secremanager.NewLoader(secretsmanager.New(session.New()))
```

To load AWS Secrets:

```go
var s map[string]string // you can use a struct as well

err := loader.Load(awsSecretsID, &s)
```
