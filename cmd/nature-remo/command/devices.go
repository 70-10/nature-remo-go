package command

import (
	"fmt"

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

	for i, d := range devices {
		if i != 0 {
			c.Ui.Output("--------------------------------------------------------")
		}

		c.Ui.Output(fmt.Sprint("ID               : ", d.ID))
		c.Ui.Output(fmt.Sprint("Name             : ", d.Name))
		c.Ui.Output(fmt.Sprint("Temperature      : ", d.NewestEvents.Temperature.Value))
		c.Ui.Output(fmt.Sprint("Humidity         : ", d.NewestEvents.Humidity.Value))
		c.Ui.Output(fmt.Sprint("Illumination     : ", d.NewestEvents.Illumination.Value))
		c.Ui.Output(fmt.Sprint("Firmware Version : ", d.FirmwareVersion))
		c.Ui.Output(fmt.Sprint("Created At       : ", d.CreatedAt.Format("2006-01-02 15:04:05")))
	}

	return ExitCodeOK
}

func (c *DevicesCommand) Synopsis() string {
	return "Device list"
}

func (c *DevicesCommand) Help() string {
	return ""
}
