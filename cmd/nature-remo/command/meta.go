package command

import (
	"github.com/70-10/nature-remo-go/cmd/nature-remo/config"
	"github.com/mitchellh/cli"
)

type Meta struct {
	Ui     cli.Ui
	Config config.Config
}
