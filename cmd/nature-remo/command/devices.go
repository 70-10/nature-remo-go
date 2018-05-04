package command

import (
	"fmt"

	"github.com/70-10/nature-remo-go"
)

type DevicesCommand struct {
	Meta
}

func (c *DevicesCommand) Run(args []string) int {
	err := c.Config.Initialize()
	if err != nil {
		return ExitCodeInitializeConfigError
	}

	client := natureremo.NewClient(c.Config.Token)
	devices, err := client.GetDevices()
	if err != nil {
		return ExitCodeRunCommandError
	}

	for _, d := range devices {
		fmt.Println(d.CreatedAt.Format("2006-01-02 15:04:05") + " " + d.ID + " " + d.Name + " " + d.FirmwareVersion)
	}

	return ExitCodeOK
}

func (c *DevicesCommand) Synopsis() string {
	return "Device list"
}

func (c *DevicesCommand) Help() string {
	return ""
}
