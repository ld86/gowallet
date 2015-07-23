package cli

type Cli struct {
	commands map[string]func([]string) error
}

func New() *Cli {
	return &Cli{
		commands: map[string]func([]string) error{
			"g":   GetPassword,
			"get": GetPassword,
		},
	}
}

func (cli *Cli) Run(args []string) error {
	if len(args) < 1 {
		return nil
	}
	function := cli.commands[args[0]]
	return function(args[1:])
}
