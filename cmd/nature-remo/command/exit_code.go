package command

const (
	ExitCodeOK = iota
	ExitCodeInitializeConfigError
	ExitCodeParseArgsError
	ExitCodeRunCommandError
)
