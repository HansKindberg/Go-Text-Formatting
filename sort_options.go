package textfmt

import "golang.org/x/text/language"

type SortOptions struct {
	CaseSensitive bool
	Direction     SortDirection
	Enabled       bool
	Language      language.Tag
}

func NewSortOptions() SortOptions {
	return SortOptions{
		CaseSensitive: false,
		Direction:     Ascending,
		Enabled:       true,
		Language:      language.Und,
	}
}
