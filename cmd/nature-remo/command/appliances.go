package command

import (
	"github.com/70-10/nature-remo-go"
)

type AppliancesCommand struct {
	Meta
}

func (c *AppliancesCommand) Run(args []string) int {
	client := natureremo.NewClient(c.Config.Token)
	appliances, err := client.GetAppliances()
	if err != nil {
		return ExitCodeRunCommandError
	}

	for _, appliance := range appliances {
		c.Ui.Output(appliance.ID + " " + appliance.Nickname)
	}

	return ExitCodeOK
}

func (c *AppliancesCommand) Synopsis() string {
	return "Appliance list"
}

func (c *AppliancesCommand) Help() string {
	return "Help"
}
