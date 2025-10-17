package textfmt

import (
	"io"
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
	a := t.TempDir()
	t.Errorf("Directory \"%s\"", a)

	//t.Fail()

	/* 	content := "File content."

	   	tmpFile, err := os.CreateTemp(t.TempDir(), "testfile-*.txt")
	   	if err != nil {
	   		t.Fatalf("Failed to create temp file: %v", err)
	   	}
	   	defer tmpFile.Close()

	   	if _, err := tmpFile.WriteString(content); err != nil {
	   		t.Fatalf("Failed to write to temp file: %v", err)
	   	}

	   	actualPath, err := FormatFile(testFormatter{}, Options{}, tmpFile.Name())
	   	if err != nil {
	   		t.Fatalf("Unexpected error: %v", err)
	   	}

	   	formattedBytes, err := os.ReadFile(actualPath)
	   	if err != nil {
	   		t.Fatalf("Failed to read formatted file: %v", err)
	   	}

	   	expected := "Formatted: " + content
	   	if string(formattedBytes) != expected {
	   		t.Errorf("Expected %q, got %q", expected, string(formattedBytes))
	   	} */
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
