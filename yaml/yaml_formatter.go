package yaml

type YamlFormatter struct {
}

func (yamlFormatter YamlFormatter) Format(options YamlOptions, text string) string {
	// Do stuff with text here, format etc.
	return text
}
