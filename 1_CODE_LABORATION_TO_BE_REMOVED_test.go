package textfmt

import (
	"testing"

	"go.yaml.in/yaml/v4"
	"golang.org/x/text/language"
)

func TestArne(t *testing.T) {
	arne := Arne{}
	berit := arne.Tammer()

	if berit != language.Und {
		t.Errorf("expected berit to be Und, got %d", berit)
	}
}

func TestBerit(t *testing.T) {
	a := yaml.Unmarshal([]byte("key: value"), nil)
	b := yaml.Node{}

	if a.Error() == b.Alias.Anchor {
		t.Errorf("expected 1 to be 1, got %d", 2)
	}

	if 1 != 2 {
		t.Errorf("expected 1 to be 1, got %d", 2)
	}
}
