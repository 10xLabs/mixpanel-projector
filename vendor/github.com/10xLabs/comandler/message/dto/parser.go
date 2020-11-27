package dto

import (
	"encoding/json"

	"github.com/10xLabs/chandler/awsevent"
)

// Parse interprets a []byte and returns the corresponding Message value.
// If m is empty or contains invalid json representation or could not be parsed
// as Commmand or Events ErrInvalidMessage is returned.
func Parse(m json.RawMessage) ([]DTO, error) {

	events, err := ParseEvents(m)
	if err == nil {
		dtos := []DTO{}
		for _, e := range events {
			dtos = append(dtos, e)
		}

		return dtos, nil
	}

	command, err := ParseCommand(m)
	if err == nil {
		return []DTO{command}, nil
	}

	commands, err := ParseCommands(m)

	if err == nil && len(commands) > 0 {
		dtos := []DTO{}
		for _, e := range commands {
			dtos = append(dtos, e)
		}
		return dtos, nil
	}

	return nil, ErrInvalidMessage
}

// ParseCommand interprets a []byte and returns the corresponding Command value.
// If m contains invalid json representation, nil and ErrSyntax is
// returned.
// ErrEmptyCommandType is returned when there's no command type.
func ParseCommand(m json.RawMessage) (*Command, error) {
	body := struct {
		Command Command `json:"command"`
	}{}
	if err := json.Unmarshal(m, &body); err != nil {
		return nil, ErrSintax
	}
	if len(body.Command.Type) == 0 {
		return nil, ErrEmptyCommandType
	}

	return &body.Command, nil
}

// ParseCommands ...
func ParseCommands(m json.RawMessage) ([]*Command, error) {
	body := struct {
		Commands []*Command `json:"commands"`
	}{}

	if err := json.Unmarshal(m, &body); err != nil {
		return nil, ErrSintax
	}

	return body.Commands, nil
}

// ParseEvents interprets a []byte and returns the corresponding Event slice value.
// If m contains invalid json representation, nil and ErrSyntax is
// returned.
// ErrEmptyKinesisRecords is returned when there's no kinesis records.
func ParseEvents(m json.RawMessage) ([]*Event, error) {
	var ge awsevent.Event
	if err := json.Unmarshal(m, &ge); err != nil {
		return nil, ErrSintax
	}

	if len(ge.Records) == 0 {
		return nil, ErrEmptyKinesisRecords
	}

	events := []*Event{}
	for _, r := range ge.Records {
		var e Event
		if err := json.Unmarshal(r.Data(), &e); err != nil {
			return nil, ErrSintax
		}

		events = append(events, &e)
	}

	return events, nil
}
