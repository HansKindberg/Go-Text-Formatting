package textfmt

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type CollatorFactory struct{}

func (collatorFactory *CollatorFactory) Create(caseSensitive *bool, direction *SortDirection, language *language.Tag) *collate.Collator {

	collator := collate.New(*language)

	//collator.c

	return collator
}
