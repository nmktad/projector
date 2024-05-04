package projector_test

import (
	"reflect"
	"testing"

	"github.com/nmktad/projector/projector"
)

func getOpts(args []string) *projector.ProjectorOptions {
	return &projector.ProjectorOptions{
		Args:   args,
		Config: "",
		Pwd:    "",
	}
}

func TestConfigPrint(t *testing.T) {
	opts := getOpts([]string{})
	config, err := projector.NewConfig(opts)
	if err != nil {
		t.Errorf("Failed to create config: %v", err)
	}

	if !reflect.DeepEqual([]string{}, config.Args) {
		t.Errorf("Expected empty args, got %v", config.Args)
	}

	if config.Operation != projector.Print {
		t.Errorf("Expected Print operation, got %v", config.Operation)
	}
}

func TestConfigPrintKey(t *testing.T) {
	opts := getOpts([]string{"key"})
	config, err := projector.NewConfig(opts)
	if err != nil {
		t.Errorf("Failed to create config: %v", err)
	}

	if !reflect.DeepEqual([]string{"key"}, config.Args) {
		t.Errorf("Expected args to be %v, got %v", "key", config.Args)
	}

	if config.Operation != projector.Print {
		t.Errorf("Expected Print operation, got %v", config.Operation)
	}
}

func TestConfigAddKeyValue(t *testing.T) {
	opts := getOpts([]string{"add", "key", "value"})

	config, err := projector.NewConfig(opts)
	if err != nil {
		t.Errorf("Failed to create config: %v", err)
	}

	if !reflect.DeepEqual([]string{"key", "value"}, config.Args) {
		t.Errorf("Expected args to be %v, got %v", []string{"key", "value"}, config.Args)
	}

	if config.Operation != projector.Add {
		t.Errorf("Expected Print operation, got %v", config.Operation)
	}
}

func TestRemove(t *testing.T) {
	opts := getOpts([]string{"rm", "key"})
	config, err := projector.NewConfig(opts)
	if err != nil {
		t.Errorf("Failed to create config: %v", err)
	}

	if !reflect.DeepEqual([]string{"key"}, config.Args) {
		t.Errorf("Expected args to be %v, got %v", []string{"key"}, config.Args)
	}

	if config.Operation != projector.Remove {
		t.Errorf("Expected Print operation, got %v", config.Operation)
	}
}
