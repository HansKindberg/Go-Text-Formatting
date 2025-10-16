package textfmt

import (
	"golang.org/x/text/language"
)

type Arne struct{}

func (a *Arne) Tammer() language.Tag {
	return language.Und
}
