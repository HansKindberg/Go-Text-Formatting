package yaml

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSuccess(t *testing.T) {}

func TestFormat(t *testing.T) {
	tests := []struct {
		optionsFileName string
		yamlFileName    string
	}{
		{"arne", "arne"},
		{"arne-mismatch", "arne-mismatch"},
		/* 		{"Empty-01", "Arne"},
		   		{"Empty-02", "Arne"}, */
	}

	yamlFormatter := YamlFormatter{}

	for i, tt := range tests {
		tt := tt // capture range variable
		name := fmt.Sprintf("Case %03d", i+1)
		t.Run(name, func(t *testing.T) {
			t.Parallel() // 🧠 run subtests concurrently

			input, err := getInputYaml(tt.yamlFileName)
			if err != nil {
				t.Fatalf("Failed to get input YAML reader: %v", err)
			}

			expected, err := getExpectedYaml(tt.yamlFileName)
			if err != nil {
				t.Fatalf("Failed to get expected YAML reader: %v", err)
			}

			var writer bytes.Buffer
			err = yamlFormatter.Format(YamlOptions{}, input, &writer)

			if err != nil {
				t.Fatalf("Failed to format YAML: %v", err)
			}

			actualText := writer.String()

			actualText = strings.TrimSuffix(actualText, "\n")

			expectedBytes, err := io.ReadAll(expected)
			if err != nil {
				t.Fatalf("Failed to read expected YAML: %v", err)
			}

			expectedText := string(expectedBytes)

			if actualText != expectedText {
				t.Errorf("Mismatch for %s:\nDifference:\n%s", name, cmp.Diff(expectedText, actualText))
				t.Errorf("Mismatch for %s:\nExpected:\n%s\nGot:\n%s", name, truncate(expectedText, 50), truncate(actualText, 50))
			}
		})
	}
}

func getExpectedYaml(fileName string) (io.Reader, error) {
	return getYaml("./testdata/yaml_formatter/expected", fileName)
}

func getFile(directory string, extension string, fileName string) (io.Reader, error) {
	path := filepath.Join(directory, fileName+"."+extension)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getInputYaml(fileName string) (io.Reader, error) {
	return getYaml("./testdata/yaml_formatter", fileName)
}

func getYaml(directory string, fileName string) (io.Reader, error) {
	return getFile(directory, "yaml", fileName)
}

func truncate(text string, numberOfCharacters int) string {
	runes := []rune(text) // Handles Unicode

	if len(runes) > numberOfCharacters {
		return string(runes[:numberOfCharacters]) + "…"
	}

	return text
}
