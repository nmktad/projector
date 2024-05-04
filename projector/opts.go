package projector

import (
	"github.com/hellflame/argparse"
)

type ProjectorOptions struct {
	Config string
	Pwd    string
	Args   []string
}

func GetProjectorOptions() (*ProjectorOptions, error) {
	parser := argparse.NewParser("Projector", "Projector is a tool for managing your projects", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	args := parser.Strings("a", "args", &argparse.Option{
		Positional: true,
		Default:    "",
		Help:       "Arguments for the command",
	})

	config := parser.String("c", "config", &argparse.Option{
		Default:  "",
		Required: false,
		Help:     "Path to the config file",
	})

	pwd := parser.String("p", "pwd", &argparse.Option{
		Default:  "",
		Required: false,
		Help:     "Path to the project directory",
	})

	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &ProjectorOptions{
		Args:   *args,
		Config: *config,
		Pwd:    *pwd,
	}, nil
}
