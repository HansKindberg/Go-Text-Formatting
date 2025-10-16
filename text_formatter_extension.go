package textfmt

import (
	"bytes"
	"os"
	"strings"
)

// FormatBytes formats the given byte slice using the provided formatter and options.
// It returns the formatted bytes or an error if formatting fails.
func FormatBytes[T Options, F TextFormatter[T]](formatter F, options T, data []byte) ([]byte, error) {
	reader := bytes.NewReader(data)
	var writer bytes.Buffer

	if err := formatter.Format(options, reader, &writer); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

// FormatFile formats the contents of the input file and writes the result to the output file.
// It uses the provided formatter and options to perform the formatting.
// The parameters "input" and "output" are file paths.
// It returns an error if formatting fails.
func FormatFile[T Options, F TextFormatter[T]](formatter F, options T, input, output string) error {
	inFile, err := os.Open(input)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return formatter.Format(options, inFile, outFile)
}

// FormatString formats the given text using the provided formatter and options.
// It returns the formatted string or an error if formatting fails.
func FormatString[T Options, F TextFormatter[T]](formatter F, options T, text string) (string, error) {
	reader := strings.NewReader(text)
	var writer bytes.Buffer

	if err := formatter.Format(options, reader, &writer); err != nil {
		return "", err
	}

	return writer.String(), nil
}
