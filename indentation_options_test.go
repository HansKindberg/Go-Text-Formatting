package textfmt

import (
	"testing"
)

func TestIndentationOptionsDefaults(t *testing.T) {
	options := IndentationOptions{}

	if options.Character != 0 {
		t.Errorf("expected default Character to be 0, got %q", options.Character)
	}

	if options.Enabled != false {
		t.Errorf("expected default Enabled to be false, got %v", options.Enabled)
	}

	if options.Size != 0 {
		t.Errorf("expected default Size to be 0, got %d", options.Size)
	}
}

func TestIndentationOptionsAssignment(t *testing.T) {
	options := IndentationOptions{
		Character: '\t',
		Enabled:   true,
		Size:      4,
	}

	if options.Character != '\t' {
		t.Errorf("expected Character to be '\\t', got %q", options.Character)
	}

	if options.Enabled != true {
		t.Errorf("expected Enabled to be true, got %v", options.Enabled)
	}

	if options.Size != 4 {
		t.Errorf("expected Size to be 4, got %d", options.Size)
	}
}

func TestNewIndentationOptions(t *testing.T) {
	options := NewIndentationOptions()

	if options.Character != '\t' {
		t.Errorf("expected Character to be '\\t', got %q", options.Character)
	}

	if options.Enabled != true {
		t.Errorf("expected Enabled to be true, got %v", options.Enabled)
	}

	if options.Size != 1 {
		t.Errorf("expected Size to be 1, got %d", options.Size)
	}
}
