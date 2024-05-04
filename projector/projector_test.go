package projector_test

import (
	"testing"

	"github.com/nmktad/projector/projector"
)

func getData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar",
				"baz": "qux",
			},
			"/foo": {
				"foo": "binx",
			},
			"/foo/bar": {
				"foo": "bing",
			},
		},
	}
}

func getProjector(pwd string, data projector.Data) *projector.Projector {
	return projector.CreateProjector(
		&projector.Config{
			Args:      []string{},
			Operation: projector.Print,
			Pwd:       pwd,
			Config:    "Hello there",
		},
		&data,
	)
}

func TestGetValue(t *testing.T) {
	data := getData()
	p := getProjector("/foo/bar", *data)

	val, ok := p.GetValue("foo")
	if val != "bing" || !ok {
		t.Errorf("Expected 'bing', got '%s'", val)
	}

	if val, ok := p.GetValue("baz"); !ok && val != "qux" {
		t.Errorf("Expected qux, got '%s'", val)
	}
}

func TestSetValueKey(t *testing.T) {
	data := getData()
	p := getProjector("/foo/bar", *data)

	if val, ok := p.GetValue("foo"); val != "bing" || !ok {
		t.Errorf("Expected 'bing', got '%s'", val)
	}

	p.SetValue("foo", "bingo")

	if val, ok := p.GetValue("foo"); val != "bingo" || !ok {
		t.Errorf("Expected 'bingo', got '%s'", val)
	}

	p.SetValue("baz", "bang")

	if val, ok := p.GetValue("baz"); val != "bang" || !ok {
		t.Errorf("Expected 'bang', got '%s'", val)
	}

	p = getProjector("/foo", *data)

	if val, ok := p.GetValue("foo"); val != "binx" || !ok {
		t.Errorf("Expected 'binx', got '%s'", val)
	}
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	p := getProjector("/foo/bar", *data)

	if val, ok := p.GetValue("foo"); val != "bing" || !ok {
		t.Errorf("Expected 'bing', got '%s'", val)
	}

	p.RemoveValue("foo")

	if val, ok := p.GetValue("foo"); !ok && val != "binx" {
		t.Errorf("Expected no value, got '%s'", val)
	}

	if val, ok := p.GetValue("baz"); val != "qux" || !ok {
		t.Errorf("Expected 'qux', got '%s'", val)
	}
}
