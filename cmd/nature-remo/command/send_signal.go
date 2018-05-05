package command

import (
	"flag"

	natureremo "github.com/70-10/nature-remo-go"
)

type SendSignalsCommand struct {
	Meta
}

func (c *SendSignalsCommand) Run(args []string) int {
	var id string
	flags := flag.NewFlagSet("send_signal", flag.ExitOnError)
	flags.Usage = func() { c.Ui.Error(c.Help()) }
	flags.StringVar(&id, "id", "", "id")
	if err := flags.Parse(args); err != nil {
		return ExitCodeParseArgsError
	}
	args = flag.Args()

	if id == "" {
		c.Ui.Error("id is requried")
		return ExitCodeRunCommandError
	}

	client := natureremo.NewClient(c.Config.Token)
	err := client.SendSignal(id)

	if err != nil {
		return ExitCodeRunCommandError
	}

	c.Ui.Output("Send " + id)

	return ExitCodeOK
}

func (c *SendSignalsCommand) Synopsis() string {
	return "Send signal"
}

func (c *SendSignalsCommand) Help() string {
	return "Help"
}
