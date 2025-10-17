package textfmt

import (
	"io"
	"os"
	"path"
	"testing"
)

type testFormatter struct{}

func (f testFormatter) Format(options Options, reader io.Reader, writer io.Writer) error {
	bytes, err := io.ReadAll(reader)

	if err != nil {
		return err
	}

	text := "Formatted: " + string(bytes)

	if _, err := writer.Write([]byte(text)); err != nil {
		return err
	}

	return nil
}

func TestFormatBytes(t *testing.T) {
	bytes := []byte("Some text.")

	actualBytes, err := FormatBytes(testFormatter{}, Options{}, bytes)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedBytes := []byte("Formatted: " + string(bytes))

	if string(actualBytes) != string(expectedBytes) {
		t.Errorf("Expected \"%s\", got \"%s\"", expectedBytes, actualBytes)
	}
}

func TestFormatFile(t *testing.T) {
	input := "./testdata/text_formatter_extension/input.txt"
	inputBytes, err := os.ReadFile(input)
	if err != nil {
		t.Fatalf("Failed to read input file: %v", err)
	}
	inputText := string(inputBytes)

	output := path.Join(t.TempDir(), "output.txt")
	err = FormatFile(testFormatter{}, Options{}, input, output)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	outputBytes, err := os.ReadFile(output)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	actualText := string(outputBytes)
	expectedText := "Formatted: " + inputText

	if actualText != expectedText {
		t.Errorf("Expected \"%s\", got \"%s\"", expectedText, actualText)
	}
}

func TestFormatString(t *testing.T) {
	text := "Some text."

	actualText, err := FormatString(testFormatter{}, Options{}, text)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedText := "Formatted: " + text

	if actualText != expectedText {
		t.Errorf("Expected \"%s\", got \"%s\"", expectedText, actualText)
	}
}
