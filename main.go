package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nmktad/projector/projector"
)

func main() {
	opts, err := projector.GetProjectorOptions()
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	config, err := projector.NewConfig(opts)
	if err != nil {
		log.Fatalf("Error creating config: %v", err)
	}

	proj := projector.NewProjector(config)

	switch config.Operation {
	case projector.Print:
		if len(config.Args) == 0 {
			data := proj.GetValueAll()
			jsonString, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("Error marshalling data: %v", err)
			}

			fmt.Printf("%s\n", string(jsonString))
		} else {
			if val, ok := proj.GetValue(config.Args[0]); ok {
				fmt.Printf("%s\n", val)
			}
		}
	case projector.Add:
		proj.SetValue(config.Args[0], config.Args[1])
		proj.Save()
	case projector.Remove:
		proj.RemoveValue(config.Args[0])
		proj.Save()
	default:
		log.Fatalf("Invalid operation: %v", config.Operation)
	}
}
