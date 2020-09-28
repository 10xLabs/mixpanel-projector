//go:generate go run -mod=mod github.com/10xLabs/chandler/subgen
package main

import (
	"github.com/10xLabs/log"
	"github.com/10xLabs/mixpanel-projector/config"
	"github.com/10xLabs/mixpanel-projector/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	config.Setup()
	handler.Setup()
	log.Setup(config.App.Log)
}

// Event ...
type Event struct {
	Event      string          `json:"event"`
	Properties EventProperties `json:"properties"`
}

// EventProperties ...
type EventProperties struct {
	Todo
	BookingID string `json:"bookingID"`
	Token     string `json:"token"`
}

type Todo interface{}

func main() {
	lambda.Start(
		handler.Handler,
	)
}
