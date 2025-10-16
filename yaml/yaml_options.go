package yaml

import textfmt "github.com/HansKindberg/Go-Text-Formatting"

type YamlOptions struct {
	textfmt.Options
}

func NewYamlOptions() YamlOptions {
	return YamlOptions{
		Options: textfmt.Options{
			Indentation: textfmt.IndentationOptions{
				Character: ' ',
				Enabled:   true,
				Size:      2,
			},
		},
	}
}
