package command

import (
	"github.com/70-10/nature-remo-go"
)

type DevicesCommand struct {
	Meta
}

func (c *DevicesCommand) Run(args []string) int {

	client := natureremo.NewClient(c.Config.Token)
	devices, err := client.GetDevices()
	if err != nil {
		return ExitCodeRunCommandError
	}

	for _, d := range devices {
		c.Ui.Output(d.CreatedAt.Format("2006-01-02 15:04:05") + " " + d.ID + " " + d.Name + " " + d.FirmwareVersion)
	}

	return ExitCodeOK
}

func (c *DevicesCommand) Synopsis() string {
	return "Device list"
}

func (c *DevicesCommand) Help() string {
	return ""
}
