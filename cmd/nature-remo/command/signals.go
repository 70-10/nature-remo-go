package command

import (
	"flag"

	natureremo "github.com/70-10/nature-remo-go"
)

type SignalsCommand struct {
	Meta
}

func (c *SignalsCommand) Run(args []string) int {
	var id string
	flags := flag.NewFlagSet("signals", flag.ExitOnError)
	flags.Usage = func() { c.Ui.Error(c.Help()) }
	flags.StringVar(&id, "id", "", "id")
	if err := flags.Parse(args); err != nil {
		return ExitCodeParseArgsError
	}
	args = flag.Args()

	if id == "" {
		c.Ui.Error("ID is requried")
		return ExitCodeRunCommandError
	}

	client := natureremo.NewClient(c.Config.Token)
	signals, err := client.GetSignals(id)
	if err != nil {
		return ExitCodeRunCommandError
	}

	for _, s := range signals {
		c.Ui.Output(s.ID + " " + s.Name)
	}

	return ExitCodeOK
}

func (c *SignalsCommand) Synopsis() string {
	return "Signal list"
}

func (c *SignalsCommand) Help() string {
	return "Help"
}
