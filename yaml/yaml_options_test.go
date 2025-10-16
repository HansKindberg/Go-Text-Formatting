package yaml

import (
	"testing"
)

func TestNewYamlOptions(t *testing.T) {
	options := NewYamlOptions()

	if options.Indentation.Character != ' ' {
		t.Errorf("expected Character to be ' ', got %q", options.Indentation.Character)
	}

	if options.Indentation.Enabled != true {
		t.Errorf("expected Enabled to be true, got %v", options.Indentation.Enabled)
	}

	if options.Indentation.Size != 2 {
		t.Errorf("expected Size to be 2, got %d", options.Indentation.Size)
	}
}
