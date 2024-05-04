package projector

import (
	"encoding/json"
	"os"
	"path"
	"slices"
)

type Data struct {
	Projector map[string]map[string]string `json:"projector"`
}

type Projector struct {
	data   *Data
	config *Config
}

func CreateProjector(config *Config, data *Data) *Projector {
	return &Projector{
		config: config,
		data:   data,
	}
}

func (p *Projector) GetValue(key string) (string, bool) {
	curr := p.config.Pwd
	found := false
	prev := ""

	out := ""

	for curr != prev {
		if val, ok := p.data.Projector[curr][key]; ok {
			out = val
			found = true
			break
		}

		prev = curr
		curr = path.Dir(curr)
	}

	return out, found
}

func (p *Projector) GetValueAll() map[string]string {
	curr := p.config.Pwd
	prev := ""

	out := make(map[string]string)
	paths := []string{}

	for curr != prev {
		paths = append(paths, curr)
		prev = curr
		curr = path.Dir(curr)
	}

	slices.Reverse(paths)
	for _, elem := range paths {
		if dir, ok := p.data.Projector[elem]; ok {
			for key, value := range dir {
				out[key] = value
			}
		}
	}

	return out
}

func (p *Projector) SetValue(key, value string) {
	pwd := p.config.Pwd

	if _, ok := p.data.Projector[pwd]; !ok {
		p.data.Projector[pwd] = make(map[string]string)
	}

	p.data.Projector[pwd][key] = value
}

func (p *Projector) Save() error {
	content, err := json.Marshal(p.data)
	if err != nil {
		return err
	}

	dir := path.Dir(p.config.Config)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0644); err != nil {
			return err
		}
	}

	if err := os.WriteFile(p.config.Config, content, 0644); err != nil {
		return err
	}

	return nil
}

func (p *Projector) RemoveValue(key string) {
	pwd := p.config.Pwd

	if _, ok := p.data.Projector[pwd]; ok {
		delete(p.data.Projector[pwd], key)
	}
}

func defaultProjector(config *Config) *Projector {
	return &Projector{
		data: &Data{
			Projector: make(map[string]map[string]string),
		},
		config: config,
	}
}

func NewProjector(config *Config) *Projector {
	if _, err := os.Stat(config.Config); err == nil {
		content, err := os.ReadFile(config.Config)
		if err != nil {
			return defaultProjector(config)
		}

		var data Data

		if err := json.Unmarshal(content, &data); err != nil {
			return defaultProjector(config)
		}

		return &Projector{
			data:   &data,
			config: config,
		}
	}

	return defaultProjector(config)
}
