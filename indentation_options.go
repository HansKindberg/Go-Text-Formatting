package textfmt

type IndentationOptions struct {
	Character rune
	Enabled   bool
	Size      int
}

func NewIndentationOptions() IndentationOptions {
	return IndentationOptions{
		Character: '\t',
		Enabled:   true,
		Size:      1,
	}
}
