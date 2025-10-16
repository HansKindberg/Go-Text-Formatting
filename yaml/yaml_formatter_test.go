package yaml

import (
	"io/ioutil"
	"os"
	"testing"

	"go.yaml.in/yaml/v4"
)

func Test_4(t *testing.T) {
	//t.Fail()
	t.Errorf("Det blev ett Arne fel")
}

func Test_5(t *testing.T) {
	source, err := os.Open("./in-5.yaml")
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer source.Close()
}

func Test_6(t *testing.T) {
	decoder := yaml.NewDecoder(nil)
	if decoder == nil {
		t.Errorf("Decoder should not be nil")
	}

	source, err := os.Open("./in-5.yaml")
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}

	defer source.Close()
}

func getFileContent(fileName string) string {
	content, err := os.ReadFile("./in-5.yaml")

	if err == nil {

	}

	if content == nil {

	}

	fileBytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	yamlContent := string(fileBytes)

	if yamlContent == "" {

	}

	return "Arne"
}
