package command

import (
	"fmt"
	"time"

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
		c.Ui.Output(fmt.Sprint("ID                : ", d.ID))
		c.Ui.Output(fmt.Sprint("Name              : ", d.Name))
		c.Ui.Output(fmt.Sprint("Temperature       : ", d.NewestEvents.Temperature.Value))
		c.Ui.Output(fmt.Sprint("Temperature Time  : ", format(d.NewestEvents.Temperature.CreatedAt)))
		c.Ui.Output(fmt.Sprint("Humidity          : ", d.NewestEvents.Humidity.Value))
		c.Ui.Output(fmt.Sprint("Humidity Time     : ", format(d.NewestEvents.Humidity.CreatedAt)))
		c.Ui.Output(fmt.Sprint("Illumination      : ", d.NewestEvents.Illumination.Value))
		c.Ui.Output(fmt.Sprint("Illumination Time : ", format(d.NewestEvents.Illumination.CreatedAt)))
		c.Ui.Output(fmt.Sprint("Firmware Version  : ", d.FirmwareVersion))
		c.Ui.Output(fmt.Sprint("Created At        : ", d.CreatedAt.In(time.Local).Format("2006-01-02 15:04:05")))
	}

	return ExitCodeOK
}

func (c *DevicesCommand) Synopsis() string {
	return "Device list"
}

func (c *DevicesCommand) Help() string {
	return ""
}

func format(t time.Time) string {
	return t.In(time.Local).Format("2006-01-02 15:04:05")
}
