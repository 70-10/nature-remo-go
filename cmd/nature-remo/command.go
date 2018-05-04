package main

import (
	"github.com/70-10/nature-remo-go/cmd/nature-remo/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"config": func() (cli.Command, error) {
			return &command.ConfigCommand{
				Meta: *meta,
			}, nil
		},
		"devices": func() (cli.Command, error) {
			return &command.DevicesCommand{
				Meta: *meta,
			}, nil
		},
	}
}
