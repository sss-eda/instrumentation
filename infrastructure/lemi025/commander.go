package lemi025

import (
	"fmt"

	"github.com/sss-eda/instrumentation/application/commanding"
)

// Commander TODO
type Commander struct {
	connection Connection
}

// NewCommander TODO
func NewCommander(
	connection Connection,
) (*Commander, error) {
	return &Commander{
		connection: connection,
	}, nil
}

// SendCommand TODO
func (commander *Commander) SendCommand(
	command commanding.Command,
) error {
	var err error

	switch command.Type() {
	case ReadConfig:
		err = commander.connection.ReadConfig()
	case ReadTime:
		err = commander.connection.ReadTime()
	case SetTime:
		err = commander.connection.SetTime()
	default:
		err = fmt.Errorf("Invalid command type: %q", command.Type().String())
	}

	return err
}
