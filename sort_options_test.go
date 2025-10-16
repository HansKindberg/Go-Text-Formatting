package textfmt

import (
	"testing"

	"golang.org/x/text/language"
)

func TestSortOptionsDefaults(t *testing.T) {
	options := SortOptions{}

	if options.CaseSensitive != false {
		t.Errorf("expected CaseSensitive to be false, got %v", options.CaseSensitive)
	}

	if options.Direction != Ascending {
		t.Errorf("expected Direction to be Ascending, got %q", options.Direction)
	}

	if options.Enabled != false {
		t.Errorf("expected Enabled to be false, got %v", options.Enabled)
	}

	if options.Language != language.Und {
		t.Errorf("expected Language to be Und, got %d", options.Language)
	}
}

func TestSortOptionsAssignment(t *testing.T) {
	options := SortOptions{
		CaseSensitive: true,
		Direction:     Descending,
		Enabled:       false,
		Language:      language.English,
	}

	if options.CaseSensitive != true {
		t.Errorf("expected CaseSensitive to be true, got %v", options.CaseSensitive)
	}

	if options.Direction != Descending {
		t.Errorf("expected Direction to be Descending, got %q", options.Direction)
	}

	if options.Enabled != false {
		t.Errorf("expected Enabled to be false, got %v", options.Enabled)
	}

	if options.Language != language.English {
		t.Errorf("expected Language to be English, got %d", options.Language)
	}
}

func TestNewSortOptions(t *testing.T) {
	options := NewSortOptions()

	if options.CaseSensitive != false {
		t.Errorf("expected CaseSensitive to be false, got %v", options.CaseSensitive)
	}

	if options.Direction != Ascending {
		t.Errorf("expected Direction to be Ascending, got %q", options.Direction)
	}

	if options.Enabled != true {
		t.Errorf("expected Enabled to be true, got %v", options.Enabled)
	}

	if options.Language != language.Und {
		t.Errorf("expected Language to be Und, got %d", options.Language)
	}
}
