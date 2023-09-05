package cli

import "github.com/hellflame/argparse"

type Opts struct {
	Args []string
}

func GetOpts() (*Opts, error) {
	parser := argparse.NewParser("globalparser", "Global parser to get all args", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	args := parser.Strings("a", "args", &argparse.Option{
		Positional: true,
		Required:   false,
		Default:    "",
	})

	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &Opts{
		Args: *args,
	}, nil
}
