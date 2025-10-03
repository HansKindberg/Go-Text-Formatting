package textfmt

import (
	"testing"
)

func TestIndentationOptionsDefaults(t *testing.T) {
	opt := IndentationOptions{}

	if opt.Character != 0 {
		t.Errorf("expected default Character to be 0, got %q", opt.Character)
	}
	if opt.Enabled != false {
		t.Errorf("expected default Enabled to be false, got %v", opt.Enabled)
	}
	if opt.Size != 0 {
		t.Errorf("expected default Size to be 0, got %d", opt.Size)
	}
}

func TestIndentationOptionsAssignment(t *testing.T) {
	opt := IndentationOptions{
		Character: '\t',
		Enabled:   true,
		Size:      4,
	}

	if opt.Character != '\t' {
		t.Errorf("expected Character to be '\\t', got %q", opt.Character)
	}
	if opt.Enabled != true {
		t.Errorf("expected Enabled to be true, got %v", opt.Enabled)
	}
	if opt.Size != 4 {
		t.Errorf("expected Size to be 4, got %d", opt.Size)
	}
}

func TestNewIndentationOptions(t *testing.T) {
	opt := NewIndentationOptions()

	if opt.Character != '\t' {
		t.Errorf("expected Character to be '\\t', got %q", opt.Character)
	}
	if opt.Enabled != true {
		t.Errorf("expected Enabled to be true, got %v", opt.Enabled)
	}
	if opt.Size != 1 {
		t.Errorf("expected Size to be 1, got %d", opt.Size)
	}
}
