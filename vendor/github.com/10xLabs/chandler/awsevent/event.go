package awsevent

import (
	"github.com/10xLabs/log"
	"github.com/aws/aws-lambda-go/events"
)

// Event ...
type Event struct {
	MustDeleteCache bool          `json:"mustDeleteCache"`
	IsReplay        bool          `json:"isReplay"`
	Records         []EventRecord `json:"Records"`
}

// EventRecord ...
type EventRecord struct {
	Kinesis           *events.KinesisRecord                  `json:"kinesis"`
	MessageAttributes *map[string]events.SQSMessageAttribute `json:"messageAttributes"`
}

// Data ...
func (g *EventRecord) Data() []byte {
	if g.Kinesis != nil {
		return g.Kinesis.Data
	}
	if g.MessageAttributes != nil {
		return (*g.MessageAttributes)["data"].BinaryValue
	}

	log.WithFields(log.Fields{
		"record": g,
	}).Info("Kinesis data or Message Attributes data not found")

	return nil
}
