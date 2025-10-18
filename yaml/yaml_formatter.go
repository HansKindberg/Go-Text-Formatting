package yaml

import (
	"io"

	"go.yaml.in/yaml/v4"
)

type YamlFormatter struct{}

func (f YamlFormatter) Format(options YamlOptions, reader io.Reader, writer io.Writer) error {
	decoder := yaml.NewDecoder(reader)

	var tree interface{}
	if err := decoder.Decode(&tree); err != nil {
		return err
	}

	// TODO: sort maps recursively here

	encoder := yaml.NewEncoder(writer)

	if options.Indentation.Enabled {
		encoder.SetIndent(options.Indentation.Size)
	}

	return encoder.Encode(tree) // Add a trailing newline, https://chatgpt.com/c/68f34b46-2c2c-832c-8a4a-9d113ab5502c
}

/* func (yamlFormatter *YamlFormatter) Format(options *YamlOptions, text string) (string, error) {
	if yamlFormatter == nil {
		return "", fmt.Errorf("yamlFormatter cannot be nil")
	}

	if options == nil {
		return "", fmt.Errorf("options cannot be nil")
	}

	// Do stuff with text here, format etc.
	return text, nil
} */

// https://chatgpt.com/c/68e53de7-a5c8-832e-b1b3-395db3a65b29
// https://chatgpt.com/c/68e5f56f-4fd0-8332-9f52-f6d7936dabc0
/* func (f YamlFormatter) Format(options YamlOptions, reader io.Reader, writer io.Writer) error {
	decoder := yaml.NewDecoder(reader)

	var tree interface{}
	if err := decoder.Decode(&tree); err != nil {
		return err
	}

	if options == nil {}

	// TODO: sort maps recursively here

	encoder := yaml.NewEncoder(writer)
	encoder.SetIndent(2)
	return encoder.Encode(tree)
} */
