package textfmt

/* package textfmt

import (
	"strings"
	"testing"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func TestSomethingElse(t *testing.T) {

	names := []string{"Åsa", "Anna", "Åke", "Anders", "Älva"}
	// Use Swedish collation, ignoring case
	//////////////////////////////////collator := collate.New(language.Swedish, collate.IgnoreCase)
	//collator := collate.New(language.Und)
	collator := collate.New(language.Swedish, Ascending)
	collator.SortStrings(names)
	//fmt.Println(names)

	arne := strings.Compare("a", "a")

	if arne != 0 {
		t.Errorf("expected arne to be 0, got %q", arne)
	}

	arne = strings.Compare("a", "b")

	if arne != -1 {
		t.Errorf("expected arne to be -1, got %q", arne)
	}

	arne = strings.Compare("b", "a")

	if arne != 1 {
		t.Errorf("expected arne to be 1, got %q", arne)
	}

	arne = strings.Compare("A", "a")

	if arne != -1 {
		t.Errorf("expected arne to be -1, got %q", arne)
	}

	arne = strings.Compare("a", "A")

	if arne != 1 {
		t.Errorf("expected arne to be 1, got %q", arne)
	}

	arne = strings.Compare("å", "ä")

	if arne != -1 {
		t.Errorf("expected arne to be -1, got %q", arne)
	}
}

func TestSomething(t *testing.T) {

	arne := strings.Compare("a", "a")

	if arne != 0 {
		t.Errorf("expected arne to be 0, got %q", arne)
	}

	arne = strings.Compare("a", "b")

	if arne != -1 {
		t.Errorf("expected arne to be -1, got %q", arne)
	}

	arne = strings.Compare("b", "a")

	if arne != 1 {
		t.Errorf("expected arne to be 1, got %q", arne)
	}

	arne = strings.Compare("A", "a")

	if arne != -1 {
		t.Errorf("expected arne to be -1, got %q", arne)
	}

	arne = strings.Compare("a", "A")

	if arne != 1 {
		t.Errorf("expected arne to be 1, got %q", arne)
	}

	arne = strings.Compare("å", "ä")

	if arne != -1 {
		t.Errorf("expected arne to be -1, got %q", arne)
	}
}
*/
