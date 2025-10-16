package textfmt

import (
	"testing"
)

func TestTextFormatterDeclaration(t *testing.T) {
	// Just a compile time test.
	var _ TextFormatter[Options]
}
