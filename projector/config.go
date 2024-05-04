package projector

import (
	"fmt"
	"log"
	"os"
	"path"
)

type Operation int

const (
	Print Operation = iota
	Add
	Remove
)

type Config struct {
	Pwd       string
	Config    string
	Args      []string
	Operation Operation
}

func getPwd(opts *ProjectorOptions) (string, error) {
	if opts.Pwd != "" {
		return opts.Pwd, nil
	}

	return os.Getwd()
}

func getConfig(opts *ProjectorOptions) (string, error) {
	if opts.Config != "" {
		return opts.Config, nil
	}

	configDir, err := os.UserConfigDir()
	if err == nil {
		return path.Join(configDir, "projector", "projector.json"), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to create config location: %v", err)
	}

	return path.Join(homeDir, ".projector.json"), nil
}

func getOperation(opts *ProjectorOptions) Operation {
	if len(opts.Args) == 0 {
		return Print
	}

	switch opts.Args[0] {
	case "add":
		return Add
	case "rm":
		return Remove
	default:
		return Print
	}
}

func getArgs(opts *ProjectorOptions) ([]string, error) {
	if len(opts.Args) == 0 {
		return []string{}, nil
	}

	switch getOperation(opts) {
	case Add:
		if len(opts.Args) != 3 {
			return nil, fmt.Errorf("add requires 2 arguments but got %d", len(opts.Args)-1)
		}
		return opts.Args[1:], nil
	case Remove:
		if len(opts.Args) != 2 {
			return nil, fmt.Errorf("rm requires 1 argument but got %d", len(opts.Args)-1)
		}
		return opts.Args[1:], nil
	case Print:
		if len(opts.Args) > 1 {
			return nil, fmt.Errorf("print requires no arguments but got %d", len(opts.Args)-1)
		}
		return opts.Args, nil
	default:
		return nil, fmt.Errorf("invalid operation options for operation: %s", opts.Args[0])
	}
}

func NewConfig(opts *ProjectorOptions) (*Config, error) {
	pwd, err := getPwd(opts)
	if err != nil {
		return nil, err
	}

	config, err := getConfig(opts)
	if err != nil {
		return nil, err
	}

	args, err := getArgs(opts)
	if err != nil {
		return nil, err
	}

	return &Config{
		Pwd:       pwd,
		Config:    config,
		Args:      args,
		Operation: getOperation(opts),
	}, nil
}
